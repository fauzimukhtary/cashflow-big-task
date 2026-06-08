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
		fmt.Println("| 1  | Tambah Data")
		fmt.Println("| 2  | Tampilkan Tabel Iuran Kas")
		fmt.Println("| 3  | Urutkan & Total Pembayaran")
		fmt.Println("| 4  | Update Status Pembayaran")
		fmt.Println("| 5  | Mahasiswa Belum Bayar")
		fmt.Println("| 6  | Hapus Mahasiswa")
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
				fmt.Println("========== Urutkan & Total ===========")
				MenuLaporan(&IuranMaster, JumlahMahasiswa)
			case 4:
				fmt.Println("============== Update  Kas ==============")
				UpdateStatusBayar(&IuranMaster, JumlahMahasiswa)
			case 5:
				fmt.Println("============= Belum Bayar ===============")
				BelumBayar(IuranMaster, JumlahMahasiswa)
			case 6:
				fmt.Println("============== Hapus  Data ==============")
				HapusData(&IuranMaster, &JumlahMahasiswa, &JumlahTerhapus)
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
				fmt.Println("           Terima kasih telah menggunakan CashFlow - Sampai Jumpa!           ")
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

/*
* wlm
 */

func BelumBayar(Tabel IuranBulanan, NMahasiswa int) {
	/* 	Prosedur untuk menampilkan daftar nama mahasiswa yang belum
	 *	membayar iuran pada tanggal tertentu.
	 *
	 *	IS. Terdefinisi Tabel bertipe data IuranBulanan dan NMahasiswa
	 *
	 *	FS. Menampilkan daftar nama mahasiswa yang memiliki status "-----"
	 *  pada tanggal yang diinputkan pengguna.
	 */
	var TanggalCek, i int
	var Ketemu bool

	if NMahasiswa == 0 {
		fmt.Println("********* Database masih kosong *********")
	} else {
		fmt.Print("Masukkan tanggal yang ingin dicek (1-31): ")
		fmt.Scan(&TanggalCek)
		if TanggalCek < 1 || TanggalCek > HARI {
			fmt.Println("Tanggal tidak valid (Gunakan rentang 1-31).")
		} else {
			fmt.Printf("\nMahasiswa yang BELUM BAYAR pada tanggal %d:\n", TanggalCek)
			fmt.Println("-------------------------------------------")
			Ketemu = false
			for i = 0; i < NMahasiswa; i++ {
				// Mengecek jika status masih "-----"
				if Tabel.Tanggal[i][TanggalCek-1] == "-----" {
					fmt.Printf("- [%d] %s\n", Tabel.ID[i], Tabel.Nama[i])
					Ketemu = true
				}
			}

			if !Ketemu {
				fmt.Println("Hebat! Semua mahasiswa sudah bayar pada tanggal ini.")
			}
			fmt.Println("-------------------------------------------")
		}
	}
}

func HitungTotalBayar(Data TanggalBayar) int {
	/*	Fungsi untuk menghitung berapa kali status "LUNAS" muncul
	 *	dalam array TanggalBayar milik seorang mahasiswa.
	 */
	var count, i int
	for i = 0; i < HARI; i++ {
		if Data[i] == "LUNAS" {
			count++
		}
	}
	return count
}

func MenuLaporan(Tabel *IuranBulanan, NMahasiswa int) {
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
	var pilihan, i, j, idx_ekstrim int
	var tID int
	var tNama string
	var tTanggal TanggalBayar

	if NMahasiswa == 0 {
		fmt.Println("********* Database masih kosong *********")
		return
	}

	fmt.Println("Pilih urutan tampilan:")
	fmt.Println("1. Total Bayar Terbanyak ke Terkecil (Descending)")
	fmt.Println("2. Total Bayar Terkecil ke Terbanyak (Ascending)")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)

	if pilihan != 1 && pilihan != 2 {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	// ALGORITMA SELECTION SORT SESUAI MODUL
	i = 1
	for i <= NMahasiswa-1 {
		idx_ekstrim = i - 1
		j = i
		for j < NMahasiswa {
			totalJ := HitungTotalBayar(Tabel.Tanggal[j])
			totalEkstrim := HitungTotalBayar(Tabel.Tanggal[idx_ekstrim])

			if pilihan == 1 {
				// Descending: Cari total yang lebih besar
				if totalEkstrim < totalJ {
					idx_ekstrim = j
				} else if totalEkstrim == totalJ {
					// Jika TOTAL SAMA, pilih yang ID-nya LEBIH KECIL
					if Tabel.ID[idx_ekstrim] > Tabel.ID[j] {
						idx_ekstrim = j
					}
				}
			} else {
				// Ascending: Cari total yang lebih kecil
				if totalEkstrim > totalJ {
					idx_ekstrim = j
				} else if totalEkstrim == totalJ {
					// Jika TOTAL SAMA, pilih yang ID-nya LEBIH KECIL
					if Tabel.ID[idx_ekstrim] > Tabel.ID[j] {
						idx_ekstrim = j
					}
				}
			}
			j = j + 1
		}

		// Proses Pertukaran (Swap) semua field agar data tetap sinkron
		tID = Tabel.ID[idx_ekstrim]
		Tabel.ID[idx_ekstrim] = Tabel.ID[i-1]
		Tabel.ID[i-1] = tID

		tNama = Tabel.Nama[idx_ekstrim]
		Tabel.Nama[idx_ekstrim] = Tabel.Nama[i-1]
		Tabel.Nama[i-1] = tNama

		tTanggal = Tabel.Tanggal[idx_ekstrim]
		Tabel.Tanggal[idx_ekstrim] = Tabel.Tanggal[i-1]
		Tabel.Tanggal[i-1] = tTanggal

		i = i + 1
	}

	// TAMPILKAN HASIL
	if pilihan == 1 {
		fmt.Println("\n--- Laporan: Terbanyak ke Terkecil (Tie-breaker: ID Kecil) ---")
	} else {
		fmt.Println("\n--- Laporan: Terkecil ke Terbanyak (Tie-breaker: ID Kecil) ---")
	}

	fmt.Printf("| %-3v | %-30v | %-12v |\n", "ID", "Nama Mahasiswa", "Total Bayar")
	fmt.Println("--------------------------------------------")
	for i = 0; i < NMahasiswa; i++ {
		fmt.Printf("| %-3d | %-30s | %-12d |\n",
			Tabel.ID[i],
			Tabel.Nama[i],
			HitungTotalBayar(Tabel.Tanggal[i]),
		)
	}
}
