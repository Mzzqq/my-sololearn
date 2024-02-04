package main

import (
	"fmt"
	"strings"
)

func talk() {
	fmt.Println(strings.ToUpper("GO"))
}

func main() {
	talk()
	talk()
	talk()
}
