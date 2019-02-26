package sexp

import "reflect"

type Transformable interface {
	Transform(interface{}) bool
}

func transformIndexable(container Indexable, t interface{}) (ok bool) {
	end := container.Len()
	switch t := t.(type) {
	case func(interface{}) interface{}:						for i := 0; i < end; i++ {
																container.Set(i, t(container.At(i)))
															}
															ok = true

	case func(int, interface{}) interface{}:				for i := 0; i < end; i++ {
																container.Set(i, t(i, container.At(i)))
															}
															ok = true

	case func(interface{}, interface{}) interface{}:		for i := 0; i < end; i++ {
																container.Set(i, t(i, container.At(i)))
															}
															ok = true

	default:												if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
																switch t.Type().NumIn() {
																case 1:				for i := 0; i < end; i++ {
																						container.Set(i, t.Call(valueslice(container.At(i)))[0].Interface())
																					}
																					ok = true

																case 2:				for i := 0; i < end; i++ {
																						container.Set(i, t.Call(valueslice(i, container.At(i)))[0].Interface())
																					}
																					ok = true
																}
															}
	}
	return
}

func transformMappable(container Mappable, t interface{}) (ok bool) {
	switch t := t.(type) {
	case func(interface{}) interface{}:					for _, v := range container.Keys() {
															container.Set(v, t(container.At(v)))
														}
														ok = true

	case func(interface{}, interface{}) interface{}:	for _, v := range container.Keys() {
															container.Set(v, t(v, container.At(v)))
														}
														ok = true

	default:											if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
															switch t.Type().NumIn() {
															case 1:				for _, v := range container.Keys() {
																					container.Set(v, t.Call(valueslice(container.At(v))))
																				}
																				ok = true

															case 2:				for _, v := range container.Keys() {
																					container.Set(v, t.Call(valueslice(v, container.At(v))))
																				}
																				ok = true
															}
														}
	}
	return
}

func transformSlice(s reflect.Value, t interface{}) (ok bool) {
	var v	reflect.Value
	end := s.Len()
	switch t := t.(type) {
	case func(interface{}) interface{}:					for i := 0; i < end; i++ {
															v = s.Index(i)
															v.Set(reflect.ValueOf(t(v.Interface())))
														}
														ok = true

	case func(int, interface{}) interface{}:			for i := 0; i < end; i++ {
															v = s.Index(i)
															v.Set(reflect.ValueOf(t(i, v.Interface())))
														}
														ok = true

	case func(interface{}, interface{}) interface{}:	for i := 0; i < end; i++ {
															v = s.Index(i)
															v.Set(reflect.ValueOf(t(i, v.Interface())))
														}
														ok = true
	}
	return
}

func transformMap(m reflect.Value, t interface{}) (ok bool) {
	switch t := t.(type) {
	case func(interface{}) interface{}:	for iter := m.MapRange(); iter.Next(); {
															m.SetMapIndex(iter.Key(), reflect.ValueOf(t(iter.Value().Interface())))
														}
														ok = true

	case func(interface{}, interface{}) interface{}: for iter := m.MapRange(); iter.Next(); {
															m.SetMapIndex(iter.Key(), reflect.ValueOf(t(iter.Key().Interface(), iter.Value().Interface())))
														}
														ok = true
	}
	return
}

func transform(container, t interface{}) (ok bool) {
	var v	reflect.Value

	switch c := reflect.ValueOf(container); c.Kind() {
	case reflect.Slice:		if ok = transformSlice(c, t); !ok {
								if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
									end := c.Len()
									switch t.Type().NumIn() {
									case 1:				for i := 0; i < end; i++ {
															v = c.Index(i)
															v.Set(t.Call([]reflect.Value{ v })[0])
														}
														ok = true

									case 2:				for i := 0; i < end; i++ {
															v = c.Index(i)
															v.Set(t.Call([]reflect.Value{ reflect.ValueOf(i), v })[0])
														}
														ok = true
									}
								}
							}

	case reflect.Map:		if ok = transformMap(c, t); !ok {
								if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
									switch t.Type().NumIn() {
									case 1:   for iter := c.MapRange(); iter.Next(); {
															c.SetMapIndex(iter.Key(), t.Call([]reflect.Value{ iter.Value() })[0])
														}
														ok = true

									case 2:   for iter := c.MapRange(); iter.Next(); {
															c.SetMapIndex(iter.Key(), t.Call([]reflect.Value{ iter.Key(), iter.Value() })[0])
														}
														ok = true
									}
								}
							}
	}
	return
}

func Transform(container, t interface{}) (ok bool) {
	switch container := container.(type) {
	case Transformable:		ok = container.Transform(t)

	case Indexable:			ok = transformIndexable(container, t)

	case Mappable:			ok = transformMappable(container, t)

	default:				ok = transform(container, t)
	}
	return
}
