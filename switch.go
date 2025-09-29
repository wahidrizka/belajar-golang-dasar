package main

import (
	"fmt"
)

func main() {
	name := "Wahid"

	switch name {
	case "Wahid":
		fmt.Println("Hello Wahid")
	case "Rizka":
		fmt.Println("Hello Rizka")
	default:
		fmt.Println("Hello, who are you?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
