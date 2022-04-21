package main

import "fmt"

func main() {
	namaMahasiswa()
}

func namaMahasiswa() {
	var names [3]string
	names[0] = "stheven"
	names[1] = "erlangga"
	names[2] = "sanjaya"
	fmt.Println(names[0], names[1], names[2])
}
