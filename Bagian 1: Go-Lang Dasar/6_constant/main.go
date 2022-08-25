package main

import "fmt"

func main() {
	/**
	 * konstanta harus di deklarasikan langsung dengan nilai
	 * tidak error meskipun tidak digunakan
	**/

	const firstName string = "Ifan" // deklarasi constanta dengan tipe data dan nilai
	const lastName = "Fairuz"       // deklarasi constanta dengan nilai
	const age = 22                  // deklarasi constanta dengan nilai

	// age = 21 - * Error (nilai tidak bisa diubah untuk constanta)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
