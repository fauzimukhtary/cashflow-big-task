package main

import "fmt"

const NMAX int = 999
const HARI int = 31

type TanggalBayar [HARI]string

type ArrNamaMahasiswa [NMAX]string
type ArrID [NMAX]int
type ArrTanggal [NMAX]TanggalBayar

type IuranBulanan struct {
	ID      ArrID
	Nama    ArrNamaMahasiswa
	Tanggal ArrTanggal
}

func main() {
	MainMenuDisplay()
	Program()
}

func MainMenuDisplay() {
	/*
	 *	Prosedur berisi output untuk dekorasi pada menu utama
	 */

	fmt.Println("+----------------------------------------------------------+")
	fmt.Println("|   _____   ___   _____ _   _ ______ _     _____  _    _   |")
	fmt.Println("|  /  __ \\ / _ \\ /  ___| | | ||  ___| |   |  _  || |  | |  |")
	fmt.Println("|  | /  \\// /_\\ \\\\ `--.| |_| || |_  | |   | | | || |  | |  |")
	fmt.Println("|  | |    |  _  | `--. \\  _  ||  _| | |   | | | || |/\\| |  |")
	fmt.Println("|  | \\__/\\| | | |/\\__/ / | | || |   | |___\\ \\_/ /\\  /\\  /  |")
	fmt.Println("|   \\____/\\_| |_/\\____/\\_| |_/\\_|   \\_____/\\___/  \\/  \\/   |")
	fmt.Println("|                                                          |")
	fmt.Println("+----------------------------------------------------------+")
	fmt.Println()
	fmt.Println("Selamat Datang di CashFlow - Aplikasi Manajemen Kas Kelas")
}

func Program() {
	/*
	 *	Prosedur berisi Program utama dengan pilihan menu yang dapat
	 *	mengakses fitur-fitur utama aplikasi.
	 */

	var IuranMaster IuranBulanan
	var JumlahMahasiswa int
	var LoopInMainMenu, LoopInMenu bool
	var Pilihan int

	LoopInMainMenu = true

	for LoopInMainMenu {
		Pilihan = 0
		fmt.Println("Silakan pilih menu yang diinginkan:")
		fmt.Println("|  1 | Tambah Data")
		fmt.Println("|  2 | Tampilkan Tabel Iuran Kas")
		fmt.Println("|  3 | Update Status Pembayaran")
		fmt.Println("|  4 | Hapus Mahasiswa")
		fmt.Println("|  9 | Insert Data Dummy")
		fmt.Println("| 99 | Keluar Aplikasi")
		fmt.Println()

		LoopInMenu = true

		for LoopInMenu {
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&Pilihan)

			switch Pilihan {
			default:
				fmt.Println("Pilihan Anda tidak valid")
				LoopInMenu = false
			case 1:
				fmt.Println("============== Tambah Data ==============")
				TambahData(&IuranMaster, &JumlahMahasiswa)
			case 2:
				fmt.Println("============ Tabel Iuran Kas ============")
				TampilkanData(IuranMaster, JumlahMahasiswa)
			case 3:
			case 4:
			case 5:
			case 99:
				fmt.Println("Terima kasih telah menggunakan CashFlow!")
				LoopInMainMenu = false
			}

			LoopInMenu = !LoopInMenu
			fmt.Println()
		}
	}
}

func TambahData(Tabel *IuranBulanan, NMahasiswa *int) {
	/*
	 * 	IS.	terdefinisi sembarang Tabel bertipe array data IuranBulanan,
	 		dan NMahasiswa bertipe integer
	 * 	FS.	Tabel terisi nama mahasiswa baru pada indeks NMahasiswa saat itu.
	 */
	var i int
	var TMPNama string

	fmt.Print("Masukkan Nama Mahasiswa Baru (Tanpa Spasi): ")
	fmt.Scan(&TMPNama)

	Tabel.Nama[*NMahasiswa] = TMPNama
	Tabel.ID[*NMahasiswa] = *NMahasiswa + 1
	for i = 0; i < HARI; i++ {
		Tabel.Tanggal[*NMahasiswa][i] = "-----"
	}
	*NMahasiswa++

	fmt.Printf("Mahasiswa: %v, berhasil ditambahkan ke Sistem Iuran\n", TMPNama)
}

func TampilkanData(Tabel IuranBulanan, NMahasiswa int) {
	var i, j int
	var mulai, akhir int
	var PilihanTanggal int
	var LoopInDateMenu bool

	if NMahasiswa == 0 {
		fmt.Println("********* Database masih kosong *********")
	} else {

		fmt.Println("Pilih rentang tanggal:")
		fmt.Println("1. 1-10")
		fmt.Println("2. 11-20")
		fmt.Println("3. 21-31")

		LoopInDateMenu = true

		for LoopInDateMenu {
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&PilihanTanggal)

			switch PilihanTanggal {
			default:
				fmt.Println("Pilihan Anda tidak valid")
				LoopInDateMenu = false
			case 1:
				mulai = 1
				akhir = 10
			case 2:
				mulai = 11
				akhir = 20
			case 3:
				mulai = 21
				akhir = 31
			}

			LoopInDateMenu = !LoopInDateMenu
			fmt.Println()
		}

		fmt.Print("| ID |       Nama Mahasiswa       |                                    Tanggal\n")
		fmt.Print("|    |                            |")

		for i = mulai; i <= akhir; i++ {
			fmt.Printf("   %-2d  |", i)
		}

		for i = 0; i < NMahasiswa; i++ {
			fmt.Printf("\n| %-2d | %-26v |", Tabel.ID[i], Tabel.Nama[i])
			for j = mulai - 1; j <= akhir-1; j++ {
				fmt.Printf(" %-5v |", Tabel.Tanggal[i][j])
			}
		}
		fmt.Println()

	}
}
