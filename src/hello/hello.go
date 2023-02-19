package hello

import "fmt"

func sayHello(name string) string {
	if len(name) == 0 {
		return "Hello"
	}

	return fmt.Sprintf("Hello %s", name)
}

func sayBye(name string) string {
	if len(name) == 0 {
		return "Bye"
	}

	return fmt.Sprintf("Bye %s", name)
}
