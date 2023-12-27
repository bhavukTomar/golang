package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to our pizza app :")
	fmt.Println("Please rate our app between 1 to 5 :")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating us : ", input)

	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Our new rating after adding 1 is : ", newRating+1)
	}
}
