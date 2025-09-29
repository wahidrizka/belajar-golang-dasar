package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Wahid"
	// person["address"] = "Cilacap"

	person := map[string]string{
		"name":    "Wahid",
		"address": "Cilacap",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Wahid"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
