include $(GOROOT)/src/Make.inc

TARG=sexp

GOFILES=\
	interfaces.go\
	node.go\
	slice.go\
	linear_list.go\
	cyclic_list.go

include $(GOROOT)/src/Make.pkg