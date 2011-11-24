package sexp

import "github.com/feyeleanor/raw"
import "reflect"

/*
func Merge(d, s Mapping) {
	Each(s.Keys(), func(k interface{}) {
		d.Store(k, s.At(k))
	})
}
*/

/*
	Combine takes two parameters and applies a function to them.
	In the case of scalar parameters the function is applied once and the result returned.
	In the case of vector parameters the function is applied to each matched pair of elements and the result returned.
	A vector result will always have the same type as the left-hand parameter and the number of elements in the result
	vector will be the same as the number of elements in the lef-hand parameter with trailing unmatched values combined
	with the correct zero value for the contained type.
*/
func Combine(left, right interface{}, f func(interface{}, interface{}) interface{}) (r interface{}) {
	switch left := left.(type) {
	case Combinable:	r = left.Combine(right, f)
	case Indexable:		r = combineIndexable(left, right, f)
	default:			r = combineValue(left, right, f)
	}
	return
}

func blank(container interface{}) reflect.Value {
	return reflect.Zero(reflect.ValueOf(container).Type().Elem())
}

func combineIndexable(left Indexable, right interface{}, f func(interface{}, interface{}) interface{}) (result interface{}) {
	switch right := right.(type) {
	case Indexable:		makeSlice := func(length int) (r Indexable) {
							raw.CatchAll(func() {
								r = Reallocate(left, length, length).(Indexable)
							})
							return 
						}

						var s Indexable
						switch l, r := left.Len(), right.Len(); {
						case l == r:		if s := makeSlice(l); result != nil {
												for i := 0; i < l; i++ {
													s.Set(i, f(left.At(i), right.At(i)))
												}
											}

						case l > r:			if s := makeSlice(l); result != nil {
												for i := 0; i < r; i++ {
													s.Set(i, f(left.At(i), right.At(i)))
												}

												blank := reflect.Zero(reflect.ValueOf(left).Type().Elem()).Interface()
												for i := r; i < l; i++ {
													s.Set(i, f(left.At(i), blank))
												}
											}

						case l < r:			if s := makeSlice(r); result != nil {
												for i := 0; i < l; i++ {
													s.Set(i, f(left.At(i), right.At(i)))
												}

												blank := reflect.Zero(reflect.ValueOf(left).Type().Elem()).Interface()
												for i := l; i < r; i++ {
													s.Set(i, f(blank, right.At(i)))
												}
											}
						}
						result = s

	default:			switch right := reflect.ValueOf(right); right.Kind() {
						case reflect.Slice:		makeSlice := func(length int) (r reflect.Value) {
													raw.CatchAll(func() {
														Reallocate(left, length, length)
														r = reflect.ValueOf(left)
													})
													return 
												}

												var s reflect.Value
												CombineAndSet := func(i int, l interface{}, r reflect.Value) {
													s.Index(i).Set(reflect.ValueOf(f(l, r.Interface())))
												}

												switch l, r := left.Len(), right.Len(); {
												case l == r:		if s = makeSlice(l); s.IsValid() {
																		for i := 0; i < l; i++ {
																			CombineAndSet(i, left.At(i), right.Index(i))
																		}
																	}

												case l > r:			if s = makeSlice(l); s.IsValid() {
																		for i := 0; i < r; i++ {
																			CombineAndSet(i, left.At(i), right.Index(i))
																		}
																		for i := r; i < l; i++ {
																			CombineAndSet(i, left.At(i), reflect.Value{})
																		}
																	}

												case l < r:			if s = makeSlice(r); s.IsValid() {
																		for i := 0; i < l; i++ {
																			CombineAndSet(i, left.At(i), right.Index(i))
																		}
																		for i := l; i < r; i++ {
																			CombineAndSet(i, nil, right.Index(i))
																		}
																	}
												}
												result = s.Interface()

						case reflect.Map:		m := reflect.MakeMap(right.Type())
												CombineAndSet := func(i, l interface{}, r reflect.Value) {
													m.SetMapIndex(reflect.ValueOf(i), reflect.ValueOf(f(l, r.Interface())))
												}

												for i := left.Len() - 1; i > 0; i-- {
													CombineAndSet(i, left.At(i), right.MapIndex(reflect.ValueOf(i)))
												}

												for _, k := range right.MapKeys() {
													i := int(k.Int())
													if left.At(i) == nil {
														CombineAndSet(i, left.At(i), right.MapIndex(k))
													}
												}
												result = m.Interface()
						}
	}
	return
}

func combineValue(Left, Right interface{}, f func(interface{}, interface{}) interface{}) (result interface{}) {
	left := reflect.ValueOf(Left)
	blank := reflect.Zero(left.Type().Elem())
	switch left.Kind() {
	case reflect.Slice:		switch right := Right.(type) {
							case Indexable:			makeSlice := func(length int) (r reflect.Value) {
														raw.CatchAll(func() {
															Reallocate(left, length, length)
															r = reflect.ValueOf(left)
															})
														return 
													}

													var s reflect.Value
													CombineAndSet := func(i int, l reflect.Value, r interface{}) {
														s.Index(i).Set(reflect.ValueOf(f(l.Interface(), r)))
													}

													switch l, r := left.Len(), right.Len(); {
													case l == r:		if s = makeSlice(l); s.IsValid() {
																			for i := 0; i < l; i++ {
																				CombineAndSet(i, left.Index(i), right.At(i))
																			}
																		}

													case l > r:			if s = makeSlice(l); s.IsValid() {
																			for i := 0; i < r; i++ {
																				CombineAndSet(i, left.Index(i), right.At(i))
																			}
																			for i := r; i < l; i++ {
																				CombineAndSet(i, left.Index(i), blank)
																			}
																		}

													case l < r:			if s = makeSlice(r); s.IsValid() {
																			for i := 0; i < l; i++ {
																				CombineAndSet(i, left.Index(i), right.At(i))
																			}
																			for i := l; i < r; i++ {
																				CombineAndSet(i, blank, right.At(i))
																			}
																		}
													}
													result = s.Interface()

							default:				switch right := reflect.ValueOf(right); right.Kind() {
													case reflect.Slice:		makeSlice := func(length int) reflect.Value {
																				return reflect.MakeSlice(left.Type(), length, length)
																			}
																			
																			var s reflect.Value
																			CombineAndSet := func(i int, l, r reflect.Value) {
																				s.Index(i).Set(reflect.ValueOf(f(l.Interface(), r.Interface())))
																			}

																			switch l, r := left.Len(), right.Len(); {
																			case l == r:		if s = makeSlice(l); s.IsValid() {
																									for i := 0; i < l; i++ {
																										CombineAndSet(i, left.Index(i), right.Index(i))
																									}
																								}

																			case l > r:			if s = makeSlice(l); s.IsValid() {
																									for i := 0; i < r; i++ {
																										CombineAndSet(i, left.Index(i), right.Index(i))
																									}
																									for i := r; i < l; i++ {
																										CombineAndSet(i, left.Index(i), blank)
																									}
																								}

																			case l < r:			if s = makeSlice(r); s.IsValid() {
																									for i := 0; i < l; i++ {
																										CombineAndSet(i, left.Index(i), right.Index(i))
																									}
																									for i := l; i < r; i++ {
																										CombineAndSet(i, blank, right.Index(i))
																									}
																								}
																			}
																			result = s.Interface()

													case reflect.Map:		if map_type := right.Type(); map_type.Key().Kind() == reflect.Int {
																				n := reflect.MakeMap(map_type)
																				for i := 0; i < left.Len(); i++ {
																					n.SetMapIndex(reflect.ValueOf(i), left.Index(i))
																				}
																				result = combineValue(n.Interface(), Right, f)
																			}
													}
							}
	case reflect.Map:		switch right := Right.(type) {
/*							case Map:				m := reflect.MakeMap(left.Type())
													CombineAndSet := func(k interface{}, l reflect.Value, r interface{}) {
														m.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(f(l.Interface(), r)))
													}

													right.EachWithKey(func(k, r interface{}) {
														if l := left.MapIndex(reflect.ValueOf(k)); l.IsValid() {
															CombineAndSet(k, l, r)
														}
													})
*/
							case Indexable:			m := reflect.MakeMap(left.Type())
													CombineAndSet := func(i int, l reflect.Value, r interface{}) {
														m.SetMapIndex(reflect.ValueOf(i), reflect.ValueOf(f(l.Interface(), r)))
													}

													for i := left.Len() - 1; i > 0; i-- {
														CombineAndSet(i, left.MapIndex(reflect.ValueOf(i)), right.At(i))
													}

													for i := right.Len() - 1; i > 0; i-- {
														k := reflect.ValueOf(i)
														if !m.MapIndex(k).IsValid() {
															CombineAndSet(i, left.MapIndex(k), right.At(i))
														}
													}
													result = m.Interface()

							default:				switch right := reflect.ValueOf(right); right.Kind() {
													case reflect.Map:		m := reflect.MakeMap(left.Type())
																			CombineAndSet := func(k reflect.Value) {
																				lv := left.MapIndex(k)
																				rv := right.MapIndex(k)

																				var x interface{}
																				if lv.IsValid() {
																					if rv.IsValid() {
																						x = f(lv.Interface(), rv.Interface())																						
																					} else {
																						x = f(lv.Interface(), blank.Interface())
																					}
																				} else {
																					x = f(blank.Interface(), rv.Interface())
																				}
																				m.SetMapIndex(k, reflect.ValueOf(x))
																			}

																			for _, k := range left.MapKeys() {
																				CombineAndSet(k)
																			}

																			for _, k := range right.MapKeys() {
																				CombineAndSet(k)
																			}
																			result = m.Interface()

													case reflect.Slice:		if map_type := left.Type(); map_type.Key().Kind() == reflect.Int {
																				n := reflect.MakeMap(map_type)
																				for i := 0; i < right.Len(); i++ {
																					n.SetMapIndex(reflect.ValueOf(i), right.Index(i))
																				}
																				result = combineValue(Left, n.Interface(), f)
																			}
													}
							}
	}
	return
}