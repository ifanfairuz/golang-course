package main

import "fmt"

func main() {
	// string di apit oleh petik dua "Nilai String"
	fmt.Println("Ifan")
	fmt.Println("Ifan Fairuz")

	fmt.Println(len("Ifan"))      // len fungsi menghitung jumlah karakter string
	fmt.Println("Ifan"[0])        // mengambil byte code berbentuk UInt8 di mulai dari index 0 = char I dengan byte code 73
	fmt.Println("Ifan Fairuz"[1]) // mengambil byte code berbentuk UInt8 di mulai dari index 0 = char f dengan byte code 102
}
