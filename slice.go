package main

import "fmt"

func main() {
	names := []string{"Wahid", "Rizka", "Fathur", "Rohman", "Aurelia", "Syahwa"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	var slice4 []string = names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"

	fmt.Println(daySlice1) // Sabtu Baru, Minggu Baru
	fmt.Println(days)      // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru

	daySlice2 := append(daySlice1, "Libur baru")
	daySlice2[0] = "Sabtu Lama"
	// daysBaru := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu", "Libur baru"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Wahid"
	newSlice[1] = "Wahid"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	var newSlice2 = append(newSlice, "Wahid")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Aurelia"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	iniSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
