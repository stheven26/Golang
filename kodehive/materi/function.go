package main

import "fmt"

type BlackList func(string) bool

// func add(x, y int) int {
// 	total := x + y
// 	return total
// }

// func registerUser(nama string, blackList func(string) bool)

func registerUser(nama string, blackList BlackList) {
	if blackList(nama) {
		fmt.Println("Anda diblokir!", nama)
	} else {
		fmt.Println("Selamat datang", nama)
	}
}

func main() {
	blokirUser := func(nama string) bool {
		return nama != "stheven"
	}
	registerUser("erlangga", blokirUser)
	registerUser("stheven", blokirUser)

	// sum := add(20, 30)
	// fmt.Println(sum)
}
