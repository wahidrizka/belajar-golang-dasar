package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Wahid")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)             // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Wahid")) // tidak bisa diakses
}
