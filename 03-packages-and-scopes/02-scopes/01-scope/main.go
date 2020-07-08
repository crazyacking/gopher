package main

// file scope
import "fmt"

// package scope
const ok = true

func main() {
	var hello = "Hello"

	fmt.Println(hello, ok)
}
