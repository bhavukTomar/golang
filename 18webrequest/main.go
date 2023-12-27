package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Lets learn web requests")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Type of response is %T \n", response)
	fmt.Println(response)

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println(string(databytes))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
