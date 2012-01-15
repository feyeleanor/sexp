package sexp

import "reflect"


func valueslice(n... interface{}) (s []reflect.Value) {
	s = make([]reflect.Value, len(n), len(n))
	for i, v := range n {
		s[i] = reflect.ValueOf(v)
	}
	return
}


type indexableSlice	[]interface{}
func (i indexableSlice) Len() int { return len(i) }
func (i indexableSlice) At(x int) interface{} { return i[x] }
func (i indexableSlice) Set(x int, v interface{}) { i[x] = v }
func (i indexableSlice) Clear(x int) { i[x] = nil }


func makeAddressable(value reflect.Value) reflect.Value {
	if !value.CanAddr() {
		ptr := reflect.New(value.Type()).Elem()
		ptr.Set(value)
		value = ptr
	}
	return value
}

func assign(location, value reflect.Value) {
	location = makeAddressable(location)
	location.Set(value)
}

func boundOffset(container Linear, base, offset int) int {
	last_index := container.Len() - 1
	if base + offset >= last_index {
		offset = last_index - base + 1
	}
	if offset < 0 {
		offset = 0
	}
	return offset
}

func Equal(x, y interface{}) bool {
	if x, ok := x.(Equatable); ok {
		return x.Equal(y)
	}
	if y, ok := y.(Equatable); ok {
		return y.Equal(x)
	}
	return reflect.DeepEqual(x, y)
}

func WaitFor(f func()) {
	done := make(chan bool)
	go func() {
		f()
		done <- true
	}()
	<-done
}