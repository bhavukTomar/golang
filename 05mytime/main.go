package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to study off time in golang")

	currentTime := time.Now()

	fmt.Println("Current Time is : ", currentTime)
	fmt.Println("Current Time is : ", currentTime.Format("01-02-2006 15:04:05 Monday"))

	currentDate := time.Date(2021, time.April, 6, 15, 23, 33, 0, time.Local)
	fmt.Println("current date is : ", currentDate)
	fmt.Println("current date is : ", currentDate.Format("01-02-2006 Monday"))
}
