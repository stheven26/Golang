package main

import "fmt"

func add(x, y int) int {
	total := x + y
	return total
}

func main() {
	sum := add(20, 30)
	fmt.Println(sum)
	// colors := [...]string{"red", "blue", "green"}

	// tones := [...]string{"dark", "light"}

	// if colors == tones {
	// 	fmt.Print(colors, tones)
	// } else {
	// 	fmt.Print("tidak sama")
	// }
	// kota := [5]string{"solo", "jakarta", "medan"}
	// fmt.Print(cap(kota))
	// var names [3]string
	// names[len(names)-1] = "!"
	// names[1] = "think" + names[2]
	// names[0] = "Don't"
	// names[0] += " "
	// fmt.Println(names[0], names[1], names[2])

	// gadgets := [3]string{"Confused Drone"}
	// gears := [...]string{"Confused Drone"}
	// fmt.Print(reflect.ValueOf(gadgets))
	// fmt.Print(reflect.ValueOf(gears))

	// gadgets := [3]string{"Confused Drone", "Broken Phone"}
	// gears := gadgets
	// gears[2] = "Shinny Mouse"
	// fmt.Printf("%q\n", gadgets)
	// fmt.Printf("%q\n", gadgets)
	// rows := 5

	// for i := 1; i <= rows; i++ {
	// 	k := 0
	// 	for space := 1; space <= rows-i; space++ {
	// 		fmt.Print("  ")
	// 	}
	// 	for {
	// 		fmt.Print(" *")
	// 		k++
	// 		if k == 2*i-1 {
	// 			break
	// 		}
	// 	}
	// 	fmt.Println("")
	// }
	// var sum int
	// for i := 0; i < 11; i++ {
	// 	sum += i
	// }
	// fmt.Print(sum)
	// main:
	// 	for {
	// 		switch "A" {
	// 		case "A":
	// 			break
	// 		case "B":
	// 			continue main
	// 		}
	// 	}
	// for i := 2; i <= 9; i++ {
	// 	if i%3 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	fmt.Println("Masukkan nilai dalam satuan feet: ")

	var feet float64
	fmt.Scanln(&feet)

	var meter float64 = feet * 0.3048

	fmt.Printf("Nilai %.1f feet sama dengan %.4f meter", feet, meter)
	// fmt.Print((true && false) || (false && true) || !(false && false))
	// main.namaMahasiswa()
	// namaDepan := "Stheven"

	// var namaBelakang string
	// namaBelakang = "Erlangga"

	// var _ string = "coba aja ya gaes ya" //variabel yang ga ingin diambil nilainya

	// //var nama, alamat string = "stheven", "jakarta"
	// nama, umur, nilaiUlangan := "stheven", 21, 96.8

	// f.Printf("nama saya adalah %s %s\n", namaDepan, namaBelakang)
	// f.Printf("Nama %s Umur saya %d dan nilai ulangan saya adalah %.2f", nama, umur, nilaiUlangan)

	// var number = [...]int{1, 2, 3, 4}

	// if number%2 == 0 {
	// 	fmt.Printf(number)
	// } else {
	// 	fmt.Printf("bilangan ganjil")
	// }
}
