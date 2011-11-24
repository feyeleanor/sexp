package sexp

import (
	"fmt"
	"github.com/feyeleanor/raw"
	"github.com/feyeleanor/slices"
	"reflect"
)

func Len(container interface{}) (l int) {
	switch container := container.(type) {
	case Linear:			l = container.Len()
	default:				switch v := reflect.ValueOf(container); v.Kind() {
							case reflect.Slice:		fallthrough
							case reflect.Map:		l = v.Len()
							}
	}
	return
}

func Cap(container interface{}) (l int) {
	switch container := container.(type) {
	case FixedSize:			l = container.Cap()
	default:				switch v := reflect.ValueOf(container); v.Kind() {
							case reflect.Slice:		fallthrough
							case reflect.Map:		l = v.Cap()
							}
	}
	return
}

func Each(container, f interface{}) {
	switch container := container.(type) {
	case Iterable:			container.Each(f)

	case IndexedReader:		end := container.Len()
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

	case MappedReader:		switch f := f.(type) {
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


	default:				switch c := reflect.ValueOf(container); c.Kind() {
							case reflect.Slice:		end := c.Len()
													switch f := f.(type) {
													case func(interface{}):					for i := 0; i < end; i++ {
																								f(c.Index(i).Interface())
																							}


													case func(int, interface{}):			for i := 0; i < end; i++ {
																								f(i, c.Index(i).Interface())
																							}


													case func(interface{}, interface{}):	for i := 0; i < end; i++ {
																								f(i, c.Index(i).Interface())
																							}

													default:								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
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
													
							case reflect.Map:		switch f := f.(type) {
													case func(interface{}):					for _, key := range c.MapKeys() {
																								f(c.MapIndex(key).Interface())
																							}

													case func(interface{}, interface{}):	for _, key := range c.MapKeys() {
																								f(key.Interface(), c.MapIndex(key).Interface())
																							}

													default:								if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
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
}

func Cycle(container interface{}, count int, f func(interface{})) (i int) {
	switch {
	case count == 0:	for ; ; i++ { Each(container, f) }
	default:			for ; i < count; i++ { Each(container, f) }
	}
	return
}

func Transform(container interface{}, f func(interface{}) interface{}) {
	switch container := container.(type) {
	case Transformable:		container.Transform(f)
	case Indexable:			end := container.Len()
							for i := 0; i < end; i++ {
								container.Set(i, f(container.At(i)))
							}

	default:				switch c := reflect.ValueOf(container); c.Kind() {
							case reflect.Slice:		end := c.Len()
													for i := 0; i < end; i++ {
														v := c.Index(i)
														v.Set(reflect.ValueOf(f(v.Interface())))
													}
							case reflect.Map:		for _, key := range c.MapKeys() {
														c.SetMapIndex(key, reflect.ValueOf(f(c.MapIndex(key))))
													}
							}
	}
}

func Collect(container interface{}, f func(interface{}) interface{}) (r interface{}) {
	switch container := container.(type) {
	case Collectable:		r = container.Collect(f)
	default:				switch c := reflect.ValueOf(container); c.Kind() {
							case reflect.Slice:		end := c.Len()
													s := reflect.MakeSlice(c.Type(), end, end)
													for i := 0; i < end; i++ {
														s.Index(i).Set(reflect.ValueOf(f(c.Index(i).Interface())))
													}
													r = s
							case reflect.Map:		m := reflect.MakeMap(c.Type())
													for _, key := range c.MapKeys() {
														m.SetMapIndex(key, reflect.ValueOf(f(c.MapIndex(key))))
													}
													r = m
							}
	}
	return
}

func Reduce(container, seed interface{}, f func(interface{}, interface{}) interface{}) (r interface{}) {
	r = seed
	Each(container, func(x interface{}) {
		r = f(r, x)
	})
	return
}

//	While processes values from a container whilst a condition is true or until the end of the container is reached.
//	Returns the count of items which pass the test.
func While(container interface{}, f func(interface{}) bool) (i int) {
	raw.Catch(func() {
		Each(container, func(x interface{}) {
			if f(x) {
				i++
			} else {
				raw.Throw()
			}
		})
	})
	return
}

//	Until processes values from a container until a condition is true or until the end of the container is reached.
//	Returns the count of items which fail the test.
func Until(container interface{}, f func(interface{}) bool) (i int) {
	raw.Catch(func() {
		Each(container, func(x interface{}) {
			if f(x) {
				raw.Throw()
			} else {
				i++
			}
		})
	})
	return
}

func Any(container interface{}, f func(interface{}) bool) (b bool) {
	if l := Until(container, f); l > 0 {
		 b = l < Len(container)
	}
	return 
}

func All(container interface{}, f func(interface{}) bool) (b bool) {
	if l := While(container, f); l > 0 {
		b = l == Len(container)
	}
	return
}

func None(container interface{}, f func(interface{}) bool) (b bool) {
	return Until(container, f) == Len(container)
}

func One(container interface{}, f func(interface{}) bool) (b bool) {
	raw.Catch(func() {
		Each(container, func(x interface{}) {
			if f(x) {
				if b {
					b = false
					raw.Throw()
				} else {
					b = true
				}
			}
		})
	})
	return
}

func Count(container interface{}, f func(interface{}) bool) (n int) {
	Each(container, func(x interface{}) {
		if f(x) { n++ }
	})
	return
}


func Density(container interface{}, f func(interface{}) bool) (r float64) {
	if l := Len(container); l > 0 {
		r = float64(Count(container, f)) / float64(l)
	}
	return 
}

func Dense(container interface{}, d, t float64, f func(interface{}) bool) bool {
	r := Density(container, f)
	return r - t > d
}

func Most(container interface{}, t float64, f func(interface{}) bool) bool {
	return Dense(container, 0.5, t, f)
}

func Reverse(container interface{}) {
	switch container := container.(type) {
	case Reversible:		container.Reverse()
	case Indexable:			end := container.Len() - 1
							for i := 0; i < end; i++ {
								x, y := container.At(i), container.At(end)
								container.Set(i, y)
								container.Set(end, x)
								end--
							}
	case reflect.Value:		switch container.Kind() {
							case reflect.Slice:		end := container.Len() - 1
													for i := 0; i < end; i++ {
														x, y := container.Index(i), container.Index(end)
														temp := x.Interface()
														x.Set(y)
														y.Set(reflect.ValueOf(temp))
														end--
													}
							}
	default:				Reverse(reflect.ValueOf(container))
	}
}

/*
	Calculated the depth of nesting of a container.
	A scalar value and an empty container will both return a depth of zero.
	All other containers will return a depth of 1+.
*/
func Depth(container interface{}) (d int) {
	switch container := container.(type) {
	case Nested:			if r := container.Depth() + 1; r > d {
								d = r
							}
	default:				Each(container, func(v interface{}) {
								if r := Depth(v) + 1; r > d {
									d = r
								}
							})
	}
	return
}

func Flatten(container interface{}) {
	switch container := container.(type) {
	case Flattenable:		container.Flatten()
	default:				Transform(container, func(v interface{}) interface{} {
								Flatten(v)
								return v
							})
	}
}

func Append(container, value interface{}) {
	switch container := container.(type) {
	case Appendable:		container.Append(value)
	case Resizeable:		length := container.Len() + 1
							Resize(container, length, length)
							length--
							container.(Indexable).Set(length, value)
	case reflect.Value:		switch container.Kind() {
							case reflect.Slice:		container.Set(reflect.Append(container, reflect.ValueOf(value)))
													
							}
	default:				Append(reflect.ValueOf(container), value)
	}
}

func Repeat(container interface{}, count int) {
	if count > 0 {
		switch container := container.(type) {
		case Repeatable:		container.Repeat(count)
		case Resizeable:		length := container.Len()
								Resize(container, length, length * count)
								for start := length; count > 1; count-- {
									BlockCopy(container, start, 0, length)
									start += length
								}
		}
	}
}

func Subslice(container interface{}, start, end int) (r interface{}) {
	if start < 0 {
		start = 0
	}
	switch container := container.(type) {
	case Sliceable:				r = container.Subslice(start, end)
	case Indexable:				LastIndex := container.Len() - 1
								if end > LastIndex {
									end = LastIndex
								}
								l := end - start
								if c, ok := container.(Typed); ok {
									s := reflect.MakeSlice(c.Type(), l, l)
									for i, j := 0, start; j < end; j++ {
										s.Index(i).Set(reflect.ValueOf(container.At(j)))
										i++
									}
									r = s.Interface()
								} else {
									s := make([]interface{}, l, l)
									for i, j := 0, start; j < end; j++ {
										s[i] = container.At(j)
										i++
									}
									r = s
								}
	default:					switch container := reflect.ValueOf(container); container.Kind() {
								case reflect.Slice:			LastIndex := container.Len() - 1
															if end > LastIndex {
																end = LastIndex
															}
															l := end - start
															s := reflect.MakeSlice(container.Type(), l, l)
															for i, j := 0, start; j < end; j++ {
																s.Index(i).Set(container.Index(j))
																i++
															}
															r = s.Interface()
								}
	}
	return
}

func Overwrite(destination interface{}, offset int, source interface{}) {
	switch destination := destination.(type) {
	case Overwriteable:			destination.Overwrite(offset, source)
	default:					switch d := reflect.ValueOf(destination); d.Kind() {
								case reflect.Slice:			switch s := reflect.ValueOf(source); s.Kind() {
															case reflect.Slice:			reflect.Copy(d.Slice(offset, d.Len() - 1), s)
															}
								}
	}
}

func BlockCopy(container interface{}, d, s, n int) {
	if d > -1 && s > -1 && d != s && n > 0 {
		switch container := container.(type) {
		case Blitter:			container.BlockCopy(d, s, n)
		case Indexable:			switch {
								case d > s:		n = boundOffset(container, d, n)
												s += n
												for end := d + n; d < end; {
													end--
													s--
													container.Set(end, container.At(s))
												}

								case d < s: 	n = boundOffset(container, s, n)
												d += n
												for end := s + n; end > s; {
													end--
													d--
													container.Set(d, container.At(end))
												}
								}

		default:				if c := reflect.ValueOf(container); c.Kind() == reflect.Slice {
									switch {
									case d > s:		n = boundOffset(c, d, n)
													s += n
													for end := d + n; d < end; {
														end--
														s--
														c.Index(end).Set(c.Index(s))
													}

									case d < s:		n = boundOffset(c, s, n)
													d += n
													for end := s + n; end > s; {
														end--
														d--
														c.Index(d).Set(c.Index(end))
													}
									}
								}
		}
	}
}

func BlockSwap(container interface{}, d, s, n int) {
	if d > -1 && s > -1 && d != s && n > 0 {
		temp := Subslice(container, d, d + n)
		BlockCopy(container, d, s, n)
		Overwrite(container, s, temp)
	}
}

func BlockClear(container interface{}, d, n int) {
	if d > -1 && n > 0 {
		switch container := container.(type) {
		case Blitter:			container.BlockClear(d, n)

		case Indexable:			n = boundOffset(container, d, n)
								for end := d + n; d < end; d++ {
									container.Clear(d)
								}

		default:				if c := reflect.ValueOf(container); c.Kind() == reflect.Slice {
									blank := reflect.Zero(c.Type().Elem())
									n = boundOffset(c, d, n)
									for end := d + n; d < end; d++ {
										c.Index(d).Set(blank)
									}
								}
		}
	}
}

/*
	Create a new memory container and copy contents across to it.
	Returns nil when reallocation fails.
*/
func Reallocate(container interface{}, length, capacity int) (r interface{}) {
	switch c := container.(type) {
	case Resizeable:			c.Reallocate(length, capacity)
								r = c

	default:					println("not resizeable")
								if c := reflect.ValueOf(container); c.Kind() == reflect.Slice {
println("so brute-force slice reallocation")
fmt.Printf("container = %v\n", container)
									if length > capacity {
										length = capacity
									}
fmt.Printf("desired: length = %v, capacity = %v\n", length, capacity)
									if c.Cap() != capacity {
fmt.Printf("creating new %v slice\n", c.Type())
										n := reflect.MakeSlice(c.Type(), length, capacity)
										reflect.Copy(n, c)
										c = n
									}

									if c.Len() != length {
										c = makeAddressable(c)
										c.SetLen(length)
									}
fmt.Printf("c = %v\n", c.Interface())

									r = c.Interface()
								}
	}
	return
}

/*
	Resize a container by a delta of n elements at the insertion point x.
	Returns nil when expansion fails.
*/
func Resize(container interface{}, x, n int) (r interface{}) {
	r = container
	if x > -1 {
		switch block := r.(type) {
		case Resizeable:			length := block.Len() + n
									capacity := block.Cap()
									if n > 0 {
										if length > capacity {
											capacity = length
										}
										block.Reallocate(length, capacity)
										BlockCopy(block, x + n, x, length - n)
										BlockClear(block, x, n)
									} else {
										BlockCopy(block, x + n, x, n)
										block.Reallocate(length, capacity)
									}

		default:					if c := reflect.ValueOf(r); c.Kind() == reflect.Slice {
										if x <= c.Len() {
											if length := c.Len() + n; length > c.Cap() {
												Reallocate(container, length, length)
												r = container
											}
fmt.Printf("container = %v, x = %v, n = %v\n", container, x, n)
											BlockCopy(container, x + n, x, n)
fmt.Printf("container = %v, x = %v, n = %v\n", container, x, n)
											BlockClear(container, x, n)
fmt.Printf("container = %v\n", container)
										}
									}
		}
	}
	return
}

func First(container interface{}, i int) interface{} {
	return Subslice(container, 0, i - 1)
}

func Last(container interface{}, i int) interface{} {
	l := Len(container)
	return Subslice(container, l - i, l - 1)
}

/*
func PopFirst(container interface{}) interface{} {
	return s.At(0), s.Section(1, s.Len())
}

func PopLast(container interface{}) interface{} {
	return s.At(l), s.Section(0, l)
}
*/

func Feed(container interface{}, c chan interface{}, f func(x interface{}) interface{}) {
	switch container := container.(type) {
	case Feeder:		container.Feed(c, f)

	default:			go func() {
							Each(container, func(v interface{}) {
								c <- f(v)
							})
						}()
	}
}

func Pipe(container interface{}, f func(x interface{}) interface{}) (c chan interface{}) {
	switch container := container.(type) {
	case Piper:			c = container.Pipe(f)

	case Feeder:		c = make(chan interface{})
						go func() {
							WaitFor(func() {
								container.Feed(c, f)
							})
							close(c)
						}()

	default:			c = make(chan interface{})
						go func() {
							Each(container, func(v interface{}) {
								c <- f(v)
							})
							close(c)
						}()
	}
	return
}