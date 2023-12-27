package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("We are handling web Requests now")
	PerformFormPostRequest()
}

func PerformGetRequews() {
	const myGetUrl = "http://localhost:8000/get"

	response, err := http.Get(myGetUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseBuilder strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseBuilder.Write(content)

	fmt.Println("ByteCount is : ", byteCount)
	fmt.Println(responseBuilder.String())

	fmt.Println(string(content))
}

func PerformPostRequest() {
	const myPostUrl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"name": "Bhavuk"
		}
	`)

	response, err := http.Post(myPostUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformFormPostRequest() {
	const myPostUrl = "http://localhost:8000/postform"

	//fake json payload

	data := url.Values{}
	data.Add("name", "Bhavuk Tomar")

	response, err := http.PostForm(myPostUrl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}
