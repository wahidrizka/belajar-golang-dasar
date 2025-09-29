package main

func getFullName() (string, string) {
	return "Wahid", "Rizka"
}

func main() {
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullName()
	println(firstName)
}
