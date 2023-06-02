package main

import "fmt"

func main() {

	var number int

	fmt.Println("Please enter your number: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("This is an even number.")
	} else {
		fmt.Println("This is an odd number.")
	}
}
