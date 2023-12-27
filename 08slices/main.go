package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Start learning about slices")

	var fruitList = []string{"Apple", "Mango", "Banana"}

	fmt.Printf("Type of fruitlist is : %T\n", fruitList)
	fmt.Println("1", fruitList)

	fruitList = append(fruitList, "Tomamto", "Peach")
	fmt.Println("2", fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println("3", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("4", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println("5", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 330
	highScores[1] = 340
	highScores[2] = 350
	highScores[3] = 360

	fmt.Println("6", highScores)

	highScores = append(highScores, 789, 440, 234)

	fmt.Println("7", highScores)

	sort.Ints(highScores)
	fmt.Println("8", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//How to remove a value from slice from based on index

	var courses = []string{"reactjs", "swift", "java", "golang", "angular"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("9", courses)
}
