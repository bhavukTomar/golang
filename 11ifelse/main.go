package main

import "fmt"

func main() {
	fmt.Println("In this class we are going to study control statements")

	loginDetails := 23
	var result string

	if loginDetails < 10 {
		result = "in is case"
	} else if loginDetails > 10 {
		result = "in else if case"
	} else {
		result = "in else case"
	}

	fmt.Println("*", result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num > 3 {
		fmt.Println("Number is greater than 3")
	} else {
		fmt.Println("Number is less than 3")
	}
}
