package main

import "fmt"

type Laptop struct {
	Username
	City  string
	State string
	Email string
}

type Username struct {
	Name  string
	Phone int
	Age   int
	City  string
	State string
	Email string
}

func main() {
	var details Username

	details.Name = "Brock"
	details.Phone = 9090456789
	details.Age = 99
	details.City = "KondaGaon"
	details.State = "SriLanka"
	details.Email = "brock@gmail.com"

	printUser(details)
}

func printUser(userDetails Username) {

	fmt.Println(userDetails.Name)
	fmt.Println(userDetails.Phone)
	fmt.Println(userDetails.Age)
	fmt.Println(userDetails.City)
	fmt.Println(userDetails.State)
	fmt.Println(userDetails.Email)
}