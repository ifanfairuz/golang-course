package main

import "fmt"

func main() {
	var name1 = "ifan"
	var name2 = "ifan"

	var sama = name1 == name2
	fmt.Println(sama)

	name1 = "masrur"
	name2 = "ryan"

	sama = name1 == name2
	fmt.Println(sama)
}
