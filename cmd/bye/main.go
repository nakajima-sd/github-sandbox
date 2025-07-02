package main

import (
	"fmt"

	"github.com/nakajima-sd/github-sandbox/bye"
)

func main() {
	ok := bye.Bye()
	fmt.Println("returned value:", ok)
}

