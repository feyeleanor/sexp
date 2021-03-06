package sexp

import R "reflect"

type Iterable interface {
	Each(function interface{}) bool
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

	case func(...interface{}):				p := make([]interface{}, end, end)
											for i := 0; i < end; i++ {
												p[i] = container.At(i)
											}
											f(p...)
											ok = true

	case func(R.Value):						for i := 0; i < end; i++ {
												f(R.ValueOf(container.At(i)))
											}
											ok = true

	case func(int, R.Value):				for i := 0; i < end; i++ {
												f(i, R.ValueOf(container.At(i)))
											}
											ok = true

	case func(interface{}, R.Value):		for i := 0; i < end; i++ {
												f(i, R.ValueOf(container.At(i)))
											}
											ok = true

	case func(R.Value, R.Value):			for i := 0; i < end; i++ {
												f(R.ValueOf(i), R.ValueOf(container.At(i)))
											}
											ok = true


	case func(...R.Value):					p := make([]R.Value, end, end)
											for i := 0; i < end; i++ {
												p[i] = R.ValueOf(container.At(i))
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													p := make([]R.Value, end, end)
													for i := 0; i < end; i++ {
														p[i] = R.ValueOf(container.At(i))
													}
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		for i := 0; i < end; i++ {
																			p[0] = R.ValueOf(container.At(i))
																			f.Call(p)
																		}
																		ok = true

													case 2:				//	f(i, v)
																		p := make([]R.Value, 2, 2)
																		for i := 0; i < end; i++ {
																			p[0], p[1] = R.ValueOf(i), R.ValueOf(container.At(i))
																			f.Call(p)
																		}
																		ok = true
													}
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

	case func(...interface{}):				keys := container.Keys()
											p := make([]interface{}, len(keys), len(keys))
											for i, v := range keys {
												p[i] = container.At(v)
											}
											f(p...)
											ok = true

	case func(R.Value):						for _, v := range container.Keys() {
												f(R.ValueOf(container.At(v)))
											}
											ok = true

	case func(interface{}, R.Value):		for _, v := range container.Keys() {
												f(v, R.ValueOf(container.At(v)))
											}
											ok = true

	case func(R.Value, R.Value):			for _, v := range container.Keys() {
												f(R.ValueOf(v), R.ValueOf(container.At(v)))
											}
											ok = true

	case func(...R.Value):					keys := container.Keys()
											p := make([]R.Value, len(keys), len(keys))
											for i, v := range keys {
												p[i] = R.ValueOf(container.At(v))
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													// f(...v)
													keys := container.Keys()
													p := make([]R.Value, len(keys), len(keys))
													for i, v := range keys {
														p[i] = R.ValueOf(container.At(v))
													}
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		for _, v := range container.Keys() {
																			p[0] = R.ValueOf(container.At(v))
																			f.Call(p)
																		}
																		ok = true

													case 2:				// f(i, v)
																		p := make([]R.Value, 2, 2)
																		for _, v := range container.Keys() {
																			p[0], p[1] = R.ValueOf(v), R.ValueOf(container.At(v))
																			f.Call(p)
																		}
																		ok = true
													}
												}
											}
	}
	return
}

func rangeSlice(s R.Value, f interface{}) (ok bool) {
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

	case func(...interface{}):				p := make([]interface{}, end, end)
											for i := 0; i < end; i++ {
												p[i] = s.Index(i).Interface()
											}
											f(p...)
											ok = true

	case func(R.Value):						for i := 0; i < end; i++ {
												f(s.Index(i))
											}
											ok = true

	case func(int, R.Value):				for i := 0; i < end; i++ {
												f(i, s.Index(i))
											}
											ok = true

	case func(interface{}, R.Value):		for i := 0; i < end; i++ {
												f(i, s.Index(i))
											}
											ok = true

	case func(R.Value, R.Value):			for i := 0; i < end; i++ {
												f(R.ValueOf(i), s.Index(i))
											}
											ok = true

	case func(...R.Value):					p := make([]R.Value, end, end)
											for i := 0; i < end; i++ {
												p[i] = s.Index(i)
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												end := s.Len()
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													p := make([]R.Value, end, end)
													for i := 0; i < end; i++ {
														p[i] = s.Index(i)
													}
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		for i := 0; i < end; i++ {
																			p[0] = s.Index(i)
																			f.Call(p)
																		}
																		ok = true

													case 2:				//	f(i, v)
																		p := make([]R.Value, 2, 2)
																		for i := 0; i < end; i++ {
																			p[0], p[1] = R.ValueOf(i), s.Index(i)
																			f.Call(p)
																		}
																		ok = true
													}
												}
											}
	}
	return
}

func rangeMap(m R.Value, f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}): for iter := m.MapRange(); iter.Next(); {
												f(iter.Value().Interface())
											}
											ok = true

	case func(interface{}, interface{}):	for iter := m.MapRange(); iter.Next(); {
												f(iter.Key().Interface(), iter.Value().Interface())
											}
											ok = true

	case func(...interface{}):				p := make([]interface{}, m.Len(), m.Len())
                      for i, iter := 0, m.MapRange(); iter.Next(); i++ {
												p[i] = iter.Value().Interface()
											}
											f(p...)
											ok = true

	case func(R.Value):	for iter := m.MapRange(); iter.Next(); {
												f(iter.Value())
											}
											ok = true

	case func(interface{}, R.Value):  for iter := m.MapRange(); iter.Next(); {
												f(iter.Key().Interface(), iter.Value())
											}
											ok = true

	case func(R.Value, R.Value): for iter := m.MapRange(); iter.Next(); {
												f(iter.Key(), iter.Value())
											}
											ok = true

	case func(...R.Value):					p := make([]R.Value, m.Len(), m.Len())
                      for i, iter := 0, m.MapRange(); iter.Next(); i++ {
												p[i] = iter.Value()
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													p := make([]R.Value, m.Len(), m.Len())
                          for i, iter := 0, m.MapRange(); iter.Next(); i++ {
														p[i] = iter.Value()
													}
													f.Call(p)
													ok = true
												} else {
													switch t.NumIn() {
													case 1:			//	f(v)
																	p := make([]R.Value, 1, 1)
                                  for iter := m.MapRange(); iter.Next(); {
																		p[0] = iter.Value()
																		f.Call(p)
																	}
																	ok = true

													case 2:			//	f(i, v)
																	p := make([]R.Value, 2, 2)
                                  for iter := m.MapRange(); iter.Next(); {
																		p[0], p[1] = iter.Key(), iter.Value()
																		f.Call(p)
																	}
																	ok = true
													}
												}
											}
	}
	return
}

func rangeChannel(c R.Value, f interface{}) (ok bool) {
	switch f := f.(type) {
	case func(interface{}):					for v, open := c.Recv(); open; {
												f(v.Interface())
												v, open = c.Recv()
											}
											ok = true

	case func(int, interface{}):			i := 0
											for v, open := c.Recv(); open; i++ {
												f(i, v.Interface())
												v, open = c.Recv()
											}
											ok = true

	case func(interface{}, interface{}):	i := 0
											for v, open := c.Recv(); open; i++ {
												f(i, v.Interface())
												v, open = c.Recv()
											}
											ok = true

	case func(...interface{}):				p := make([]interface{}, 0, 4)
											for v, open := c.Recv(); open; {
												p = append(p, v.Interface())
												v, open = c.Recv()
											}
											f(p...)
											ok = true

	case func(R.Value):						for v, open := c.Recv(); open; {
												f(v)
												v, open = c.Recv()
											}
											ok = true

	case func(int, R.Value):				i := 0
											for v, open := c.Recv(); open; i++ {
												f(i, v)
												v, open = c.Recv()
											}
											ok = true

	case func(interface{}, R.Value):		i := 0
											for v, open := c.Recv(); open; i++ {
												f(i, v)
												v, open = c.Recv()
											}
											ok = true

	case func(...R.Value):					p := make([]R.Value, 0, 4)
											for v, open := c.Recv(); open; {
												p = append(p, v)
												v, open = c.Recv()
											}
											f(p...)
											ok = true

	default:								if f := R.ValueOf(f); f.Kind() == R.Func {
												if t := f.Type(); t.IsVariadic() {
													//	f(...v)
													p := make([]R.Value, 0, 4)
													for v, open := c.Recv(); open; {
														p = append(p, v)
														v, open = c.Recv()
													}
													f.Call(p)
												} else {
													switch t.NumIn() {
													case 1:				//	f(v)
																		p := make([]R.Value, 1, 1)
																		for v, open := c.Recv(); open; {
																			p[0] = v
																			f.Call(p)
																			v, open = c.Recv()
																		}
																		ok = true

													case 2:				//	f(i, v)
																		p := make([]R.Value, 2, 2)
																		i := 0
																		for v, open := c.Recv(); open; i++ {
																			p[0], p[1] = R.ValueOf(i), v
																			f.Call(p)
																			v, open = c.Recv()
																		}
																		ok = true
													}
												}
											}
	}
	return
}

func rangeGenericFunction(g, f R.Value) (ok bool) {
	switch tg := g.Type(); tg.NumIn() {
	case 0:			if f := R.ValueOf(f); f.Kind() == R.Func {
						if tf := f.Type(); tf.IsVariadic() {
							//	f(...v)
							i := 0
							pg := []R.Value{}
							pf := make([]R.Value, 0, 4)
							for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
								pf = append(pf, v[0])
								i++
							}
							f.Call(pf)
							ok = true
						} else {
							switch tf.NumIn() {
							case 1:		//	f(v)
										pg := []R.Value{}
										pf := make([]R.Value, 1, 1)
										for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
											pf[0] = v[0]
											f.Call(pf)
										}
										ok = true

							case 2:		//	f(i, v)
										i := 0
										pg := []R.Value{}
										pf := make([]R.Value, 2, 2)
										for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
											pf[0], pf[1] = R.ValueOf(i), v[0]
											f.Call(pf)
											i++
										}
										ok = true
							}
						}
					}

	case 1:			if f := R.ValueOf(f); f.Kind() == R.Func {
						if tf := f.Type(); tf.IsVariadic() {
							//	f(...v)
							i := 0
							pg := []R.Value{ R.ValueOf(0) }
							pf := make([]R.Value, 0, 4)
							for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
								pf = append(pf, v[0])
								i++
								pg[0] = R.ValueOf(i)
							}
							f.Call(pf)
							ok = true
						} else {
							switch tf.NumIn() {
							case 1:		//	f(v)
										i := 0
										p := []R.Value{ R.ValueOf(0) }
										for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
											p[0] = v[0]
											f.Call(p)
											i++
											p[0] = R.ValueOf(i)
										}
										ok = true

							case 2:		//	f(i, v)
										i := 0
										pg := []R.Value{ R.ValueOf(0) }
										pf := make([]R.Value, 2, 2)
										for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
											pf[0], pf[1] = pg[0], v[0]
											f.Call(pf)
											i++
											pg[0] = R.ValueOf(i)
										}
										ok = true
							}
						}
					}
	}
	return
}

func rangeFunction(g R.Value, f interface{}) (ok bool) {
	if tg := g.Type(); tg.NumOut() == 2 {
		switch tg.NumIn() {
		case 0:			switch f := f.(type) {
						case func(interface{}):					p := []R.Value{}
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0].Interface())
																}
																ok = true

						case func(int, interface{}):			p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																}
																ok = true

						case func(interface{}, interface{}):	p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																}
																ok = true

						case func(...interface{}):				pg := []R.Value{}
																pf := make([]interface{}, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0].Interface())
																}
																f(pf...)
																ok = true

						case func(R.Value):						p := []R.Value{}
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0])
																}
																ok = true

						case func(int, R.Value):				p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																}
																ok = true

						case func(interface{}, R.Value):		p := []R.Value{}
																for i, v := 0, g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																}
																ok = true

						case func(...R.Value):					pg := []R.Value{}
																pf := make([]R.Value, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0])
																}
																f(pf...)
																ok = true

						default:								rangeGenericFunction(g, R.ValueOf(f))
						}

		case 1:			switch f := f.(type) {
						case func(interface{}):					i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0].Interface())
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(int, interface{}):			i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(interface{}, interface{}):	i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[0].IsNil(); v = g.Call(p) {
																	f(i, v[0].Interface())
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(...interface{}):				i := 0
																pg := []R.Value{ R.ValueOf(0) }
																pf := make([]interface{}, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0].Interface())
																	i++
																	pg[i] = R.ValueOf(i)
																}
																f(pf...)
																ok = true

						case func(R.Value):						i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(v[0])
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(int, R.Value):				i := 0
																p := []R.Value{ R.ValueOf(0) }
																for v := g.Call(p); !v[1].Bool(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(interface{}, R.Value):		i := 0
																p := []R.Value{ R.ValueOf(i) }
																for v := g.Call(p); !v[0].IsNil(); v = g.Call(p) {
																	f(i, v[0])
																	i++
																	p[0] = R.ValueOf(i)
																}
																ok = true

						case func(...R.Value):					i := 0
																pg := []R.Value{ R.ValueOf(i) }
																pf := make([]R.Value, 0, 4)
																for v := g.Call(pg); !v[1].Bool(); v = g.Call(pg) {
																	pf = append(pf, v[0])
																	i++
																	pf[i] = R.ValueOf(i)
																}
																f(pf...)
																ok = true

						default:								rangeGenericFunction(g, R.ValueOf(f))
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

	default:				switch c := R.ValueOf(container); c.Kind() {
							case R.Slice:		ok = rangeSlice(c, f)

							case R.Map:			ok = rangeMap(c, f)

							case R.Chan:		ok = rangeChannel(c, f)

							case R.Func:		ok = rangeFunction(c, f)
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
