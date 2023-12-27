package main

import "fmt"

func main() {
	fmt.Println("Hey there lets learn array")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[2] = "orange"
	fruitList[3] = "banana"

	fmt.Println("Your fruit list is : ", fruitList)
	fmt.Println("Your fruit list is : ", len(fruitList))

	var vegList = [5]string{"potato", "beans", "onion"}
	fmt.Println("Your vegie list is : ", vegList)
	fmt.Println("Your vegie list is : ", len(vegList))
}
