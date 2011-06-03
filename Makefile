include $(GOROOT)/src/Make.inc

TARG=sexp

GOFILES=\
	interfaces.go\
	slice.go\
	node.go\
	cons_cell.go\
	list_header.go\
	linear_list.go\
	cyclic_list.go

include $(GOROOT)/src/Make.pkg