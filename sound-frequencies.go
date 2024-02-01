package main

import "fmt"

func main() {
	var f int
	fmt.Scanln(&f)
	//your code goes here
	switch {
	case f > 20000:
		fmt.Println("Ultrasound")
	case f >= 20:
		fmt.Println("Audible")
	case f >= 0:
		fmt.Println("Infrasound")
	default:
		fmt.Println("Wrong Input")

	}
}
