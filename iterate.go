package sexp

import(
	"github.com/feyeleanor/slices"
	"reflect"
)

func eachIndexedReader(container IndexedReader, f interface{}) {
	end := container.Len()
	switch f := f.(type) {
	case func(interface{}):					for i := 0; i < end; i++ {
												f(container.At(i))
											}

	case func(int, interface{}):			for i := 0; i < end; i++ {
												f(i, container.At(i))
											}

	case func(interface{}, interface{}):	for i := 0; i < end; i++ {
												f(i, container.At(i))
											}

	default:								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
												switch f.Type().NumIn() {
												case 1:				for i := 0; i < end; i++ {
																		f.Call(slices.VList(container.At(i)))
																	}

												case 2:				for i := 0; i < end; i++ {
																		f.Call(slices.VList(i, container.At(i)))
																	}

												default:			panic(f)
												}
											} else {
												panic(f)
											}
	}
}

func eachMappedReader(container MappedReader, f interface{}) {
	switch f := f.(type) {
	case func(interface{}):					for _, v := range container.Keys() {
												f(container.At(v))
											}

	case func(interface{}, interface{}):	for _, v := range container.Keys() {
												f(v, container.At(v))
											}

	default:								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
												switch f.Type().NumIn() {
												case 1:				for _, v := range container.Keys() {
																		f.Call(slices.VList(container.At(v)))
																	}

												case 2:				for _, v := range container.Keys() {
																		f.Call(slices.VList(v, container.At(v)))
																	}

												default:			panic(f)
												}
											} else {
												panic(f)
											}
	}
}

func rangeSlice(s reflect.Value, f interface{}) (ok bool) {
	end := s.Len()
	switch f := f.(type) {
	case func(interface{}):					for i := 0; i < end; i++ {
												f(s.Index(i).Interface())
											}
											ok = true

	case func(int, interface{}):			for i := 0; i < end; i++ {
												f(i, s.Index(i).Interface())
											}
											ok = true

	case func(interface{}, interface{}):	for i := 0; i < end; i++ {
												f(i, s.Index(i).Interface())
											}
											ok = true
	}
	return
}

func rangeMap(m reflect.Value, f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):					for _, key := range m.MapKeys() {
												f(m.MapIndex(key).Interface())
											}
											ok = true

	case func(interface{}, interface{}):	for _, key := range m.MapKeys() {
												f(key.Interface(), m.MapIndex(key).Interface())
											}
											ok = true
	}
	return
}

func each(container, f interface{}) {
	switch c := reflect.ValueOf(container); c.Kind() {
	case reflect.Slice:		if !rangeSlice(c, f) {
								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
									end := c.Len()
									switch f.Type().NumIn() {
									case 1:				for i := 0; i < end; i++ {
															f.Call([]reflect.Value{ c.Index(i) })
														}

									case 2:				for i := 0; i < end; i++ {
															f.Call([]reflect.Value{ reflect.ValueOf(i), c.Index(i) })
														}

									default:			panic(f)
									}
								} else {
									panic(f)
								}
							}
							
	case reflect.Map:		if !rangeMap(c, f) {
								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
									switch f.Type().NumIn() {
									case 1:				for _, key := range c.MapKeys() {
															f.Call([]reflect.Value{ c.MapIndex(key) })
														}

									case 2:				for _, key := range c.MapKeys() {
															f.Call([]reflect.Value{ key, c.MapIndex(key) })
														}

									default:			panic(f)
									}
								} else {
									panic(f)
								}
							}
	}
}

func Each(container, f interface{}) {
	switch container := container.(type) {
	case Iterable:			container.Each(f)

	case IndexedReader:		eachIndexedReader(container, f)

	case MappedReader:		eachMappedReader(container, f)

	default:				each(container, f)
							
	}
}

func Cycle(container interface{}, count int, f func(interface{})) (i int) {
	switch {
	case count == 0:	for ; ; i++ { Each(container, f) }
	default:			for ; i < count; i++ { Each(container, f) }
	}
	return
}