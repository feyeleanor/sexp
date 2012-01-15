package sexp

 import "reflect"

type PartiallyIterable interface {
	While(bool, interface{}) (bool, int)
}


func whileIndexedReader(container IndexedReader, r bool, f interface{}) (ok bool, count int) {
	end := container.Len()
	switch f := f.(type) {
	case func(interface{}) bool:				for i := 0; i < end; i++ {
													if f(container.At(i)) != r {
														break
													}
													count++
												}
												ok = true

	case func(int, interface{}) bool:			for i := 0; i < end; i++ {
													if f(i, container.At(i)) != r {
														break
													}
													count++
												}
												ok = true

	case func(interface{}, interface{}) bool:	for i := 0; i < end; i++ {
													if f(i, container.At(i)) != r {
														break
													}
													count++
												}
												ok = true

	default:									if f := reflect.ValueOf(f); f.Kind() == reflect.Func {
													switch f.Type().NumIn() {
													case 1:				for ; count < end && f.Call(valueslice(container.At(count)))[0].Bool() != r; count++ {}
																		ok = true

													case 2:				for ; count < end && f.Call(valueslice(count, container.At(count)))[0].Bool() != r; count++ {}
																		ok = true
													}
												}
	}
	return
}

func While(container, f interface{}) (ok bool, count int) {
	switch container := container.(type) {
	case PartiallyIterable:	ok, count = container.While(true, f)

	case IndexedReader:		ok, count = whileIndexedReader(container, true, f)

//	case MappedReader:		ok, count = whileMappedReader(container, f)
	}
	return
}

func Until(container, f interface{}) (ok bool, count int) {
	switch container := container.(type) {
	case PartiallyIterable:	ok, count = container.While(false, f)

	case IndexedReader:		ok, count = whileIndexedReader(container, false, f)
	}
	return
}