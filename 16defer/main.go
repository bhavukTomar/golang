package main

import "fmt"

func main() {
	defer fmt.Println("Lets learn")
	defer fmt.Println("one")
	defer fmt.Println("two")
	mydefer()
	fmt.Println("about defer")
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
