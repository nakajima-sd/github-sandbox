package main

import (
	"fmt"

	"github.com/nakajima-sd/github-sandbox/hello"
)

func main() {
	ok := hello.Hello()
	fmt.Println("returned value:", ok)
}

