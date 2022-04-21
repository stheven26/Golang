package main

import (
	"fmt"
	"time"
)

// func perulanganAngka(angka int) {
// 	fmt.Printf("Perulangan ke %d\n", angka)
// }
var itemsChannel = make(chan string)
var cleanedItemChannel = make(chan string)

func main() {
	TestCreateChannel()
	// for i := 1; i <= 10; i++ {
	// 	go perulanganAngka(i)
	// }
	items := [10]string{
		"batu",
		"emas",
		"kerikil",
		"besi",
		"emas",
		"emas",
		"emas",
		"batu",
		"ikan",
		"kerang",
	}

	go menyelam(items)
	go membersihkan()
	go menyimpan()
	time.Sleep(5 * time.Second)
}

func TestCreateChannel() {
	channel := make(chan string)

	go func() {
		ucapanSelamat := "Selamat Stheven!"
		channel <- ucapanSelamat
		fmt.Println("Selesai memasukkan data ke channel")
	}()
	data := <-channel
	fmt.Println(data)
}

func menyelam(items [10]string) {
	for _, item := range items {
		if item == "emas" {
			fmt.Println("Berhasil mendapatkan", item)
			itemsChannel <- item
		}
	}
}

func membersihkan() {
	for itemKotor := range itemsChannel {
		fmt.Println("Berhasil membersihkan", itemKotor)
		cleanedItemChannel <- itemKotor
	}
}

func menyimpan() {
	for cleanedItem := range cleanedItemChannel {
		fmt.Println("Berhasil menyimpan", cleanedItem)
	}
}
