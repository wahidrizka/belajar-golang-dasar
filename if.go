package main

import "fmt"

func main() {
	name := "Wahid"

	if name == "Wahid" {
		fmt.Println("Hello, Wahid!")
	} else if name == "Aurelia" {
		fmt.Println("Hello, Aurelia!")
	} else if name == "Syahwa" {
		fmt.Println("Hello, Syahwa!")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
