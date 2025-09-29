package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name

	return hello
}

func main() {
	result := getHello("Wahid")
	fmt.Println(result)

	fmt.Println(getHello("Rizka"))
	fmt.Println(getHello("Fathurrohman"))
}
