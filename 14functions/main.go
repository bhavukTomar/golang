package main

import "fmt"

func main() {
	fmt.Println("We are going to learn about functions")
	greeter()

	result := adder(3, 4)

	result2, message := adder2(3, 4)

	proResult := proAdder(2, 3, 4, 5)

	fmt.Println("Pro result : ", proResult)

	fmt.Printf("Result : %v, Value : %v \n", result2, message)

	fmt.Println("Result : ", result)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func adder2(valOne int, valTwo int) (int, string) {
	return valOne + valTwo, "This is my message"
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func greeter() {
	fmt.Println("Namaste from golang")
}
