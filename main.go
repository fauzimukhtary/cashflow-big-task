// FUNCTION
// PROCEDURE
// RECURSIVE
// FIND MIN/MAX
// SEQUENTIAL SEARCH
// BINARY SEARCH ✔
// SELECTION SORT ✔
// INSERTION SORT

package main

import "fmt"

const NMAX int = 999
const HARI int = 31

type ArrTanggal [HARI]string

type DataIuran struct {
	ID      int
	Nama    string
	Tanggal ArrTanggal
}

type IuranBulanan [NMAX]DataIuran

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
	var SortedByID bool
	var NominalKas int

	LoopInMainMenu = true
	JumlahMahasiswa = 0
	SortedByID = true

	for LoopInMainMenu {
		Pilihan = 0
		fmt.Println("Silakan pilih menu yang diinginkan:")
		fmt.Println("| 1  | Tambah Data")
		fmt.Println("| 2  | Tampilkan Tabel Iuran Kas")
		fmt.Println("| 3  | Update Status Pembayaran")
		fmt.Println("| 4  | Urutkan & Total Pembayaran")
		fmt.Println("| 5  | Status Bayar Per Tanggal")
		fmt.Println("| 6  | Hapus Mahasiswa")
		fmt.Println("| 7  | Ubah Nominal Kas")
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
				fmt.Println("============== Urutkan Data ==============")
				UrutkanLaporanByPembayaran(IuranMaster, JumlahMahasiswa, &SortedByID, NominalKas)
			case 5:
				fmt.Println("============== Belum Bayar ==============")
				BelumBayarPerTanggal(IuranMaster, JumlahMahasiswa)
			case 6:
				fmt.Println("============== Hapus  Data ==============")
				HapusData(&IuranMaster, &JumlahMahasiswa, &JumlahTerhapus)
			case 7:
				fmt.Println("=========== Ubah Nominal Kas ============")
				fmt.Print("Masukkan nominal dalam rupiah: ")
				fmt.Scan(&NominalKas)
				fmt.Println("Berhasil update nominal")
			case 99:
				fmt.Println("+-----------------------------------------------------------------------------------+")
				fmt.Println("|                                                                                   |")
				fmt.Println("|  _____ _____ ____  ___ __  __     _     _  __     _     ____ ___ _   _     _      |")
				fmt.Println("| |_   _| ____|  _ ||_ _|  |/  |   / |   | |/ /    / |   / ___|_ _| | | |   | |     |")
				fmt.Println("|   | | |  _| | |_) || || ||/| |  / _ |  | ' /    / _ |  |___ || || |_| |   | |     |")
				fmt.Println("|   | | | |___|  _ < | || |  | | / ___ | | . |   / ___ |  ___) | ||  _  |   |_|     |")
				fmt.Println("|   |_| |_____|_| |_|___|_|  |_|/_/   |_||_||_| /_/   |_||____/___|_| |_|   (_)     |")
				fmt.Println("|                                                                                   |")
				fmt.Println("+-----------------------------------------------------------------------------------+")
				fmt.Println("           Terima kasih telah menggunakan CashFlow - Sampai Jumpa!")
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
	fmt.Scan(&Tabel[*NMahasiswa].Nama)

	Tabel[*NMahasiswa].ID = *NMahasiswa + NTerhapus + 1
	for i = 0; i < HARI; i++ {
		Tabel[*NMahasiswa].Tanggal[i] = "-----"
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
			fmt.Printf("\n| %-2d | %-26v |", Tabel[i].ID, Tabel[i].Nama)
			for j = Mulai - 1; j <= Akhir-1; j++ {
				fmt.Printf(" %-5v |", Tabel[i].Tanggal[j])
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
				Tabel[IdxOfID].Nama,
				Tabel[IdxOfID].ID,
			)
			fmt.Println("Ketikkan tanggal-tanggal yang ingin diupdate menjadi LUNAS")
			fmt.Println("Note: Masukkan angka-angka yang dipisahkan oleh spasi")
			fmt.Println("Note: Masukkan angka diluar jangkauan untuk menghentikan input")
			fmt.Println("*) contoh: 12 14 19 31 -1")

			fmt.Print("\nTanggal yang akan diupdate: ")
			fmt.Scan(&Date)

			for Date > 0 && Date <= HARI {
				Tabel[IdxOfID].Tanggal[Date-1] = "LUNAS"
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
			fmt.Printf("| %-2d | %-26v |\n", Tabel[i].ID, Tabel[i].Nama)
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
				Tabel[IdxOfID].Nama,
				Tabel[IdxOfID].ID,
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
					Tabel[i].ID = Tabel[i+1].ID
					Tabel[i].Nama = Tabel[i+1].Nama
					for j = 0; j < HARI; j++ {
						Tabel[i].Tanggal[j] = Tabel[i+1].Tanggal[j]
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
		if IDTarget == Tabel[Tengah].ID {
			IndeksDicari = Tengah
		} else if IDTarget < Tabel[Tengah].ID {
			BatasKanan = Tengah - 1
		} else {
			BatasKiri = Tengah + 1
		}
	}

	return IndeksDicari
}

/*
 * wlm
 */

func BelumBayarPerTanggal(Tabel IuranBulanan, NMahasiswa int) {

	/* 	Prosedur untuk menampilkan daftar nama mahasiswa yang belum
	 *	membayar iuran pada tanggal tertentu.
	 *
	 *	IS. Terdefinisi Tabel bertipe data IuranBulanan dan NMahasiswa
	 *
	 *	FS. Menampilkan daftar nama mahasiswa yang memiliki status "-----"
	 *  pada tanggal yang diinputkan pengguna.
	 */

	var Date, i int
	var Ketemu bool
	var LoopInDateChoice bool

	if NMahasiswa == 0 {
		fmt.Println("********* Database masih kosong *********")
	} else {
		LoopInDateChoice = true

		for LoopInDateChoice {
			fmt.Print("Masukkan tanggal yang ingin dicek (1-31): ")
			fmt.Scan(&Date)
			if Date < 1 || Date > HARI {
				fmt.Print("Tanggal tidak valid (Gunakan rentang 1-31).\n\n")
			} else {
				LoopInDateChoice = false
			}
		}

		fmt.Printf("\nMahasiswa yang BELUM BAYAR pada tanggal %d:\n", Date)
		fmt.Println("-------------------------------------------")

		// Mencetak daftar mahasiswa yg belum bayar pada tanggal tersebut
		Ketemu = false
		for i = 0; i < NMahasiswa; i++ {
			// Mengecek jika status masih "-----"
			if Tabel[i].Tanggal[Date-1] == "-----" {
				fmt.Printf("- [%d] %s\n", Tabel[i].ID, Tabel[i].Nama)
				Ketemu = true
			}
		}

		if !Ketemu {
			fmt.Println("Hebat! Semua mahasiswa sudah bayar pada tanggal ini.")
		}
		fmt.Println("-------------------------------------------")
	}
}

func HitungTotalBayar(Data ArrTanggal) int {

	/*	Fungsi utama untuk menghitung berapa kali status "LUNAS" muncul
	 *	dalam array TanggalBayar milik seorang mahasiswa.
	 *  Memanggil fungsi rekursif lain
	 */

	var i int

	i = 0
	return HitungTotalBayarRecursive(Data, i)
}

func HitungTotalBayarRecursive(Date ArrTanggal, i int) int {

	/*	Fungsi secara rekursif menghitung berapa kali status "LUNAS" muncul
	 *	dalam array TanggalBayar milik seorang mahasiswa.
	 */

	if i == HARI-1 {
		if Date[i] == "LUNAS" {
			return 1
		} else {
			return 0
		}
	} else {
		if Date[i] == "LUNAS" {
			return 1 + HitungTotalBayarRecursive(Date, i+1)
		} else {
			return HitungTotalBayarRecursive(Date, i+1)
		}
	}
}

func UrutkanLaporanByPembayaran(Tabel IuranBulanan, NMahasiswa int, SortedByID *bool, Nominal int) {

	/* 	Prosedur untuk mengurutkan data mahasiswa berdasarkan total pembayaran.
	 *	Jika total pembayaran sama, maka diurutkan berdasarkan ID terkecil.
	 *	Menggunakan algoritma SELECTION SORT sesuai modul dosen.
	 *
	 *	IS. Terdefinisi Tabel bertipe IuranBulanan dan NMahasiswa.
	 *
	 *	FS. Tabel terurut (Ascending/Descending) berdasarkan total bayar.
	 *	    Jika ada total yang sama, ID terkecil akan diutamakan.
	 *	    Hasil akhir tercetak di layar.
	 */

	var pilihan, i int
	var LoopInAscDesc bool

	if NMahasiswa == 0 {
		fmt.Println("********* Database masih kosong *********")
	} else if Nominal == 0 {
		fmt.Println("****** Nominal Kas Belum Di-Setting *****")
	} else {

		fmt.Println("Pilih urutan tampilan:")
		fmt.Println("1. Total Bayar Terkecil ke Terbanyak (Ascending)")
		fmt.Println("2. Total Bayar Terbanyak ke Terkecil (Descending)")

		LoopInAscDesc = true

		for LoopInAscDesc {
			fmt.Printf("\nPilihan Anda: ")
			fmt.Scan(&pilihan)

			if pilihan != 1 && pilihan != 2 {
				fmt.Println("Pilihan tidak valid.")
			} else {
				LoopInAscDesc = false
			}
		}

		SelectionSortByPembayaran(&Tabel, NMahasiswa, pilihan)
		*SortedByID = false

		// TAMPILKAN HASIL
		if pilihan == 1 {
			fmt.Println("**** Laporan:  Terkecil ke Terbanyak ****")
		} else {
			fmt.Println("**** Laporan:  Terbanyak ke Terkecil ****")
		}

		fmt.Printf("| %-3v | %-30v | %-12v |\n", "ID", "Nama Mahasiswa", "Total Bayar")
		fmt.Println("--------------------------------------------")

		// Cetak Hasil
		for i = 0; i < NMahasiswa; i++ {
			fmt.Printf("| %-3d | %-30s | %-12d |\n",
				Tabel[i].ID,
				Tabel[i].Nama,
				Nominal*HitungTotalBayar(Tabel[i].Tanggal),
			)
		}

	}

}

func SelectionSortByPembayaran(Tabel *IuranBulanan, NMahasiswa int, choice int) {

	/*
	 *
	 */

	var i, j, Min, Max int
	var TotalJ, TotalMin, TotalMax int
	var temp DataIuran

	if choice == 1 {

		for i = 0; i < NMahasiswa-1; i++ {

			Min = i
			TotalMin = HitungTotalBayar(Tabel[Min].Tanggal) // hitung berapa kali mhs pada indeks Min membayar

			for j = i + 1; j < NMahasiswa; j++ {

				TotalJ = HitungTotalBayar(Tabel[j].Tanggal) // hitung berapa kali mhs pada indeks j membayar

				if TotalJ < TotalMin {

					Min = j
					TotalMin = HitungTotalBayar(Tabel[Min].Tanggal)

				} else if TotalMin == TotalJ && Tabel[j].ID < Tabel[Min].ID { // Bila Total Sama, ID Yg lebih kecil dipilih

					Min = j
					TotalMin = HitungTotalBayar(Tabel[Min].Tanggal)

				}

			}

			// Swapping
			temp = Tabel[Min]
			Tabel[Min] = Tabel[i]
			Tabel[i] = temp
		}

	} else {

		for i = 0; i < NMahasiswa-1; i++ {

			Max = i
			TotalMax = HitungTotalBayar(Tabel[Max].Tanggal) // hitung berapa kali mhs pada indeks Max membayar

			for j = i + 1; j < NMahasiswa; j++ {

				TotalJ = HitungTotalBayar(Tabel[j].Tanggal) // hitung berapa kali mhs pada indeks j membayar

				if TotalJ < TotalMax {

					Max = j
					TotalMax = HitungTotalBayar(Tabel[Max].Tanggal)

				} else if TotalMax == TotalJ && Tabel[j].ID < Tabel[Max].ID { // Bila Total Sama, ID Yg lebih kecil dipilih

					Max = j
					TotalMax = HitungTotalBayar(Tabel[Max].Tanggal)

				}

			}

			// Swapping
			temp = Tabel[Max]
			Tabel[Max] = Tabel[i]
			Tabel[i] = temp
		}

	}

}

func InsertionSortByID(Tabel *IuranBulanan, NMahasiswa int) {
	var i, j int
	var temp DataIuran

	for i = 1; i < NMahasiswa; i++ {
		j = i
		temp = Tabel[j]
		for j > 1 && temp.ID < Tabel[j-1].ID {
			Tabel[j] = Tabel[j-1]
			j--
		}
		Tabel[j] = temp
	}
}

func LaporanKeuangan(Tabel *IuranBulanan, NMahasiswa int) {

}
