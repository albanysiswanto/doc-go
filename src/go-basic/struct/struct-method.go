package main

import "fmt"

type Customers struct {
	Name, Address string
	Age           int
}

func (customer Customers) sayHello(name string) {
	fmt.Println("Hello,", name, "My Name is", customer.Name)
}

func main() {
	rully := Customers{Name: "Rully"}
	rully.sayHello("Agus")
}
