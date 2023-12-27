package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Lets learn about file system")

	content := "Hello there you are in new file"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Size of file is : ", length)

	defer file.Close()

	readFile(file.Name())

}

func readFile(filePath string) {
	databyte, err := os.ReadFile(filePath)
	checkNilErr(err)

	fmt.Println(string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
