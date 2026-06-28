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

	/*	Prosedur berisi output untuk dekorasi pada menu utama
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
		Pilihan = 99
		fmt.Println("Silakan pilih menu yang diinginkan:")
		fmt.Println("| 1  | Tambah Data")
		fmt.Println("| 2  | Tampilkan Tabel Iuran Kas")
		fmt.Println("| 3  | Update Status Pembayaran")
		fmt.Println("| 4  | Urutkan & Total Pembayaran")
		fmt.Println("| 5  | Status Bayar Per Tanggal")
		fmt.Println("| 6  | Hapus Mahasiswa")
		fmt.Println("| 7  | Ubah Nominal Kas")
		fmt.Println("| 8  | Cari (ID/Nama)")
		fmt.Println("| 9  | Laporan Keuangan (Per Tanggal & Keseluruhan)")
		fmt.Println("| 0  | Keluar Aplikasi")
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
			case 8:
				fmt.Println("=============== Pencarian ===============")
				MenuCariMahasiswa(IuranMaster, JumlahMahasiswa, NominalKas)
			case 9:
				fmt.Println("=========== Laporan Total Kas ===========")
				LaporanKeuangan(IuranMaster, JumlahMahasiswa, NominalKas)
			case 0:
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
	 *	Ketentuan: ID yang telah dihapus tidak bisa digunakan
	 *	lagi pada penambahan data selanjutnya.
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

func MenuCariMahasiswa(Tabel IuranBulanan, NMahasiswa int, Nominal int) {

	/*	Prosedur menu pencarian. Pengguna dapat memilih untuk mencari
	 *	mahasiswa berdasarkan ID (Binary Search), berdasarkan Nama
	 *	(Sequential Search), atau mencari mahasiswa dengan pembayaran
	 *	ekstrim (Min/Max).
	 *
	 *	IS. Terdefinisi Tabel IuranBulanan, NMahasiswa, dan Nominal.
	 *	FS. Menampilkan hasil pencarian sesuai metode yang dipilih.
	 */

	var pilihan, IdxHasil, IDTarget int
	var NamaTarget string
	var LoopInSrcMethod bool

	if NMahasiswa == 0 {
		fmt.Println("********* Database masih kosong *********")
	} else if Nominal == 0 {
		fmt.Println("****** Nominal Kas Belum Di-Setting *****")
	} else {

		fmt.Println("Pilih metode pencarian:")
		fmt.Println("1. Cari berdasarkan ID (Binary Search)")
		fmt.Println("2. Cari berdasarkan Nama (Sequential Search)")
		fmt.Println("3. Cari Pembayaran Tertinggi/Terendah (Min/Max)")

		LoopInSrcMethod = true

		for LoopInSrcMethod {
			fmt.Print("\nPilihan Anda: ")
			fmt.Scan(&pilihan)

			if pilihan < 1 || pilihan > 3 {
				fmt.Println("Pilihan tidak valid")
			} else {
				LoopInSrcMethod = false
			}
		}

		switch pilihan {
		case 1:
			fmt.Print("Masukkan ID yang dicari: ")
			fmt.Scan(&IDTarget)
			IdxHasil = CariBerdasarID(Tabel, NMahasiswa, IDTarget)
			if IdxHasil == -1 {
				fmt.Println("ID", IDTarget, "Tidak ditemukan")
			} else {
				TampilDataIndex(Tabel, IdxHasil, Nominal)
			}
		case 2:
			fmt.Print("Masukkan Nama yang dicari (Tanpa Spasi): ")
			fmt.Scan(&NamaTarget)
			IdxHasil = CariBerdasarNama(Tabel, NMahasiswa, NamaTarget)
			if IdxHasil == -1 {
				fmt.Println("Nama", NamaTarget, "Tidak ditemukan")
			} else {
				TampilDataIndex(Tabel, IdxHasil, Nominal)
			}
		case 3:
			CariEkstrim(Tabel, NMahasiswa, Nominal)
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

func CariBerdasarNama(Tabel IuranBulanan, NMahasiswa int, NamaTarget string) int {

	/*	Mencari indeks mahasiswa berdasarkan Nama menggunakan algoritma
	 *	SEQUENTIAL SEARCH (linear search), dilakukan dari indeks pertama
	 *	hingga ditemukan kecocokan nama atau mencapai akhir tabel.
	 *
	 *	IS. Terdefinisi Tabel IuranBulanan dan NMahasiswa.
	 *	FS. Mengembalikan indeks mahasiswa dengan Nama yang cocok,
	 *		atau -1 jika tidak ditemukan.
	 */

	var i int
	var IndeksDicari int

	i = 0
	IndeksDicari = -1

	for i < NMahasiswa && IndeksDicari == -1 {
		if Tabel[i].Nama == NamaTarget {
			IndeksDicari = i
		}
		i++
	}

	return IndeksDicari
}

func CariEkstrim(Tabel IuranBulanan, NMahasiswa int, Nominal int) {

	/* 	Prosedur untuk menampilkan mahasiswa dengan pembayaran
	 *	tertinggi (Max) dan terendah (Min).
	 *
	 *	IS. Terdefinisi Tabel IuranBulanan, NMahasiswa, dan Nominal.
	 *	FS. Menampilkan dua Mahasiswa dengan pembayaran tertinggi dan terendah
	 */

	var i, IdxExt int

	IdxExt = 0

	// ALGORITMA FIND MAX (Paling Banyak Lunas)
	for i = 1; i < NMahasiswa; i++ {
		if HitungTotalBayar(Tabel[i].Tanggal) > HitungTotalBayar(Tabel[IdxExt].Tanggal) {
			IdxExt = i
		}
	}

	fmt.Println("\n>>> MAHASISWA DENGAN PEMBAYARAN TERTINGGI <<<")
	TampilDataIndex(Tabel, IdxExt, Nominal)

	// ALGORITMA FIND MIN (Paling Sedikit Lunas)
	for i = 1; i < NMahasiswa; i++ {
		if HitungTotalBayar(Tabel[i].Tanggal) < HitungTotalBayar(Tabel[IdxExt].Tanggal) {
			IdxExt = i
		}
	}

	fmt.Println("\n>>> MAHASISWA DENGAN PEMBAYARAN TERENDAH <<<")
	TampilDataIndex(Tabel, IdxExt, Nominal)

}

func TampilDataIndex(Tabel IuranBulanan, Idx, Nominal int) {

	/* 	Prosedur untuk menampilkan data pembayaran mahasiswa yang
	 * 	berada dalam index tertentu
	 *
	 *	IS. Terdefinisi Tabel bertipe data IuranBulanan dan Idx (index mahasiswa)
	 *
	 *	FS. Menampilkan Data Pembayaran mahasiswa pada index Idx
	 */

	var TotalBayar int

	TotalBayar = HitungTotalBayar(Tabel[Idx].Tanggal)

	fmt.Printf("\n-------------------------------------------\n")
	fmt.Printf("Nama\t\t: %s\n", Tabel[Idx].Nama)
	fmt.Printf("ID\t\t: %d\n", Tabel[Idx].ID)
	fmt.Printf("Total Bayar\t: %d, (%d Hari Lunas)\n", TotalBayar*Nominal, TotalBayar)
	fmt.Printf("-------------------------------------------\n")
}

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
	 *	    Jika ada total yang sama, ID terkecil (Ascending) atau terbesar
	 *		(Descending) didahulukan. Hasil akhir tercetak di layar.
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

		SortingPembayaran(&Tabel, NMahasiswa, pilihan)
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

func SortingPembayaran(Tabel *IuranBulanan, NMahasiswa int, choice int) {

	/* 	Prosedur untuk mengurutkan data mahasiswa berdasarkan total pembayaran
	 *	sesuai pilihan yang dimasukkan pengguna.
	 *	Jika pilihan = 1 (Ascending) maka data diurutkan dengan menggunakan
	 *  algoritma SELECTION SORT, Jika pilihan = 2 (Descending) maka data
	 *  diurutkan dengan menggunakan algoritma INSERTION SORT
	 *
	 *	IS. Terdefinisi Tabel bertipe IuranBulanan, NMahasiswa (jumlah mahasiswa),
	 *		dan choice (pilihan Ascending/Descending
	 *
	 *	FS. Tabel terurut berdasarkan total bayar, secara ascending (pilihan = 1)
	 *		atau descending (pilihan = 2)
	 */

	var i, j, Ext int
	var TotalJ, TotalExt, TotalTemp int
	var temp DataIuran

	if choice == 1 {

		// Pengurutan secara ascending dengan algoritma selection sort
		// Jika ada total yg sama maka ID yg lebih besar didahulukan

		for i = 0; i < NMahasiswa-1; i++ {

			Ext = i
			TotalExt = HitungTotalBayar(Tabel[Ext].Tanggal)

			for j = i + 1; j < NMahasiswa; j++ {

				TotalJ = HitungTotalBayar(Tabel[j].Tanggal)

				if TotalJ < TotalExt {

					Ext = j
					TotalExt = HitungTotalBayar(Tabel[Ext].Tanggal)

				} else if TotalExt == TotalJ && Tabel[j].ID < Tabel[Ext].ID {
					Ext = j
					TotalExt = HitungTotalBayar(Tabel[Ext].Tanggal)

				}

			}

			// Swapping
			temp = Tabel[Ext]
			Tabel[Ext] = Tabel[i]
			Tabel[i] = temp
		}

	} else {

		// Pengurutan secara descending dengan algoritma insertion sort
		// Jika ada total yg sama maka ID yg lebih besar didahulukan

		for i = 1; i < NMahasiswa; i++ {

			j = i
			temp = Tabel[j]
			TotalTemp = HitungTotalBayar(temp.Tanggal)
			TotalJ = HitungTotalBayar(Tabel[j-1].Tanggal)

			for j > 0 && (TotalTemp > TotalJ || (TotalTemp == TotalJ && Tabel[j-1].ID > temp.ID)) {

				Tabel[j] = Tabel[j-1]
				j--

				if j > 0 {
					TotalJ = HitungTotalBayar(Tabel[j-1].Tanggal)
				}

			}

			Tabel[j] = temp
		}

	}
}

func LaporanKeuangan(Tabel IuranBulanan, NMahasiswa int, Nominal int) {

	/*	Prosedur untuk menampilkan laporan keuangan kas kelas, berisi:
	 *	1. Total per Tanggal: jumlah mahasiswa yang LUNAS pada setiap
	 *		tanggal (1-31) beserta total uang masuk pada tanggal tersebut.
	 *	2. Total Keseluruhan: total uang masuk dari seluruh mahasiswa pada
	 *		seluruh tanggal.
	 *
	 *	IS.	Terdefinisi Tabel bertipe IuranBulanan, NMahasiswa, dan Nominal.
	 *
	 *	FS. Tercetak di layar tabel Total per Tanggal dan Total Keseluruhan.
	 */

	var tgl int
	var JumlahLunas int
	var TotalKeseluruhan int

	if NMahasiswa == 0 {
		fmt.Println("********* Database masih kosong *********")
	} else if Nominal == 0 {
		fmt.Println("****** Nominal Kas Belum Di-Setting *****")
	} else {
		fmt.Print("\n=========== Total Per Tanggal ===========\n")
		fmt.Printf("| %-3v | %-13v | %-15v |\n", "Tgl", "Jml Lunas", "Total Masuk")
		fmt.Print("-----------------------------------------\n")

		TotalKeseluruhan = 0

		for tgl = 1; tgl <= HARI; tgl++ {
			JumlahLunas = JumlahLunasPerTanggal(Tabel, NMahasiswa, tgl)
			fmt.Printf("| %-3d | %-13d | Rp %-12d |\n", tgl, JumlahLunas, JumlahLunas*Nominal)
			TotalKeseluruhan += JumlahLunas * Nominal
		}

		fmt.Print("-----------------------------------------\n")
		fmt.Print("================ Summary ================\n")
		fmt.Printf("Total Uang Kas Terkumpul:\n>> Rp %d\n", TotalKeseluruhan)
		fmt.Print("-----------------------------------------\n")
	}
}

func JumlahLunasPerTanggal(Tabel IuranBulanan, NMahasiswa int, Tanggal int) int {

	/*	Fungsi untuk menghitung berapa banyak mahasiswa yang berstatus
	 *	LUNAS pada satu tanggal tertentu.
	 *
	 *	IS. Terdefinisi Tabel bertipe IuranBulanan, NMahasiswa, dan Tanggal
	 *		(rentang 1-31).
	 *	FS. Mengembalikan jumlah mahasiswa LUNAS pada Tanggal tersebut.
	 */

	var i, Jumlah int

	Jumlah = 0
	for i = 0; i < NMahasiswa; i++ {
		if Tabel[i].Tanggal[Tanggal-1] == "LUNAS" {
			Jumlah++
		}
	}

	return Jumlah
}
