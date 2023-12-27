package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("Intial value of pointer is : ", ptr)

	myVariable := 23

	ptr := &myVariable

	fmt.Println("Intial value of pointer is : ", ptr)
	fmt.Println("Intial value of pointer is : ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Intial value of pointer is : ", myVariable)

}
