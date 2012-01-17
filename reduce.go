package sexp

import "reflect"

type Reducible interface {
	Reduce(seed, function interface{}) (r interface{}, ok bool)
}

func reduceIndexedReader(c IndexedReader, seed, f interface{}) (r interface{}, ok bool) {
	end := c.Len()
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:	for i := 0; i < end; i++ {
															seed = f(seed, c.At(i))
														}
														ok = true

	default:											if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
															switch f.Type().NumIn() {
															case 2:				for i := 0; i < end; i++ {
																					f.Call(valueslice(i, c.At(i)))
																				}
																				ok = true
															}
														}
	}
	return
}

func reduceMappedReader(c MappedReader, seed, function interface{}) (r interface{}, ok bool) {
	return
}

func reduceIterable(container Iterable, seed, f interface{}) (r interface{}, ok bool) {
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:		r = seed
															Each(container, func(x interface{}) {
																r = f(r, x)
															})
															ok = true
	}
	return
}

func reduceSlice(s reflect.Value, seed, f interface{}) (r interface{}, ok bool) {
	end := s.Len()
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:	v := reflect.New(s.Type().Elem()).Elem()
														v.Set(reflect.ValueOf(seed))
														for i := 0; i < end; i++ {
															v = reflect.ValueOf(f(v.Interface(), s.Index(i).Interface()))
														}
														r = v.Interface()
														ok = true

	default:											if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
															v := reflect.New(s.Type().Elem()).Elem()
															v.Set(reflect.ValueOf(seed))
															switch f.Type().NumIn() {
															case 2:				for i, end := 0, s.Len(); i < end; i++ {
																					v = f.Call([]reflect.Value{ v, s.Index(i) })[0]
																				}
																				ok = true
															}
															if ok {
																r = v.Interface()
															}
														}
	}
	return
}

func reduceMap(m reflect.Value, seed, f interface{}) (r interface{}, ok bool) {
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:	v := reflect.New(m.Type().Elem()).Elem()
														v.Set(reflect.ValueOf(seed))
														for _, key := range m.MapKeys() {
															v = reflect.ValueOf(f(v.Interface(), m.MapIndex(key).Interface()))
														}
														r = v.Interface()
														ok = true

	default:											if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
															v := reflect.New(m.Type().Elem()).Elem()
															v.Set(reflect.ValueOf(seed))
															switch f.Type().NumIn() {
															case 2:				for _, key := range m.MapKeys() {
																					v = f.Call([]reflect.Value{ v, m.MapIndex(key) })[0]
																				}
																				ok = true
															}
															if ok {
																r = v.Interface()
															}
														}
	}
	return
}

func reduceChan(c reflect.Value, seed, f interface{}) (r interface{}, ok bool) {
	switch f := f.(type) {
	case func(interface{}, interface{}) interface{}:	v := reflect.New(c.Type().Elem()).Elem()
														v.Set(reflect.ValueOf(seed))
														for {
															if x, open := c.Recv(); open {
																v = reflect.ValueOf(f(v.Interface(), x))
															} else {
																break
															}
														}
														r = v.Interface()
														ok = true

	default:											if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
															v := reflect.New(c.Type().Elem()).Elem()
															v.Set(reflect.ValueOf(seed))
															switch f.Type().NumIn() {
															case 2:				for {
																					if x, open := c.Recv(); open {
																						v = f.Call([]reflect.Value{ v, x })[0]
																					} else {
																						break
																					}
																				}
																				ok = true
															}
															if ok {
																r = v.Interface()
															}
														}
	}
	return
}

func reduceFunction(g reflect.Value, seed, f interface{}) (r interface{}, ok bool) {
	if t := g.Type(); t.NumOut() == 2 {
		switch t.NumIn() {
		case 0:			switch f := f.(type) {
						case func(interface{}) interface{}:		v := reflect.New(g.Type().Out(0)).Elem()
																v.Set(reflect.ValueOf(seed))
																for x := g.Call([]reflect.Value{}); !x[1].Bool(); x = g.Call([]reflect.Value{}) {
																	v = reflect.ValueOf(f(x[0].Interface()))
																}
																r = v.Interface()
																ok = true
						}

		case 1:			count := 0
						switch f := f.(type) {
						case func(interface{}) interface{}:		v := reflect.New(g.Type().Out(0)).Elem()
																v.Set(reflect.ValueOf(seed))
																for x := g.Call(valueslice(count)); !x[1].Bool(); x = g.Call(valueslice(count)) {
																	v = reflect.ValueOf(f(x[0].Interface()))
																	count++
																}
																r = v.Interface()
																ok = true
						}
		}
	}
	return
}

func reduce(container, seed, f interface{}) (r interface{}, ok bool) {
	switch c := reflect.ValueOf(container); c.Kind() {
	case reflect.Invalid:	r, ok = seed, true

	case reflect.Slice:		r, ok = reduceSlice(c, seed, f)
							
	case reflect.Map:		r, ok = reduceMap(c, seed, f)

	case reflect.Chan:		r, ok = reduceChan(c, seed, f)

	case reflect.Func:		r, ok = reduceFunction(c, seed, f)
	}
	return
}

func Reduce(container, seed interface{}, f interface{}) (r interface{}, ok bool) {
	switch c := container.(type) {
	case Reducible:				println("reducible")
								r, ok = c.Reduce(seed, f)

	case IndexedReader:			println("IndexedReader")
								r, ok = reduceIndexedReader(c, seed, f)

	case MappedReader:			println("MappedReader")
								r, ok = reduceMappedReader(c, seed, f)

	case Iterable:				println("Iterable")
								r, ok = reduceIterable(c, seed, f)

	default:					switch c := reflect.ValueOf(container); c.Kind() {
								case reflect.Invalid:	r, ok = seed, true

								case reflect.Slice:		r, ok = reduceSlice(c, seed, f)

								case reflect.Map:		r, ok = reduceMap(c, seed, f)
								}
	}
	return
}