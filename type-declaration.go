package main

import "fmt"

func main() {
	/*
		Type Declaration
		- Untuk membuat tipe data yang sudah ada menjadi tipe data baru
	*/
	type numberPhone string

	var phone numberPhone = "0823232323"
	var contohPhone = "0823232321"
	fmt.Println(phone)
	fmt.Println(numberPhone(contohPhone))
}
