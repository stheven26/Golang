package main

import "fmt"

type Mahasiswa struct {
	Nama, ProgramStudi string
	Semester, Umur     int
	Status             bool
}

func (mahasiswa Mahasiswa) grandOpening(dosen string) {
	fmt.Println("Perkenalkan nama saya adalah", dosen, "selamat datang", mahasiswa.Nama)
}

func anonymousStruct() {
	MataKuliah := struct {
		Nama string
		SKS  int
	}{
		Nama: "Sistem Kendali",
		SKS:  3,
	}
	if MataKuliah.SKS > 2 {
		fmt.Println("Mata kuliah", MataKuliah.Nama, "Jumlah SKS nya adalah", MataKuliah.SKS, "Estimasi waktu kuliah", MataKuliah.SKS*50, "Menit")
	} else {
		fmt.Println("Mata kuliah", MataKuliah.Nama, "Jumlah SKS nya adalah", MataKuliah.SKS, "Estimasi waktu kuliah dibawah 100 menit")
	}
}

func main() {
	mahasiswa := Mahasiswa{
		Nama:         "Stheven",
		ProgramStudi: "Teknik Elektro",
		Semester:     8,
		Umur:         21,
		Status:       true,
	}
	fmt.Println(mahasiswa.Nama)

	mahasiswa1 := Mahasiswa{Nama: "Stheven"}
	mahasiswa1.grandOpening("Ifkar")

	// anonymousStruct()
	// defer fmt.Println("defer")
	// fmt.Println("Biasa")
	// defer fmt.Println("defer 2")
}
