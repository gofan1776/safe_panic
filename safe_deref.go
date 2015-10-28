package safe_panic

import (
	"fmt"
	"runtime"
)

type Message string

type printer struct{}

// use pointer to save memory from copying a whole string.
func (p *printer) print(msg *Message) {
	fmt.Println(msg)
}

// a factory like my CS 202 professor taught
func newPrinter() *printer {
	_printer := new(printer)
	return _printer
}

// more abstraction is better because it helps keep the logic
// decoupled.
func stringEquals(stringOne, s2 string) bool {
	return stringOne == s2
}

func Recoverer(message Message) {
	if e, isError := recover().(error); isError {
		if errorTwo, okay := e.(runtime.Error); okay {
			if stringEquals(errorTwo.Error(), errorConstant) {
				p := newPrinter()
				p.print(&msg)
			}
		}
	}
}
