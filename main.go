package main

import "fmt"

const NMAX int = 999

type stTanggal struct {
	date  int
	month int
	year  int
}

type tanggal [31]stTanggal

type arrNamaMahasiswa [NMAX]string
type arrID [NMAX]int
type arrTanggal [NMAX]tanggal

type Bulanan struct {
	ID      arrID
	Nama    arrNamaMahasiswa
	Tanggal arrTanggal
}

type Iuran [12]Bulanan

func main() {
	var lanjut bool

	lanjut = true

	for lanjut {
		fmt.Println("+----------------------------------------------------------+")
		fmt.Println("|   _____   ___   _____ _   _ ______ _     _____  _    _   |")
		fmt.Println("|  /  __ \\ / _ \\ /  ___| | | ||  ___| |   |  _  || |  | |  |")
		fmt.Println("|  | /  \\// /_\\ \\\\ `--.| |_| || |_  | |   | | | || |  | |  |")
		fmt.Println("|  | |    |  _  | `--. \\  _  ||  _| | |   | | | || |/\\| |  |")
		fmt.Println("|  | \\__/\\| | | |/\\__/ / | | || |   | |___\\ \\_/ /\\  /\\  /  |")
		fmt.Println("|   \\____/\\_| |_/\\____/\\_| |_/\\_|   \\_____/\\___/  \\/  \\/   |")
		fmt.Println("|                                                          |")
		fmt.Println("+----------------------------------------------------------+")

		lanjut = false
	}
}
