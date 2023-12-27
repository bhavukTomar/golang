package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Moday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index : %v and Value : %v \n", index, day)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 9 {
			goto lco
		}

		if rougeValue == 5 {
			// break
			rougeValue++
			continue
		}
		fmt.Println(rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Redirecting to bhavuk.tomar")
}
