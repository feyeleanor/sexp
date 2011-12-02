package sexp

import(
	"github.com/feyeleanor/slices"
	"reflect"
)

type Iterable interface {
	Each(interface{})
}


func rangeIndexedReader(container IndexedReader, f interface{}) {
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

func rangeMappedReader(container MappedReader, f interface{}) {
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

func rangeChannel(c reflect.Value, f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):					for {
												if v, open := c.Recv(); open {
													f(v.Interface())
												} else {
													break
												}
											}
											ok = true

	case func(int, interface{}):			for count := 0; ; count++ {
												if v, open := c.Recv(); open {
													f(count, v.Interface())
												} else {
													break
												}
											}
											ok = true

	case func(interface{}, interface{}):	for count := 0; ; count++ {
												if v, open := c.Recv(); open {
													f(count, v.Interface())
												} else {
													break
												}
											}
											ok = true
	}
	return
}

func rangeGenerator(p reflect.Value, f interface{}) (ok bool) {
	if t := p.Type(); t.NumOut() > 0 {
		switch t.NumIn() {
		case 0:			switch f := f.(type) {
						case func(interface{}):					for v := p.Call([]reflect.Value{}); !v[0].IsNil(); v = p.Call([]reflect.Value{}) {
																	f(v[0].Interface())
																}
																ok = true

						case func(int, interface{}):			for count, v := 0, p.Call([]reflect.Value{}); !v[0].IsNil(); v = p.Call([]reflect.Value{}) {
																	f(count, v[0].Interface())
																	count++
																} 
																ok = true

						case func(interface{}, interface{}):	for count, v := 0, p.Call([]reflect.Value{}); !v[0].IsNil(); v = p.Call([]reflect.Value{}) {
																	f(count, v[0].Interface())
																	count++
																}
																ok = true
						}

		case 1:			count := 0
						switch f := f.(type) {
						case func(interface{}):					for v := p.Call(slices.VList(count)); !v[0].IsNil(); v = p.Call(slices.VList(count)) {
																	f(v[0].Interface())
																	count++
																}
																ok = true

						case func(int, interface{}):			for v := p.Call(slices.VList(count)); !v[0].IsNil(); v = p.Call(slices.VList(count)) {
																	f(count, v[0].Interface())
																	count++
																}
																ok = true

						case func(interface{}, interface{}):	for v := p.Call(slices.VList(count)); !v[0].IsNil(); v = p.Call(slices.VList(count)) {
																	f(count, v[0].Interface())
																	count++
																}
																ok = true
						}
		}
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

	case reflect.Chan:		if !rangeChannel(c, f) {
								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
									switch f.Type().NumIn() {
									case 1:				for {
															if v, done := c.Recv(); !done {
																f.Call([]reflect.Value{ v })
															} else {
																break
															}
														}

									case 2:				for count := 0; ; count++ {
															if v, done := c.Recv(); !done {
																f.Call([]reflect.Value{ reflect.ValueOf(count), v })
															} else {
																break
															}
														}

									default:
									}
								}
							}

	case reflect.Func:		if !rangeGenerator(c, f) {
							}
	}
}

func Each(container, f interface{}) {
	switch container := container.(type) {
	case Iterable:			container.Each(f)

	case IndexedReader:		rangeIndexedReader(container, f)

	case MappedReader:		rangeMappedReader(container, f)

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