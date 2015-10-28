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

// a factory generator
func newPrinter() *printer {
	_printer := new(printer)
	return _printer
}

// more abstraction is better because it helps keep the logic
// decoupled.
func stringEquals(stringOne, s2 string) bool {
	return stringOne == s2
}

var ErrorConstant = "runtime error: invalid memory address or nil pointer dereference"

func Recoverer(message Message) {
	if e, isError := recover().(error); isError {
		if errorTwo, okay := e.(runtime.Error); okay {
			if stringEquals(errorTwo.Error(), ErrorConstant) {
				p := newPrinter()
				p.print(&message)
			}
		}
	}
}

var Generator = &Recoverer_generator{c: make(chan func(Message), 100)}

func init() {
	for i := 0; i < 100; i++ {
		Generator.Put(func(message Message) {
			if e, isError := recover().(error); isError {
				if errorTwo, okay := e.(runtime.Error); okay {
					if stringEquals(errorTwo.Error(), ErrorConstant) {
						p := newPrinter()
						p.print(&message)
					}
				}
			}
		})
	}
}

type Recoverer_generator struct{ c chan func(Message) }

func (r *Recoverer_generator) Get() (f func(Message)) {
	select {
	case f = <-r.c:
	default:
		f = func(message Message) {
			if e, isError := recover().(error); isError {
				if errorTwo, okay := e.(runtime.Error); okay {
					if stringEquals(errorTwo.Error(), ErrorConstant) {
						p := newPrinter()
						p.print(&message)
					}
				}
			}
		}
	}
	return
}

func (r *Recoverer_generator) Put(f func(Message)) {
	select {
	case r.c <- f:
	default:
	}
}
