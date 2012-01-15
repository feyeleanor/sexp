package sexp

import(
//	"github.com/feyeleanor/raw"
//	"github.com/feyeleanor/slices"
//	"reflect"
)

type PartiallyIterable interface {
	While(bool, interface{}) (int, interface{})
}


func While(container, f interface{}) (count int, v interface{}) {
	switch container := container.(type) {
	case PartiallyIterable:	count, v = container.While(true, f)
	}
	return
}

func Until(container, f interface{}) (count int, v interface{}) {
	switch container := container.(type) {
	case PartiallyIterable:	count, v = container.While(false, f)
	}
	return
}