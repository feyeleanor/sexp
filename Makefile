include $(GOROOT)/src/Make.inc

TARG=sexp

GOFILES=\
	errors.go\
	interfaces.go\
	helpers.go\
	sexp.go\
	combine.go\
	slice.go\
	slice_value.go\
	node.go\
	cons_cell.go\
	list_header.go\
	linear_list.go\
	cyclic_list.go

include $(GOROOT)/src/Make.pkg