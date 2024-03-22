package hello

import "fmt"

func SayHello(name string) {
	if name == "" {
		name = "NO NAME PROVIDED"
		return
	}

	fmt.Printf("Hello, %s!\n", name)
}