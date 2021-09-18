package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Adi",
			request: "Adi",
		},
		{
			name:    "Martin",
			request: "Martin",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Adi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Adi")
		}
	})

	b.Run("Martin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Martin")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Adi")
	}
}

func BenchmarkHelloWorldMartin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Martin")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "Adi",
			request:  "Adi",
			expected: "Hello Adiputra",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")

	assert.Equal(t, "Hello Eko", result, "Result harus Hello Eko")

	fmt.Println("Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		t.Error("Result harus Hello Eko")
	}

	fmt.Println("Done")
}

func TestHelloWorldAdi(t *testing.T) {
	result := HelloWorld("Adi")

	if result != "Hello Adi" {
		t.Fatal("Result harus Hello Adi")
	}

	fmt.Println("Done")
}
