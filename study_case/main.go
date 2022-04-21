package main

import (
	"fmt"
	"time"
)

// var ChannelPenembak = make(chan string)
// var ChannelPenangkap = make(chan string)
var ChannelMemesan = make(chan string)
var ChannelMembuat = make(chan string)
var ChannelMengantar = make(chan string)

func main() {
	// go menembak()
	// go menangkap()
	// go memasukkan()
	go memesanPesanan()
	go membuatPesanan()
	go mengantarPesanan()
	go pelangganMenerima()
	time.Sleep(10 * time.Millisecond)
}

// func menembak() {
// 	sasaranTembakan := [...]string{"burung", "pohon", "awan", "burung", "pohon", "burung", "awan"}
// 	for _, sasaran := range sasaranTembakan {
// 		if sasaran == "burung" {
// 			fmt.Println("Berhasil menembak", sasaran)
// 			ChannelPenembak <- sasaran
// 		}
// 	}
// }

// func menangkap() {
// 	for menangkap := range <- ChannelPenembak {
// 		fmt.Println("Berhasil menangkap", menangkap)
// 		ChannelPenangkap <- menangkap
// 	}
// }

// func memasukkan() {
// 	for memasukkan := range <- ChannelPenangkap {
// 		fmt.Println("Berhasil memasukkan", memasukkan)
// 	}
// }

func memesanPesanan() {
	listPesanan := [...]string{"kopi", "air", "susu", "kopi", "air", "kopi", "kopi", "air", "susu"}

	for _, memesan := range listPesanan {
		if memesan == "kopi" {
			fmt.Println("Silahkan tunggu, pesanan kopi anda sedang dibuatkan!")
			ChannelMemesan <- memesan
		} else if memesan == "susu" {
			fmt.Println("Silahkan tunggu, pesanan susu anda sedang dibuatkan!")
			ChannelMemesan <- memesan
		} else {
			fmt.Println("Tidak menerima pesanan air!")
		}
	}
}

func membuatPesanan() {
	for membuat := range ChannelMemesan {
		fmt.Println("Telah selesai membuat", membuat)
		ChannelMembuat <- membuat
	}
}

func mengantarPesanan() {
	for _, mengantar := range ChannelMembuat {
		fmt.Printf("%s sedang diantarkan ke pelanggan\n", mengantar)
		ChannelMengantar <- mengantar
	}
}

func pelangganMenerima() {
	for pelangganMenerimaPesanan := range ChannelMengantar {
		fmt.Println("Pelanggan telah menerima", pelangganMenerimaPesanan)
	}
}
