package main

const NMAX int = 999

type arrNamaMahasiswa [NMAX]string
type arrID [NMAX]int

type Master struct {
	ID      arrID
	Nama    arrNamaMahasiswa
	tanggal arrTanggal
}

type arrTanggal [NMAX]tanggal

type tanggal struct {
	date  int
	month int
	year  int
}

func main() {

}
