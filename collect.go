package sexp

import (
	"github.com/feyeleanor/slices"
	"reflect"
)

type Collectable interface {
	Collect(interface{}) interface{}
}

func makeContainer(container interface{}) (r interface{}) {
	switch c := reflect.ValueOf(container); c.Kind() {
	case reflect.Slice:			r = reflect.MakeSlice(c.Type(), c.Len(), c.Len()).Interface()

	case reflect.Map:			r = reflect.MakeMap(c.Type()).Interface()

	case reflect.Chan:			r = reflect.MakeChan(c.Type(), c.Len()).Interface()
	}
	return
}

func collectIndexable(container Indexable, t interface{}) interface{} {
	c := makeContainer(container).(Indexable)
	end := c.Len()
	switch t := t.(type) {
	case func(interface{}) interface{}:						for i := 0; i < end; i++ {
																c.Set(i, t(container.At(i)))
															}

	case func(int, interface{}) interface{}:				for i := 0; i < end; i++ {
																c.Set(i, t(i, container.At(i)))
															}

	case func(interface{}, interface{}) interface{}:		for i := 0; i < end; i++ {
																c.Set(i, t(i, container.At(i)))
															}

	default:												if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
																switch t.Type().NumIn() {
																case 1:				for i := 0; i < end; i++ {
																						c.Set(i, t.Call(slices.VList(container.At(i)))[0].Interface())
																					}

																case 2:				for i := 0; i < end; i++ {
																						c.Set(i, t.Call(slices.VList(i, container.At(i)))[0].Interface())
																					}

																default:			panic(t)
																}
															} else {
																panic(t)
															}
	}
	return c
}

func collectMappable(container Mappable, t interface{}) (r interface{}) {
	if c, ok := makeContainer(container).(Mappable); ok {
		switch t := t.(type) {
		case func(interface{}) interface{}:					for _, v := range container.Keys() {
																c.Set(v, t(container.At(v)))
															}

		case func(interface{}, interface{}) interface{}:	for _, v := range container.Keys() {
																c.Set(v, t(v, container.At(v)))
															}

		default:											if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
																switch t.Type().NumIn() {
																case 1:				for _, v := range container.Keys() {
																						c.Set(v, t.Call(slices.VList(container.At(v))))
																					}

																case 2:				for _, v := range container.Keys() {
																						c.Set(v, t.Call(slices.VList(v, container.At(v))))
																					}

																default:			panic(t)
																}
															} else {
																panic(t)
															}
		}
		r = c
	}
	return
}

func collectSlice(s reflect.Value, t interface{}) (r interface{}) {
	end := s.Len()
	c := reflect.MakeSlice(s.Type(), end, end)
	switch t := t.(type) {
	case func(interface{}) interface{}:					for i := 0; i < end; i++ {
															c.Index(i).Set(reflect.ValueOf(t(s.Index(i).Interface())))
														}
														r = c.Interface()

	case func(int, interface{}) interface{}:			for i := 0; i < end; i++ {
															c.Index(i).Set(reflect.ValueOf(t(i, s.Index(i).Interface())))
														}
														r = c.Interface()

	case func(interface{}, interface{}) interface{}:	for i := 0; i < end; i++ {
															c.Index(i).Set(reflect.ValueOf(t(i, s.Index(i).Interface())))
														}
														r = c.Interface()
	}
	return
}

func collectMap(m reflect.Value, t interface{}) (r interface{}) {
	switch t := t.(type) {
	case func(interface{}) interface{}:					c := reflect.MakeMap(m.Type())
														for _, key := range m.MapKeys() {
															c.SetMapIndex(key, reflect.ValueOf(t(m.MapIndex(key).Interface())))
														}
														r = c.Interface()

	case func(interface{}, interface{}) interface{}:	c := reflect.MakeMap(m.Type())
														for _, key := range m.MapKeys() {
															c.SetMapIndex(key, reflect.ValueOf(t(key.Interface(), m.MapIndex(key).Interface())))
														}
														r = c.Interface()
	}
	return
}

func collect(container, t interface{}) (r interface{}) {
	switch c := reflect.ValueOf(container); c.Kind() {
	case reflect.Slice:		if r = collectSlice(c, t); r == nil {
								if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
									end := c.Len()
									s := reflect.MakeSlice(c.Type(), end, end)
									switch t.Type().NumIn() {
									case 1:				for i := 0; i < end; i++ {
															s.Index(i).Set(t.Call([]reflect.Value{ c.Index(i) })[0])
														}

									case 2:				for i := 0; i < end; i++ {
															s.Index(i).Set(t.Call([]reflect.Value{ reflect.ValueOf(i), c.Index(i) })[0])
														}

									default:			panic(t)
									}
									r = s.Interface()
								} else {
									panic(t)
								}
							}
							
	case reflect.Map:		if r = collectMap(c, t); r == nil {
								if t := reflect.ValueOf(t); t.Kind() == reflect.Func {
									m := reflect.MakeMap(c.Type())
									switch t.Type().NumIn() {
									case 1:				for _, key := range c.MapKeys() {
															m.SetMapIndex(key, t.Call([]reflect.Value{ c.MapIndex(key) })[0])
														}

									case 2:				for _, key := range c.MapKeys() {
															m.SetMapIndex(key, t.Call([]reflect.Value{ key, c.MapIndex(key) })[0])
														}

									default:			panic(t)
									}
									r = m.Interface()
								} else {
									panic(t)
								}
							}
	default:
	}
	return
}

func Collect(container interface{}, t interface{}) (r interface{}) {
	switch container := container.(type) {
	case Collectable:		r = container.Collect(t)

	case Indexable:			r = collectIndexable(container, t)

	case Mappable:			r = collectMappable(container, t)

	default:				r = collect(container, t)
	}
	return
}