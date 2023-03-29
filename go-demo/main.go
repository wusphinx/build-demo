package main

import (
	"fmt"

	"example_module/greetings"
)

func main() {
	fmt.Printf("%s, world!\n", greetings.Greeting())
}
