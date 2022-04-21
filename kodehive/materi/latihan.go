package main

import (
	"fmt"
	"strings"
)

// type item struct {
// 	title string
// }

// type movie struct {
// 	item
// }

// type performance struct {
// 	item
// }

// type Mahasiswa struct {
// 	nama, jurusan string
// 	sks           int
// }

func main() {
	// var rows = 5
	// for i := 1; i <= rows; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print("* ")
	// 	}
	// 	fmt.Println()
	// }
	// for i := 1; i <= 6; i++ {
	// 	fmt.Printf("%s\n", strings.Repeat("#", i))
	// }
	for i := 1; i < 6; i++ {
		fmt.Printf("%6s\n", strings.Repeat("*", i))
	}

	// digits := [...][5]string{
	// 	{
	// 		"##",
	// 		"#",
	// 		"#",
	// 		"#",
	// 		"###",
	// 	},
	// 	[5]string{
	// 		"##",
	// 		"#",
	// 		"#",
	// 		"#",
	// 		"###",
	// 	},
	// }
	// fmt.Println(reflect.TypeOf(digits))

	// rates := [...]float64{
	// 	5: 1.5,
	// 	2.5,
	// 	0: 0.5,
	// }
	// fmt.Printf("%#v\n", rates)

	type (
		threeA [3]int
		threeB [3]int
	)
	num := threeA{1, 2, 3}
	num1 := threeA(threeB{1, 2, 3})
	fmt.Println(num == num1)

	coba := []string{"I'm", "going", "to", "stay", "\"here\""}
	fmt.Println(len(coba))

	var books string
	books = "Python"
	upper := strings.ToUpper(books)
	lower := strings.ToLower(books)
	fmt.Println(books)
	fmt.Println(upper)
	fmt.Println(lower)
	// mahasiswa := struct {
	// 	nama, jurusan string
	// 	sks           int
	// }{
	// 	"stheven", "teknik elektro", 3,
	// }
	// fmt.Printf("%T\n", mahasiswa)
	// mahasiswa := Mahasiswa{
	// 	"stheven", "elektro", 3,
	// }
	// clone := Mahasiswa{
	// 	nama:    "stheven",
	// 	jurusan: "elektro",
	// 	sks:     3,
	// }

	// fmt.Print(mahasiswa == clone)

	// nama := [...]string{"dwi", "wili", "igun", "iren", "elfan", "caca", "farhan", "dewo", "widiya", "yaya"}
	// elemen := nama[0:5]
	// elemen2 := nama[5:]
	// for i := 1; i < len(elemen); i++ {
	// 	fmt.Println(elemen[i])
	// }
	// for i := 1; i < len(elemen2); i++ {
	// 	fmt.Println(elemen2[i])
	// }

	// var names []string
	// fmt.Println("nama:", reflect.TypeOf(names), len(names), names == nil)
	// var distances []int
	// fmt.Println("jarak:", reflect.TypeOf(distances), len(distances), distances == nil)
	// var data []uint8
	// fmt.Println("data:", reflect.TypeOf(data), len(data), data == nil)
	// var ratios []float64
	// fmt.Println("rasio:", reflect.TypeOf(ratios), len(ratios), ratios == nil)
	// var alives []bool
	// fmt.Println("hidup:", reflect.TypeOf(alives), len(alives), alives == nil)

	/*
		m := movie{
			title: "avengers:end game",
			item:  item{"midnight in paris"},
		}
		fmt.Println(m.title, "&", m.item.title)
			var rows = 5
			for i := 1; i <= rows; i++ {
				for j := 1; j <= i; j++ {
					fmt.Printf("%d ", j)
				}
				fmt.Println()
			phi, nama := 3.14, "Kornelius"
			fmt.Printf("%T\n", phi)
			fmt.Printf("%T\n", nama)

			umurSaya := "20"
			fmt.Printf("Umur ku saat ini %s", umurSaya+" tahun.")
			tingkatanKeahlian := "noob programming"
			tingkatanKeahlian = "pro code"
			fmt.Printf("Aku seorang %s", tingkatanKeahlian)

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

			var num = 5
			for i := 1; i < num; i += 1 {
				fmt.Println("Perulangan ke", i)
			}

			var poin = 2

			if poin > 5 {
				switch {
				case poin == 10:
					fmt.Print("Nilai anda sempurna")
				case poin >5 && poin <10:
					fmt.Print("Nice")
				}
			}
			else {
				if poin <= 5 {
					fmt.Print("Not bad lah!")
				} else if == 3 {
					fmt.Print("Hope u pass the exam")
				} else {
					fmt.Print("Tetap semangat")
					if poin <= 2 && poin > 0{
						fmt.Print("Tetap mencoba, rajin belajar!")
					} else {
						fmt.Print("Nilai anda sangat bad")
					}
				}
			}

			switch {
			case poin == 10:
				fmt.Print("Nilai ujian anda sempurna")
			case poin >= 5 && poin <= 9:
				fmt.Print("Anda lulus ujian dengan nilai", poin)
			case poin < 5:
				fmt.Print("Anda tidak lulus ujian")
			}

			switch poin {
			case 10:
				fmt.Print("Nilai ujian anda sempurna")
				fallthrough
			case 5, 6, 7, 8, 9:
				fmt.Print("Anda lulus ujian dengan nilai ", poin)
				fallthrough
			case 1, 2, 3, 4:
				fmt.Print("Anda tidak lulus ujian")
			}

			var sum int
			for i := 0; i <= 10; i++ {
				sum += i
			}
			fmt.Print("Total penjumlahan ", sum)
			var names [3]string
			names[0] = "stheven"
			names[1] = "erlangga"
			names[2] = "sanjaya"

			// fmt.Println(names[0], names[1], names[2])
			var namaMahasiswa = [...]string{"stheven", "erlangga", "sanjaya", "dwi", "ilham"}

			for i := 0; i < len(namaMahasiswa); i++ {
				fmt.Println("Elemen ", i, ":", namaMahasiswa[i])
			}

			for _, a := range namaMahasiswa {
				fmt.Println(a)
			}
			var sliceMahasiswa = namaMahasiswa[:]
			sliceMahasiswa[2] = "haha"
			fmt.Print(sliceMahasiswa)

			num := []int{2, 3, 4}
			nums := append(num, 7, 8, 9)
			fmt.Print(nums)
	*/

}

func Identity(src string) (dst string) {
	length := len(src)

	if length >= 8 {
		prefix := src[0:4]
		var middle []string

		for i := 0; i < length-8; i++ {
			middle = append(middle, "*")
		}
		suffix := src[length-4 : length]
		dst = fmt.Sprintf("%s%s%s", prefix, strings.Join(middle, ""), suffix)
	}
	return
}
