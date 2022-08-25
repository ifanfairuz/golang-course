package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32) // akan di konversin dair int32 ke int64

	/**
	 * overlaps karena int8 hanya di range -127 sampai 128
	 * maka akan otomatis menjadi nilai terbawah dari tipe data = -127 untuk int8
	**/
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Ifan"
	var e byte = name[0]         // = byte code char I yaitu 73
	var charI string = string(e) // konversi byte ke string

	fmt.Println(charI)
}
