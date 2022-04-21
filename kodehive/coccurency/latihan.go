package main

import (
	"fmt"
	"time"
	// "time"
)

//channel sebagai parameter
// func sayHello(c chan string, nama string) {
// 	c <- nama
// }

func channelIn(channel chan<- string) {
	channel <- "Stheven"
	time.Sleep(3 * time.Millisecond)
}

func channelOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func main() {
	channel := make(chan string)
	go channelIn(channel)
	go channelOut(channel)
	time.Sleep(3 * time.Millisecond)
	// channelNama := make(chan string)
	// go sayHello(channelNama, "Stheven")
	// fmt.Println(<-channelNama)
}

// var ChannelPenyelam = make(chan string)
// var ChannelPembersihEmas = make(chan string)

// func main() {
// 	timA := [...]string{"gol", "tidak gol", "tidak gol", "gol", "tidak gol", "gol", "gol"}
// 	// timB := []int{0, 1, 1, 0, 1, 0, 0}
// 	var SkorTimB int
// 	for _, skor := range timA {
// 		if skor == "gol" {
// 			SkorTimB += 1
// 			fmt.Println("skor tim B", SkorTimB)
// 		}
// 	}
// go Penyelam()
// go pembersihEmas()
// go menyimpanEmas()
// time.Sleep(3 * time.Millisecond)
// }

// func Penyelam() {
// 	go func() {
// 		hartaKarun := [...]string{"batu", "emas", "kerikil", "emas", "batu", "kerikil", "emas"}
// 		for _, items := range hartaKarun {
// 			if items == "emas" {
// 				fmt.Printf("Berhasil mendapatkan %s\n", items)
// 				ChannelPenyelam <- items
// 			}
// 		}
// 	}()
// 	// dataEmas := <-ChannelPenyelam
// 	// fmt.Println(dataEmas)
// }

// func pembersihEmas() {
// 	for emasKotor := range ChannelPenyelam {
// 		fmt.Println("Berhasil membersihkan", emasKotor)
// 		ChannelPembersihEmas <- emasKotor
// 	}
// }

// func menyimpanEmas() {
// 	for penyimpanan := range ChannelPembersihEmas {
// 		fmt.Println("Berhasil Menyimpan", penyimpanan)
// 	}
// }
