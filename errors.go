package sexp

import "fmt"

type Errno int

func (e Errno) String() (err string) {
	if err = errText[e]; err == "" {
		err = fmt.Sprintf("errno %v", int(e))
	}
	return 
}

const(
	OK					= Errno(iota)
	EXPANSION_FAILED
)

var errText = map[Errno]string {
	EXPANSION_FAILED:		"Unable to expand container",
}