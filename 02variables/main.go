package main

import "fmt"

var jwtToken string = "ksdhkjhu"

const Token string = "gibbrish" // Capital first letter means public variable

func main() {
	var userName string = "bhavuktomar"
	fmt.Println(userName)
	fmt.Printf("Type of variable is %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of variable is %T \n", isLoggedIn)

	var smallInt int = 256
	fmt.Println(smallInt)
	fmt.Printf("Type of variable is %T \n", smallInt)

	var floatingValue float64 = 8.837498274987389889
	fmt.Println(floatingValue)
	fmt.Printf("Type of variable is %T \n", floatingValue)

	//default values and aliases
	var anotherInt int
	fmt.Println(anotherInt)
	fmt.Printf("Type of variable is %T \n", anotherInt)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Type of variable is %T \n", anotherString)

	var anotherFloat float64
	fmt.Println(anotherFloat)
	fmt.Printf("Type of variable is %T \n", anotherFloat)

	//implicit way of delcaring a variable
	var implicitVariable = "Hello"
	fmt.Println(implicitVariable)
	fmt.Printf("Type of variable is %T \n", implicitVariable)

	//declare variable without var keyword
	noVarKeywordVariable := 3
	fmt.Println(noVarKeywordVariable)
	fmt.Printf("Type of variable is %T \n", noVarKeywordVariable)

	fmt.Println(Token)
	fmt.Printf("Type of variable is %T \n", Token)

}
