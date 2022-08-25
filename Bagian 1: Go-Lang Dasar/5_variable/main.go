package main

import "fmt"

func main() {
	var name string // deklarasi variable dengan tipe data tanpa diisi

	name = "Ifan Fairuz" // assign nilai pada variable
	fmt.Println(name)    // print variable / ambil nilai variable

	name = "Ifan"     // reassign nilai pada variable
	fmt.Println(name) // print variable / ambil nilai variable

	var age = 22 // deklarasi dengan nilai
	fmt.Println(age)

	country := "Indonesia" // deklarasi tanpa sintax var tapi menggunakan :=
	fmt.Println(country)

	// deklarasi multiple variable
	var (
		firstName = "Ifan"
		lastName  = "Fairuz"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

	// error jika dideklarasikan tapi tidak digunakan / unused variable
	// language := "Bahasa Indonesia"
}
