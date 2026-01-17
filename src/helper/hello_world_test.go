package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Sandy")
	if result != "Hello Sandy" {
		panic("Result is not Hello Sandy")
	}
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Sandy")
	if result != "Hello Sandy" {
		t.Error("should Hello Sandy")
	}

	fmt.Println("Dieksekusi")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Sandy")
	if result != "Hello Sandy" {
		t.Fatal("should Hello Sandy")
	}

	fmt.Println("Tidak Dieksekusi")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Sandy")
	assert.Equal(t, "Hello Sandy", result)
	fmt.Println("Dieksekusi")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Sandy")
	require.Equal(t, "Hello Sandy", result)
	fmt.Println("Dieksekusi")
}
