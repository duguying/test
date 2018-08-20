package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	e := os.Stderr
	io.WriteString(e, "test stderr")
	fmt.Println("hello world")
}
