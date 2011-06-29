include $(GOROOT)/src/Make.inc

TARG=sexp

GOFILES=\
	errors.go\
	interfaces.go\
	helpers.go\
	sexp.go\
	combine.go\
	slice.go\
	slice_value.go

include $(GOROOT)/src/Make.pkg