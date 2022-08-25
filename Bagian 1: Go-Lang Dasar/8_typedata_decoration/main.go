package main

import "fmt"

func main() {
	// Alias untuk tipe data
	type NoKTP string
	type Age int

	var noKTP NoKTP = "123456789"
	var age Age = 22

	fmt.Println(noKTP)
	fmt.Println(age)
}
