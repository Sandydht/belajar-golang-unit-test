package helper

import (
	"fmt"
	"runtime"
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

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak bisa jalan di Mac")
	}

	result := HelloWorld("Sandy")
	require.Equal(t, "Hello Sandy", result)
}

func TestSubTest(t *testing.T) {
	t.Run("Sandy", func(t *testing.T) {
		result := HelloWorld("Sandy")
		require.Equal(t, "Hello Sandy", result)
	})

	t.Run("Dwi", func(t *testing.T) {
		result := HelloWorld("Dwi")
		require.Equal(t, "Hello Dwi", result)
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Sandy)",
			request:  "Sandy",
			expected: "Hello Sandy",
		},
		{
			name:     "HelloWorld(Dwi)",
			request:  "Dwi",
			expected: "Hello Dwi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
