package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	wahid := Man{"Wahid"}
	wahid.Married()

	fmt.Println(wahid.Name)
}
