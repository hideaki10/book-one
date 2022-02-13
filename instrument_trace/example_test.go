package trace_test

func a() {
	defer trace.Trace()()
	b()
}

func b() {
	defer trace.Trace()()
	c()
}

func c() {
	defer trace.Trace()()
	d()
}

func d() {
	defer trace.Trace()()
}

func ExampleTrace() {
	a()
	// Output:
	// enter: a
	// enter: b
	// enter: c
	// enter: d
	// exit: d
	// exit: c
	// exit: b
	// exit: a
}
