package main

import "fmt"

func main() {
	var years int
	fmt.Scanln(&years)

	//your code goes here
	result := 7
	for i := 0; i < years; i++ {
		result *= 2
	}

	fmt.Println(result)
}
