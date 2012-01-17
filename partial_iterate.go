package sexp

import R "reflect"

type PartiallyIterable interface {
	While(condition bool, function interface{}) (count int)
}

func whileIndexedReader(container IndexedReader, r bool, f interface{}) (count int) {
	if end := container.Len(); end > 0 {
		switch f := f.(type) {
		case func(interface{}) bool:				if f(container.At(0)) == r {
														count = 1
														for i := 1; i < end; i++ {
															if f(container.At(i)) != r {
																break
															}
															count++
														}
													}

		case func(int, interface{}) bool:			if f(0, container.At(0)) == r {
														count = 1
														for i := 1; i < end; i++ {
															if f(i, container.At(i)) != r {
																break
															}
															count++
														}
													}

		case func(interface{}, interface{}) bool:	if f(0, container.At(0)) == r {
														count = 1
														for i := 1; i < end; i++ {
															if f(i, container.At(i)) != r {
																break
															}
															count++
														}
													}

		case func(R.Value) bool:					if f(R.ValueOf(container.At(0))) == r {
														count = 1
														for i := 1; i < end; i++ {
															if f(R.ValueOf(container.At(i))) != r {
																break
															}
															count++
														}
													}

		case func(int, R.Value) bool:				if f(0, R.ValueOf(container.At(0))) == r {
														count = 1
														for i := 1; i < end; i++ {
															if f(i, R.ValueOf(container.At(i))) != r {
																break
															}
															count++
														}
													}

		case func(R.Value, R.Value) bool:			if f(R.ValueOf(0), R.ValueOf(container.At(0))) == r {
														count = 1
														for i := 1; i < end; i++ {
															if f(R.ValueOf(i), R.ValueOf(container.At(i))) != r {
																break
															}
															count++
														}
													}

		default:									if f := R.ValueOf(f); f.Kind() == R.Func {
														switch f.Type().NumIn() {
														case 1:				p := valueslice(container.At(0))
																			if f.Call(p)[0].Bool() == r {
																				count = 1
																				for i := 1; i < end; i++ {
																					p[0] = R.ValueOf(container.At(i))
																					if f.Call(p)[0].Bool() != r {
																						break
																					}
																					count++
																				}
																			}

														case 2:				p := valueslice(0, container.At(0))
																			if f.Call(p)[0].Bool() == r {
																				count = 1
																				for i := 1; i < end; i++ {
																					p[0] = R.ValueOf(i)
																					p[1] = R.ValueOf(container.At(i))
																					if f.Call(p)[0].Bool() != r {
																						break
																					}
																					count++
																				}
																			}
														}
													}
		}
	}
	return
}

func whileSlice(s R.Value, r bool, f interface{}) (count int) {
	if end := s.Len(); end > 0 {
		switch f := f.(type) {
		case func(interface{}) bool:					if f(s.Index(0).Interface()) == r {
															count = 1
															for i := 1; i < end; i++ {
																if f(s.Index(i).Interface()) != r {
																	break
																}
																count++
															}
														}

		case func(int, interface{}) bool:				if f(0, s.Index(0).Interface()) == r {
															count = 1
															for i := 1; i < end; i++ {
																if f(i, s.Index(i).Interface()) != r {
																	break
																}
																count++
															}
														}

		case func(interface{}, interface{}) bool:		if f(0, s.Index(0).Interface()) == r {
															count = 1
															for i := 1; i < end; i++ {
																if f(i, s.Index(i).Interface()) != r {
																	break
																}
																count++
															}
														}

		case func(R.Value) bool:						if f(s.Index(0)) == r {
															count = 1
															for i := 1; i < end; i++ {
																if f(s.Index(i)) != r {
																	break
																}
																count++
															}
														}

		case func(int, R.Value) bool:					if f(0, s.Index(0)) == r {
															count = 1
															for i := 1; i < end; i++ {
																if f(i, s.Index(i)) != r {
																	break
																}
																count++
															}
														}

		case func(interface{}, R.Value) bool:			if f(0, s.Index(0)) == r {
															count = 1
															for i := 1; i < end; i++ {
																if f(i, s.Index(i)) != r {
																	break
																}
																count++
															}
														}


		case func(R.Value, R.Value) bool:				if f(R.ValueOf(0), s.Index(0)) == r {
															count = 1
															for i := 1; i < end; i++ {
																if f(R.ValueOf(i), s.Index(i)) != r {
																	break
																}
																count++
															}
														}


		default:										if f := R.ValueOf(f); f.Kind() == R.Func {
															switch f.Type().NumIn() {
															case 1:				p := []R.Value{ s.Index(0) }
																				if f.Call(p)[0].Bool() == r {
																					count = 1
																					for i := 1; i < end; i++ {
																						p[0] = s.Index(i)
																						if f.Call(p)[0].Bool() != r {
																							break
																						}
																						count++
																					}
																				}

															case 2:				p := []R.Value{ R.ValueOf(0), s.Index(0) }
																				if f.Call(p)[0].Bool() == r {
																					count = 1
																					for i := 1; i < end; i++ {
																						p[0] = R.ValueOf(i)
																						p[1] = s.Index(i)
																						if f.Call(p)[0].Bool() != r {
																							break
																						}
																						count++
																					}
																				}
															}
														}
		}
	}
	return
}

func whileChannel(c R.Value, r bool, f interface{}) (count int) {
	switch f := f.(type) {
	case func(interface{}) bool:					for ; ; count++ {
														if v, open := c.Recv(); !open || f(v.Interface()) != r {
															break
														}
													}

	case func(int, interface{}) bool:				for ; ; count++ {
														if v, open := c.Recv(); !open || f(count, v.Interface()) != r {
															break
														}
													}

	case func(interface{}, interface{}) bool:		for ; ; count++ {
														if v, open := c.Recv(); !open || f(count, v.Interface()) != r {
															break
														}
													}

	case func(R.Value) bool:						for ; ; count++ {
														if v, open := c.Recv(); !open || f(v) != r {
															break
														}
													}

	case func(int, R.Value) bool:					for ; ; count++ {
														if v, open := c.Recv(); !open || f(count, v) != r {
															break
														}
													}

	case func(interface{}, R.Value) bool:			for ; ; count++ {
														if v, open := c.Recv(); !open || f(count, v) != r {
															break
														}
													}

	case func(R.Value, R.Value) bool:				for ; ; count++ {
														if v, open := c.Recv(); !open || f(R.ValueOf(count), v) == !r {
															break
														}
													}

	default:										if f := R.ValueOf(f); f.Kind() == R.Func {
														switch f.Type().NumIn() {
														case 1:				open := false
																			p := make([]R.Value, 1, 1)

																			for ; ; count++ {
																				if p[0], open = c.Recv(); !open || f.Call(p)[0].Bool() == r {
																					break
																				}
																			}

														case 2:				open := false
																			p := make([]R.Value, 2, 2)

																			for ; ; count++ {
																				p[0] = R.ValueOf(count)
																				if p[1], open = c.Recv(); !open || f.Call(p)[0].Bool() == r {
																					break
																				}
																			}
														}
													}
	}
	return
}

func While(container, f interface{}) (count int) {
	switch container := container.(type) {
	case PartiallyIterable:	count = container.While(true, f)

	case IndexedReader:		count = whileIndexedReader(container, true, f)

	default:				switch c := R.ValueOf(container); c.Kind() {
							case R.Slice:		count = whileSlice(c, true, f)

//							case R.Chan:		count = whileChannel(c, true, f)

//							case R.Func:		count = whileFunction(c, true, f)
							}
	}
	return
}

func Until(container, f interface{}) (count int) {
	switch container := container.(type) {
	case PartiallyIterable:	count = container.While(false, f)

	case IndexedReader:		count = whileIndexedReader(container, false, f)

	default:				switch c := R.ValueOf(container); c.Kind() {
							case R.Slice:		count = whileSlice(c, false, f)

//							case R.Chan:		count = whileChannel(c, false, f)

//							case R.Func:		count = whileFunction(c, false, f)
							}
	}
	return
}