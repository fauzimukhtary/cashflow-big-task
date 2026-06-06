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
	 *	IS.	-
	 *
	 *	FS.	Tercetak tulisan "CASHFLOW" dengan dekorasi
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

	/*	Prosedur berisi Program utama dengan pilihan menu untuk
	 *	mengakses fitur-fitur utama aplikasi.
	 *
	 *	IS.	-
	 *
	 *	FS.	Program telah selesai berjalan
	 */

	var IuranMaster IuranBulanan
	var JumlahMahasiswa, JumlahTerhapus int
	var LoopInMainMenu, LoopInMenu bool
	var Pilihan int

	LoopInMainMenu = true
	JumlahMahasiswa = 0

	for LoopInMainMenu {
		Pilihan = 0
		fmt.Println("Silakan pilih menu yang diinginkan:")
		fmt.Println("|  1 | Tambah Data")
		fmt.Println("|  2 | Tampilkan Tabel Iuran Kas")
		fmt.Println("|  3 | Update Status Pembayaran")
		fmt.Println("|  4 | Hapus Mahasiswa")
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
				TambahData(&IuranMaster, &JumlahMahasiswa, JumlahTerhapus)
			case 2:
				fmt.Println("============ Tabel Iuran Kas ============")
				TampilkanData(IuranMaster, JumlahMahasiswa, true)
			case 3:
				fmt.Println("============== Update  Kas ==============")
				UpdateStatusBayar(&IuranMaster, JumlahMahasiswa)
			case 4:
				fmt.Println("============== Hapus  Data ==============")
				HapusData(&IuranMaster, &JumlahMahasiswa, &JumlahTerhapus)
			case 99:
				fmt.Println("Terima kasih telah menggunakan CashFlow!")
				LoopInMainMenu = false
			}

			LoopInMenu = !LoopInMenu
			fmt.Println()
		}
	}
}

func TambahData(Tabel *IuranBulanan, NMahasiswa *int, NTerhapus int) {

	/*	Prosedur untuk menambah mahasiswa pada tabel iuran dan secara bawaan
	 *	mengisi seluruh status pembayaran dengan string "-----" (belum lunas)
	 *
	 *	IS.	Terdefinisi Tabel bertipe data IuranBulanan,
	 *		NMahasiswa (jumlah mahasiswa yang ada di tabel),
	 *		NTerhapus (jumlah mahasiswa yang telah terhapus dari tabel)
	 *
	 *	FS. Data Iuran Mahasiswa terisi sebanyak 1, dengan ID secara otomatis
	 *		berurut setelah ID terakhir yang pernah ada pada tabel,
	 *		Seluruh status pembayaran terisi string "-----"
	 */

	var i int
	var TMPNama string

	fmt.Print("Masukkan Nama Mahasiswa Baru (Tanpa Spasi): ")
	fmt.Scan(&TMPNama)

	Tabel.Nama[*NMahasiswa] = TMPNama
	Tabel.ID[*NMahasiswa] = *NMahasiswa + NTerhapus + 1
	for i = 0; i < HARI; i++ {
		Tabel.Tanggal[*NMahasiswa][i] = "-----"
	}
	*NMahasiswa++

	fmt.Printf("Mahasiswa: %v, berhasil ditambahkan ke Sistem Iuran\n", TMPNama)
}

func TampilkanData(Tabel IuranBulanan, NMahasiswa int, LoopDate bool) {

	/*	Prosedur untuk menampilkan data pada rentang tanggal tertentu
	 *
	 *	IS.	Terdefinisi Tabel bertipe data IuranBulanan,
	 *		NMahasiswa (jumlah mahasiswa yang ada di tabel),
	 *		LoopDate untuk menampilkan pilihan rentang tanggal jika bernilai
	 *		true, dan menampilkan pratinjau (tanggal 1-10) jika false
	 *
	 *	FS. Tercetak di layar Tabel Iuran pada rentang tanggal tertentu
	 *		sesuai masukan pengguna
	 */

	var i, j int
	var Mulai, Akhir int
	var PilihanTanggal int
	var LoopInDateMenu bool

	if NMahasiswa == 0 {
		fmt.Println("********* Database masih kosong *********")
	} else {

		if LoopDate {
			fmt.Println("Pilih rentang tanggal:")
			fmt.Println("1. 1-10")
			fmt.Println("2. 11-20")
			fmt.Println("3. 21-31")
		}

		LoopInDateMenu = LoopDate

		for LoopInDateMenu {
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&PilihanTanggal)

			switch PilihanTanggal {
			default:
				fmt.Println("Pilihan Anda tidak valid")
				LoopInDateMenu = false
			case 1:
				Mulai = 1
				Akhir = 10
			case 2:
				Mulai = 11
				Akhir = 20
			case 3:
				Mulai = 21
				Akhir = 31
			}

			LoopInDateMenu = !LoopInDateMenu
			fmt.Println()
		}

		if !LoopDate {
			Mulai = 1
			Akhir = 10
			fmt.Println("Preview data...")
		}

		fmt.Printf("| ID |       %-21v|%43v\n", "Nama Mahasiswa", "Tanggal")
		fmt.Print("|    |                            |")

		for i = Mulai; i <= Akhir; i++ {
			fmt.Printf("   %-2d  |", i)
		}

		for i = 0; i < NMahasiswa; i++ {
			fmt.Printf("\n| %-2d | %-26v |", Tabel.ID[i], Tabel.Nama[i])
			for j = Mulai - 1; j <= Akhir-1; j++ {
				fmt.Printf(" %-5v |", Tabel.Tanggal[i][j])
			}
		}
		fmt.Println()

	}
}

func UpdateStatusBayar(Tabel *IuranBulanan, NMahasiswa int) {

	/*	Prosedur untuk memperbarui status lunas iuran pada tanggal-tanggal
	 *	tertentu berdasarkan ID
	 *
	 *	IS.	Terdefinisi Tabel bertipe data IuranBulanan,
	 *		NMahasiswa (jumlah mahasiswa yang ada di tabel),
	 *
	 *	FS. Keterangan pembayaran mahasiswa pada tanggal-tanggal yang
	 *		telah dimasukkan oleh pengguna berubah menjadi "LUNAS".
	 */

	var IDToUpdate int
	var IdxOfID int
	var Date int

	TampilkanData(*Tabel, NMahasiswa, false)

	if NMahasiswa != 0 {
		fmt.Print("\nPilih ID yang akan diupdate: ")
		fmt.Scan(&IDToUpdate)

		IdxOfID = CariBerdasarID(*Tabel, NMahasiswa, IDToUpdate)

		if IdxOfID == -1 {
			fmt.Println("ID", IDToUpdate, "tidak ditemukan")
		} else {
			fmt.Printf(
				"Anda akan mengupdate data untuk: %v dengan ID %d \n\n",
				Tabel.Nama[IdxOfID],
				Tabel.ID[IdxOfID],
			)
			fmt.Println("Ketikkan tanggal-tanggal yang ingin diupdate menjadi LUNAS")
			fmt.Println("Note: Masukkan angka-angka yang dipisahkan oleh spasi")
			fmt.Println("Note: Masukkan angka diluar jangkauan untuk menghentikan input")
			fmt.Println("*) contoh: 12 14 19 31 -1")

			fmt.Print("\nTanggal yang akan diupdate: ")
			fmt.Scan(&Date)

			for Date > 0 && Date <= HARI {
				Tabel.Tanggal[IdxOfID][Date-1] = "LUNAS"
				fmt.Scan(&Date)
			}

			fmt.Println("Terdeteksi sentinel/indeks invalid:", Date, ":: input dihentikan")
			fmt.Println("Data berhasil diupdate")
		}

	}
}

func HapusData(Tabel *IuranBulanan, NMahasiswa, NTerhapus *int) {

	/*	Prosedur untuk menampilkan tabel iuran yang memperbolehkan
	 *	pengguna untuk menghapus Data dari Tabel berdasarkan ID
	 *
	 *	IS.	Terdefinisi Tabel bertipe data IuranBulanan,
	 *		NMahasiswa (jumlah mahasiswa yang ada di tabel),
	 *		NTerhapus (jumlah mahasiswa yang telah terhapus dari tabel)
	 *
	 *	FS. Data Iuran Mahasiswa pada indeks tempat ID berada telah terhapus
	 *		dan data setelahnya digeser satu indeks lebih maju;
	 *		NTerhapus bertambah dan NMahasiswa berkurang
	 */

	var i, j, IDToDelete int
	var IdxOfID int
	var Loop bool
	var ConfirmHapus byte

	if *NMahasiswa == 0 {
		fmt.Println("********* Database masih kosong *********") // <- Tercetak bila tabel masih kosong
	} else {
		fmt.Print("| ID |       Nama Mahasiswa       |\n")
		for i = 0; i < *NMahasiswa; i++ {
			fmt.Printf("| %-2d | %-26v |\n", Tabel.ID[i], Tabel.Nama[i])
		}

		fmt.Print("\nPilih ID yang akan dihapus: ")
		fmt.Scan(&IDToDelete)
		fmt.Scanln()

		IdxOfID = CariBerdasarID(*Tabel, *NMahasiswa, IDToDelete)

		if IdxOfID == -1 {
			// prosedur tidak melakukan perubahan apapun
			fmt.Println("ID", IDToDelete, "tidak ditemukan")
		} else {

			// konfirmasi penghapusan
			fmt.Printf(
				"\nAnda akan menghapus mahasiswa dengan nama: %v dengan ID %d",
				Tabel.Nama[IdxOfID],
				Tabel.ID[IdxOfID],
			)

			Loop = true
			for Loop {
				fmt.Print("\nApakah ingin melanjutkan (Y/N): ")
				fmt.Scanf("%c", &ConfirmHapus)
				switch ConfirmHapus {
				case 'y', 'Y', 'n', 'N':
					Loop = false
				default:
					fmt.Println("Input invalid")
				}
			}

			// hasil penghapusan
			switch ConfirmHapus {
			case 'y', 'Y':
				for i = IdxOfID; i < *NMahasiswa; i++ {
					Tabel.ID[i] = Tabel.ID[i+1]
					Tabel.Nama[i] = Tabel.Nama[i+1]
					for j = 0; j < HARI; j++ {
						Tabel.Tanggal[i][j] = Tabel.Tanggal[i+1][j]
					}
				}
				fmt.Println("Data berhasil dihapus")
				*NTerhapus++
				*NMahasiswa--
			case 'n', 'N':
				fmt.Println("Data batal dihapus")
			}
		}
	}
}

func CariBerdasarID(Tabel IuranBulanan, NMahasiswa, IDTarget int) int {

	/*	Mencari indeks dari ID Target menggunakan algoritma binary search
	 *	Fungsi mengembalikan indeks tenpat ID berada, dan -1 jika tidak ditemukan
	 */

	var BatasKiri, Tengah, BatasKanan int
	var IndeksDicari int

	BatasKiri = 0
	BatasKanan = NMahasiswa - 1
	IndeksDicari = -1

	for BatasKiri <= BatasKanan && IndeksDicari == -1 {
		Tengah = (BatasKiri + BatasKanan) / 2
		if IDTarget == Tabel.ID[Tengah] {
			IndeksDicari = Tengah
		} else if IDTarget < Tabel.ID[Tengah] {
			BatasKanan = Tengah - 1
		} else {
			BatasKiri = Tengah + 1
		}
	}

	return IndeksDicari
}
