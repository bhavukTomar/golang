package main

import "fmt"

func main() {
	fmt.Println("This is a class for structs")

	bhavuk := User{"Bhavuk", 23, true, "bhavuktomar@gmail.com"}
	bhavuk.GetStatus()
	fmt.Println("Bhavuk's details are", bhavuk)
	fmt.Printf("All details of bhavuk is %+v \n", bhavuk)
	fmt.Printf("Name is :%+v, Email is :%+v \n", bhavuk.Name, bhavuk.Email)

	bhavuk.NewMail()
	fmt.Printf("Name is :%+v, Email is :%+v \n", bhavuk.Name, bhavuk.Email)
}

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

func (u User) GetStatus() {
	fmt.Println("User status is :", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email status is :", u.Email)
}
