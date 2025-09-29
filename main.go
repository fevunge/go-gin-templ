package main

import (
	"fmt"

	"github.com/fevunge/let-go/loop"
	"github.com/fevunge/let-go/variables"
)

const username = "fevunge"

func main() {
	number := 42
	variables.Variables()
	fmt.Println(number, username)
	loop.Loop()
}
