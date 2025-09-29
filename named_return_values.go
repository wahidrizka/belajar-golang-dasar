package main

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Wahid"
	middleName = "Rizka"
	lastName = "Fathurrohman"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	println(a, b, c)
}
