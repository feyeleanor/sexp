package sexp

import "reflect"

type Collectable interface {
	Collect(f interface{}) (r interface{}, ok bool)
}

func makeContainer(container interface{}) (r interface{}) {
	switch c := reflect.ValueOf(container); c.Kind() {
	case reflect.Slice:			r = reflect.MakeSlice(c.Type(), c.Len(), c.Len()).Interface()

	case reflect.Map:			r = reflect.MakeMap(c.Type()).Interface()

	case reflect.Chan:			r = reflect.MakeChan(c.Type(), c.Len()).Interface()
	}
	return
}

func collectIndexable(container Indexable, t interface{}) (r interface{}, ok bool) {
	c := makeContainer(container).(Indexable)
	end := c.Len()
	switch t := t.(type) {
	case func(interface{}) interface{}:						for i := 0; i < end; i++ {
																c.Set(i, t(container.At(i)))
															}
															ok = true

	case func(int, interface{}) interface{}:				for i := 0; i < end; i++ {
																c.Set(i, t(i, container.At(i)))
															}
															ok = true

	case func(interface{}, interface{}) interface{}:		for i := 0; i < end; i++ {
																c.Set(i, t(i, container.At(i)))
															}
															ok = true

	default:												if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
																switch t.Type().NumIn() {
																case 1:				for i := 0; i < end; i++ {
																						c.Set(i, t.Call(valueslice(container.At(i)))[0].Interface())
																					}
																					ok = true

																case 2:				for i := 0; i < end; i++ {
																						c.Set(i, t.Call(valueslice(i, container.At(i)))[0].Interface())
																					}
																					ok = true
																}
															}
	}
	if ok {
		r = c
	}
	return
}

func collectMappable(container Mappable, f interface{}) (r interface{}, ok bool) {
	switch f := f.(type) {
	case func(interface{}) interface{}:					if c, allocated := makeContainer(container).(Mappable); allocated {
															for _, v := range container.Keys() {
																c.Set(v, f(container.At(v)))
															}
															r, ok = c, true
														}

	case func(interface{}, interface{}) interface{}:	if c, allocated := makeContainer(container).(Mappable); allocated {
															for _, v := range container.Keys() {
																c.Set(v, f(v, container.At(v)))
															}
															r, ok = c, true
														}

	default:											if c, allocated := makeContainer(container).(Mappable); allocated {
															if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
																switch f.Type().NumIn() {
																case 1:				for _, v := range container.Keys() {
																						c.Set(v, f.Call(valueslice(container.At(v))))
																					}
																					r, ok = c, true

																case 2:				for _, v := range container.Keys() {
																						c.Set(v, f.Call(valueslice(v, container.At(v))))
																					}
																					r, ok = c, true
																}
															}
														}
	}
	return
}

func collectSlice(s reflect.Value, f interface{}) (r interface{}, ok bool) {
	end := s.Len()
	switch f := f.(type) {
	case func(interface{}) interface{}:					c := reflect.MakeSlice(s.Type(), end, end)
														for i := 0; i < end; i++ {
															c.Index(i).Set(reflect.ValueOf(f(s.Index(i).Interface())))
														}
														r = c.Interface()
														ok = true

	case func(int, interface{}) interface{}:			c := reflect.MakeSlice(s.Type(), end, end)
														for i := 0; i < end; i++ {
															c.Index(i).Set(reflect.ValueOf(f(i, s.Index(i).Interface())))
														}
														r = c.Interface()
														ok = true

	case func(interface{}, interface{}) interface{}:	c := reflect.MakeSlice(s.Type(), end, end)
														for i := 0; i < end; i++ {
															c.Index(i).Set(reflect.ValueOf(f(i, s.Index(i).Interface())))
														}
														r = c.Interface()
														ok = true

	default:											if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
															end := s.Len()
															c := reflect.MakeSlice(s.Type(), end, end)
															switch f.Type().NumIn() {
															case 1:				for i := 0; i < end; i++ {
																					c.Index(i).Set(f.Call([]reflect.Value{ s.Index(i) })[0])
																				}
																				ok = true

															case 2:				for i := 0; i < end; i++ {
																					c.Index(i).Set(f.Call([]reflect.Value{ reflect.ValueOf(i), s.Index(i) })[0])
																				}
																				ok = true
															}
															if ok {
																r = c.Interface()
															}
														}
	}
	return
}

func collectMap(m reflect.Value, f interface{}) (r interface{}, ok bool) {
	switch f := f.(type) {
	case func(interface{}) interface{}:					c := reflect.MakeMap(m.Type())
                            for iter := m.MapRange(); iter.Next(); {
															c.SetMapIndex(iter.Key(), reflect.ValueOf(f(iter.Value().Interface())))
														}
														r = c.Interface()
														ok = true

	case func(interface{}, interface{}) interface{}:	c := reflect.MakeMap(m.Type())
                            for iter := m.MapRange(); iter.Next(); {
															c.SetMapIndex(iter.Key(), reflect.ValueOf(f(iter.Key().Interface(), iter.Value().Interface())))
														}
														r = c.Interface()
														ok = true

	default:											if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
															c := reflect.MakeMap(m.Type())
															switch f.Type().NumIn() {
															case 1:		for iter := m.MapRange(); iter.Next(); {
																					c.SetMapIndex(iter.Key(), f.Call([]reflect.Value{ iter.Value() })[0])
																				}
																				ok = true

															case 2:		for iter := m.MapRange(); iter.Next(); {
																					c.SetMapIndex(iter.Key(), f.Call([]reflect.Value{ iter.Key(), iter.Value() })[0])
																				}
																				ok = true
															}
															if ok {
																r = c.Interface()
															}
														}

	}
	return
}

func Collect(container interface{}, f interface{}) (r interface{}, ok bool) {
	switch container := container.(type) {
	case Collectable:		r, ok = container.Collect(f)

	case Indexable:			r, ok = collectIndexable(container, f)

	case Mappable:			r, ok = collectMappable(container, f)

	default:				switch c := reflect.ValueOf(container); c.Kind() {
							case reflect.Slice:		r, ok = collectSlice(c, f)

							case reflect.Map:		r, ok = collectMap(c, f)
							}
	}
	return
}
