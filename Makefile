include $(GOROOT)/src/Make.inc

TARG=sexp

GOFILES=\
	interfaces.go\
	helpers.go\
	iterate.go\
	partial_iterate.go\
	transform.go\
	collect.go\
	reduce.go\
	combine.go\
	sexp.go

include $(GOROOT)/src/Make.pkg