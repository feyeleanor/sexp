include $(GOROOT)/src/Make.inc

TARG=sexp

GOFILES=\
	errors.go\
	interfaces.go\
	helpers.go\
	sexp.go\
	combine.go

include $(GOROOT)/src/Make.pkg