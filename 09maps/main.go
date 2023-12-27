package main

import (
	"fmt"
)

func main() {
	fmt.Println("Time to learn about maps")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["PY"] = "python"
	languages["CPP"] = "C++"

	fmt.Println("Our languages are : ", languages)
	fmt.Println("JS is short for : ", languages["JS"])

	delete(languages, "JS")
	fmt.Println("Our languages are : ", languages)

	//loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("The key is %v, and the value is %v \n", key, value)
	}
}
