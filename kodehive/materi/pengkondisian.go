package main

import "fmt"

func main() {
	if age := 10; age > 60 {
		fmt.Print("Getting Older")
	} else if age > 30 {
		fmt.Print("Getting Wiser")
	} else if age > 20 {
		fmt.Print("adulthood")
	} else if age > 10 {
		fmt.Print("young blood")
	} else {
		fmt.Print("booting up")
	}

	if poin := 2; poin > 5 {
		switch {
		case poin == 10:
			fmt.Print("Nilai anda sempurna")
		case poin > 5 && poin < 10:
			fmt.Print("Nice")
		}
	}
	// else {
	// 	if poin <= 5 {
	// 		fmt.Print("Not bad lah!")
	// 	} else if == 3 {
	// 		fmt.Print("Hope u pass the exam")
	// 	} else {
	// 		fmt.Print("Tetap semangat")
	// 		if poin <= 2 && poin > 0{
	// 			fmt.Print("Tetap mencoba, rajin belajar!")
	// 		} else {
	// 			fmt.Print("Nilai anda sangat bad")
	// 		}
	// 	}
	// }

	// switch {
	// case poin == 10:
	// 	fmt.Print("Nilai ujian anda sempurna")
	// case poin >= 5 && poin <= 9:
	// 	fmt.Print("Anda lulus ujian dengan nilai", poin)
	// case poin < 5:
	// 	fmt.Print("Anda tidak lulus ujian")
	// }

	// switch poin {
	// case 10:
	// 	fmt.Print("Nilai ujian anda sempurna")
	// 	fallthrough
	// case 5, 6, 7, 8, 9:
	// 	fmt.Print("Anda lulus ujian dengan nilai ", poin)
	// 	fallthrough
	// case 1, 2, 3, 4:
	// 	fmt.Print("Anda tidak lulus ujian")
	// }

	var sum int
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Print("Total penjumlahan ", sum)

	var namaMahasiswa = [...]string{"stheven", "erlangga", "sanjaya", "dwi", "ilham"}

	for i := 0; i < len(namaMahasiswa); i++ {
		fmt.Println("Elemen ", i, ":", namaMahasiswa[i])
	}

	for _, a := range namaMahasiswa {
		fmt.Println(a)
	}
}
