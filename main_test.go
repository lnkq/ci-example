package main

import "testing"

func Test_sayHello(t *testing.T) {
	name := "Alice"
	expected := "Hello, Alice! <3"

	if got := sayHello(name); got != expected {
		t.Errorf("sayHello(%q) = %q; want %q", name, got, expected)
	}
}
