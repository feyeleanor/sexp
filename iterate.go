package sexp

import "reflect"

type Iterable interface {
	Each(interface{}) bool
}

func eachIndexedReader(container IndexedReader, f interface{}) (ok bool) {
	end := container.Len()
	switch f := f.(type) {
	case func(interface{}):					for i := 0; i < end; i++ {
												f(container.At(i))
											}
											ok = true

	case func(int, interface{}):			for i := 0; i < end; i++ {
												f(i, container.At(i))
											}
											ok = true

	case func(interface{}, interface{}):	for i := 0; i < end; i++ {
												f(i, container.At(i))
											}
											ok = true

	default:								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
												switch f.Type().NumIn() {
												case 1:				for i := 0; i < end; i++ {
																		f.Call(valueslice(container.At(i)))
																	}
																	ok = true

												case 2:				for i := 0; i < end; i++ {
																		f.Call(valueslice(i, container.At(i)))
																	}
																	ok = true
												}
											}
	}
	return
}

func eachMappedReader(container MappedReader, f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):					for _, v := range container.Keys() {
												f(container.At(v))
											}
											ok = true

	case func(interface{}, interface{}):	for _, v := range container.Keys() {
												f(v, container.At(v))
											}
											ok = true

	default:								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
												switch f.Type().NumIn() {
												case 1:				for _, v := range container.Keys() {
																		f.Call(valueslice(container.At(v)))
																	}
																	ok = true

												case 2:				for _, v := range container.Keys() {
																		f.Call(valueslice(v, container.At(v)))
																	}
																	ok = true
												}
											}
	}
	return
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

	default:								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
												end := s.Len()
												switch f.Type().NumIn() {
												case 1:				for i := 0; i < end; i++ {
																		f.Call([]reflect.Value{ s.Index(i) })
																	}
																	ok = true

												case 2:				for i := 0; i < end; i++ {
																		f.Call([]reflect.Value{ reflect.ValueOf(i), s.Index(i) })
																	}
																	ok = true
												}
											}
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

	default:								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
												switch f.Type().NumIn() {
												case 1:				for _, key := range m.MapKeys() {
																		f.Call([]reflect.Value{ m.MapIndex(key) })
																	}
																	ok = true

												case 2:				for _, key := range m.MapKeys() {
																		f.Call([]reflect.Value{ key, m.MapIndex(key) })
																	}
																	ok = true
												}
											}
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

	default:								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
												switch f.Type().NumIn() {
												case 1:				for {
																		if v, open := c.Recv(); open {
																			f.Call([]reflect.Value{ v })
																		} else {
																			break
																		}
																	}
																	ok = true

												case 2:				for count := 0; ; count++ {
																		if v, open := c.Recv(); open {
																			f.Call([]reflect.Value{ reflect.ValueOf(count), v })
																		} else {
																			break
																		}
																	}
																	ok = true
												}
											}
	}
	return
}

func rangeGenericGenerator(g, f reflect.Value) (ok bool) {
	switch tg := g.Type(); tg.NumIn() {
	case 0:			if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
						switch tf := f.Type(); tf.NumIn() {
						case 1:			if tf.In(0) == tg.Out(0) {
											for v := g.Call([]reflect.Value{}); !v[1].Bool(); v = g.Call([]reflect.Value{}) {
												f.Call([]reflect.Value{v[0]})
											}
											ok = true
										}

						case 2:			if tf.In(1) == tg.Out(0) {
											i := 0
											for v := g.Call([]reflect.Value{}); !v[1].Bool(); v = g.Call([]reflect.Value{}) {
												f.Call([]reflect.Value{reflect.ValueOf(i), v[0]})
												i++
											}
											ok = true
										}
						}
					}

	case 1:			if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
						switch tf := f.Type(); tf.NumIn() {
						case 1:			if tf.In(0) == tg.Out(0) {
											i := 0
											for v := g.Call([]reflect.Value{reflect.ValueOf(i)}); !v[1].Bool(); v = g.Call([]reflect.Value{reflect.ValueOf(i)}) {
												f.Call([]reflect.Value{v[0]})
												i++
											}
											ok = true
										}


						case 2:			if tf.In(1) == tg.Out(0) {
											i := 0
											for v := g.Call([]reflect.Value{reflect.ValueOf(i)}); !v[1].Bool(); v = g.Call([]reflect.Value{reflect.ValueOf(i)}) {
												f.Call([]reflect.Value{reflect.ValueOf(i), v[0]})
												i++
											}
											ok = true
										}
						}
					}			
	}
	return
}

/*
	A Generator is a function which when passed an index returns a resulting value generated from it, along with a boolean flag indicating
	whether or not the generator has completed its work.
*/
func rangeGenerator(g reflect.Value, f interface{}) (ok bool) {
	if tg := g.Type(); tg.NumOut() == 2 {
		switch tg.NumIn() {
		case 0:			switch f := f.(type) {
						case func(interface{}):					for v := g.Call([]reflect.Value{}); !v[1].Bool(); v = g.Call([]reflect.Value{}) {
																	f(v[0].Interface())
																}
																ok = true

						case func(int, interface{}):			for count, v := 0, g.Call([]reflect.Value{}); !v[1].Bool(); v = g.Call([]reflect.Value{}) {
																	f(count, v[0].Interface())
																	count++
																} 
																ok = true

						case func(interface{}, interface{}):	for count, v := 0, g.Call([]reflect.Value{}); !v[1].Bool(); v = g.Call([]reflect.Value{}) {
																	f(count, v[0].Interface())
																	count++
																}
																ok = true

						default:								rangeGenericGenerator(g, reflect.ValueOf(f))
						}

		case 1:			count := 0
						switch f := f.(type) {
						case func(interface{}):					for v := g.Call(valueslice(count)); !v[1].Bool(); v = g.Call(valueslice(count)) {
																	f(v[0].Interface())
																	count++
																}
																ok = true

						case func(int, interface{}):			for v := g.Call(valueslice(count)); !v[1].Bool(); v = g.Call(valueslice(count)) {
																	f(count, v[0].Interface())
																	count++
																}
																ok = true

						case func(interface{}, interface{}):	for v := g.Call(valueslice(count))[0]; !v.IsNil(); v = g.Call(valueslice(count))[0] {
																	f(count, v.Interface())
																	count++
																}
																ok = true
						}
		}
	}
	return
}

func Each(container, f interface{}) (ok bool) {
	switch container := container.(type) {
	case Iterable:			ok = container.Each(f)

	case IndexedReader:		ok = eachIndexedReader(container, f)

	case MappedReader:		ok = eachMappedReader(container, f)

	default:				switch c := reflect.ValueOf(container); c.Kind() {
							case reflect.Slice:		ok = rangeSlice(c, f)

							case reflect.Map:		ok = rangeMap(c, f)

							case reflect.Chan:		ok = rangeChannel(c, f)

							case reflect.Func:		ok = rangeGenerator(c, f)
							}
	}
	return
}

func Cycle(container interface{}, count int, f func(interface{})) (i int, ok bool) {
	switch {
	case count == 0:	for ok = true; ok; i++ { ok = Each(container, f) }

	default:			for ok = true; i < count && ok; i++ { ok = Each(container, f) }
	}
	return
}