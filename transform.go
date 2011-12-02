package sexp

import(
	"github.com/feyeleanor/slices"
	"reflect"
)

type Transformable interface {
	Transform(interface{})
}


func transformIndexable(container Indexable, t interface{}) {
	end := container.Len()
	switch t := t.(type) {
	case func(interface{}) interface{}:						for i := 0; i < end; i++ {
																container.Set(i, t(container.At(i)))
															}

	case func(int, interface{}) interface{}:				for i := 0; i < end; i++ {
																container.Set(i, t(i, container.At(i)))
															}

	case func(interface{}, interface{}) interface{}:		for i := 0; i < end; i++ {
																container.Set(i, t(i, container.At(i)))
															}

	default:												if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
																switch t.Type().NumIn() {
																case 1:				for i := 0; i < end; i++ {
																						container.Set(i, t.Call(slices.VList(container.At(i)))[0].Interface())
																					}

																case 2:				for i := 0; i < end; i++ {
																						container.Set(i, t.Call(slices.VList(i, container.At(i)))[0].Interface())
																					}

																default:			panic(t)
																}
															} else {
																panic(t)
															}
	}
}

func transformMappable(container Mappable, t interface{}) {
	switch t := t.(type) {
	case func(interface{}) interface{}:					for _, v := range container.Keys() {
															container.Set(v, t(container.At(v)))
														}

	case func(interface{}, interface{}) interface{}:	for _, v := range container.Keys() {
															container.Set(v, t(v, container.At(v)))
														}

	default:											if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
															switch t.Type().NumIn() {
															case 1:				for _, v := range container.Keys() {
																					container.Set(v, t.Call(slices.VList(container.At(v))))
																				}

															case 2:				for _, v := range container.Keys() {
																					container.Set(v, t.Call(slices.VList(v, container.At(v))))
																				}

															default:			panic(t)
															}
														} else {
															panic(t)
														}
	}
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
	case func(interface{}) interface{}:					for _, key := range m.MapKeys() {
															m.SetMapIndex(key, reflect.ValueOf(t(m.MapIndex(key).Interface())))
														}
														ok = true

	case func(interface{}, interface{}) interface{}:	for _, key := range m.MapKeys() {
															m.SetMapIndex(key, reflect.ValueOf(t(key.Interface(), m.MapIndex(key).Interface())))
														}
														ok = true
	}
	return
}

func transform(container, t interface{}) {
	var v	reflect.Value

	switch c := reflect.ValueOf(container); c.Kind() {
	case reflect.Slice:		if !transformSlice(c, t) {
								if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
									end := c.Len()
									switch t.Type().NumIn() {
									case 1:				for i := 0; i < end; i++ {
															v = c.Index(i)
															v.Set(t.Call([]reflect.Value{ v })[0])
														}

									case 2:				for i := 0; i < end; i++ {
															v = c.Index(i)
															v.Set(t.Call([]reflect.Value{ reflect.ValueOf(i), v })[0])
														}

									default:			panic(t)
									}
								} else {
									panic(t)
								}
							}
							
	case reflect.Map:		if !transformMap(c, t) {
								if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
									switch t.Type().NumIn() {
									case 1:				for _, key := range c.MapKeys() {
															c.SetMapIndex(key, t.Call([]reflect.Value{ c.MapIndex(key) })[0])
														}

									case 2:				for _, key := range c.MapKeys() {
															c.SetMapIndex(key, t.Call([]reflect.Value{ key, c.MapIndex(key) })[0])
														}

									default:			panic(t)
									}
								} else {
									panic(t)
								}
							}
	default:
	}
}

func Transform(container, t interface{}) {
	switch container := container.(type) {
	case Transformable:		container.Transform(t)

	case Indexable:			transformIndexable(container, t)

	case Mappable:			transformMappable(container, t)

	default:				transform(container, t)
	}
}