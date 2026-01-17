package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Dwi")
	if result != "Hello Sandy" {
		t.Fail()
	}
}
