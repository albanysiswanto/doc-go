package main

import "fmt"

func EndApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Panic:", message)
}

func RunApp(error bool) {
	defer EndApp()
	if error {
		panic("Ups Error")
	}
}

func main() {
	RunApp(true)
}

/*
Penjelsan:
Cara yang benar menggunakan recover function
dengan cara recover function nya digunakan didalam defer function nya.
*/
