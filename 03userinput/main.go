package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	welcome := "Please rate our pizza"
	fmt.Println(welcome)

	//comma ok, comma err (syntax for try catch)
	input, _ := reader.ReadString('\n')
	fmt.Println("Your rating is : ", input)
	fmt.Printf("Type of rating/input is : %T", input)
}
