package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("TO CREATE JSON DATA IN GO")
	// encodeJson()
	decodeJson()
}

func encodeJson() {
	lcoCourse := []course{
		{"React", 299, "bhvauk.react.in", "abc123", []string{"web-dev", "js"}},
		{"Mern", 399, "bhvauk.mern.in", "abcd123", nil},
		{"Go", 499, "bhvauk.go.in", "abce123", []string{"web-dev", "js", "mongo"}},
	}

	//package this as json data

	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonData := []byte(`
		{
			"courseName": "Go",
			"Price": 499,
			"website": "bhvauk.go.in",
			"tags": ["web-dev", "js", "mongo"]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("Valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Not Valid")
	}

	//Some case where you just want key value
	var myData map[string]interface{}

	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key is %v and Value is %v and Type is %T\n", k, v, v)
	}
}
