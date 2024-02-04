package main

import (
	"fmt"
	"strings"
)

func repeat(w string, x int) {
	fmt.Println(strings.Repeat(w+"\n", x))
}

func main() {
	var w string
	fmt.Scanln(&w)
	var x int
	fmt.Scanln(&x)

	repeat(w, x)
}
