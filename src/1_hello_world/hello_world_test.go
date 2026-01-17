package main

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Sandy")
	if result != "Hello Sandy" {
		panic("Result is not Hello Sandy")
	}
}
