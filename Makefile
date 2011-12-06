include $(GOROOT)/src/Make.inc

TARG=sexp

GOFILES=\
	interfaces.go\
	helpers.go\
	iterate.go\
	transform.go\
	collect.go\
	sexp.go\
	combine.go

include $(GOROOT)/src/Make.pkg