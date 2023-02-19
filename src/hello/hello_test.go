package hello

import (
	"fmt"
	"testing"
)

func Test_sayHello(t *testing.T) {
	var inputs = []struct {
		name     string
		expected string
	}{
		{name: "Hoge", expected: "Hello Hoge"},
		{name: "Piyo", expected: "Hello Piyo"},
		{name: "", expected: "Hello"},
	}

	for _, item := range inputs {
		var result = sayHello(item.name)

		if result != item.expected {
			t.Errorf("\"sayHello('%s')\" FAILED, expected -> %v, result -> %v", item.name, item.expected, result)
		} else {
			t.Logf("\"sayHello('%s')\" SUCCEDED, expected -> %v, result -> %v", item.name, item.expected, result)
		}

	}
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sayHello("Hoge")
	}
}

func Example_sayHello() {
	fmt.Println(sayHello("Hoge"))
	// Output: Hello Hoge
}

func Test_sayBye(t *testing.T) {
	var inputs = []struct {
		name     string
		expected string
	}{
		{name: "Hoge", expected: "Bye Hoge"},
		{name: "Piyo", expected: "Bye Piyo"},
		{name: "", expected: "Bye"},
	}

	for _, item := range inputs {
		var result = sayBye(item.name)

		if result != item.expected {
			t.Errorf("\"sayBye('%s')\" FAILED, expected -> %v, result -> %v", item.name, item.expected, result)
		} else {
			t.Logf("\"sayBye('%s')\" SUCCEDED, expected -> %v, result -> %v", item.name, item.expected, result)
		}

	}
}
