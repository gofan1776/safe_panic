package safe_panic

type T struct {
	A *struct {
		B *struct {
			C *struct {
			}
		}
	}
}

// main at the bottom like every God-fearing program should be.
func ExampleMain() {
	// just call this one super easy line!
	// spotting the null dereferenced pointer is left as an
	// exercise for the reader.
	defer Recoverer(Message("oopsie daisy!"))

	baz := func(fn1 func(), fn ...func(func())) {
		for i := range fn {
			fn[i](fn1)
		}
	}

	bar := func(fn func()) {
		fn()
	}

	quux := func(fn func()) func() {
		return fn
	}
	_ = quux

	foo := func() {
		t := new(T)
		var f struct{}

		// OOPS!
		t.A.B.C = &f
	}

	// works!
	foo()

	// even works nested!!1one
	bar(foo)

	// and stuff
	baz(foo, bar, bar, bar, bar)
}
