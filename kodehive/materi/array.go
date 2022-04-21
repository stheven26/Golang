package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "stheven"
	names[1] = "erlangga"
	names[2] = "sanjaya"
	fmt.Println(names[0], names[1], names[2])

	var namaMahasiswa = [...]string{"stheven", "erlangga", "sanjaya", "dwi", "ilham"}
	var sliceMahasiswa = namaMahasiswa[:]
	sliceMahasiswa[2] = "haha"
	fmt.Print(sliceMahasiswa)

	num := []int{2, 3, 4}
	nums := append(num, 7, 8, 9)
	fmt.Print(nums)
}
