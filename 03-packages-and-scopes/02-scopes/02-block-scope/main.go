package main

func nope() {
	// hello and ok are only visible here
	const ok = true
	var hello = "Hello"

	_ = hello
}

func main() {
	// hello and ok are not visible here

	// ERROR:
	// fmt.Println(hello, ok)
}
