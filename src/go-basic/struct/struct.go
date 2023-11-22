package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var bany Customer
	bany.Name = "Albany Siswanto"
	bany.Address = "Subang"
	bany.Age = 18
	fmt.Println(bany.Name)
	fmt.Println(bany.Address)
	fmt.Println(bany.Age)

	/*
		Struct Literals
	*/

	budi := Customer{
		Name:    "Budiman",
		Address: "Bandung",
		Age:     17,
	}
	fmt.Println(budi)

	// ATAU

	udin := Customer{"Udin", "Papua", 13}
	fmt.Println(udin)
}
