package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?courseName=java&paymentId=9088466"

func main() {
	fmt.Println("Lets learn how to handle url")

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Path)
	fmt.Println(result.Host)
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()

	fmt.Printf("Type of query params are %T \n", qParams)

	fmt.Println(qParams["courseName"])
	fmt.Println(qParams["paymentId"])

	for _, val := range qParams {
		fmt.Println("The value is :", val)
	}

	partsUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev:3000",
		Path:    "learn",
		RawPath: "courseName=java&paymentId=9088466",
	}

	anotherUrl := partsUrl.String()
	fmt.Println(anotherUrl)
}
