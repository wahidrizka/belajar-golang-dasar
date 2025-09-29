package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var wahid Customer
	fmt.Println(wahid)

	wahid.Name = "Wahid Fathurrohman"
	wahid.Address = "Indonesia"
	wahid.Age = 24

	fmt.Println(wahid)
	fmt.Println(wahid.Name)
	fmt.Println(wahid.Address)
	fmt.Println(wahid.Age)

	aurel := Customer{
		Name:    "Aurelia",
		Address: "Indonesia",
		Age:     24,
	}

	fmt.Println(aurel)

	syahwa := Customer{"Aurel", "Indonesia", 24}
	fmt.Println(syahwa)

	aurel.sayHello("Wahid")
	aurel.sayHello("Fathur")
	wahid.sayHello("Aurel")
}
