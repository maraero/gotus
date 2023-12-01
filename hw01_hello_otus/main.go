package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	msg := "Hello, OTUS!"
	fmt.Println(reverse.String(msg))
}
