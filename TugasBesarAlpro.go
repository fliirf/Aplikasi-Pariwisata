package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TempatWisata struct {
	ID        int
	Nama      string
	Lokasi    string
	Biaya     int
	Deskripsi string
	Jarak     float64
	Rating    float64
}

const NMAX = 100

type tempatWisataList [NMAX]TempatWisata

var List tempatWisataList
var ID int = 1
var n int = 0

func main() {
	for {
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("                           VisitVista                           ")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("1. Tambah Tempat Wisata")
		fmt.Println("2. Edit Tempat Wisata")
		fmt.Println("3. Hapus Tempat Wisata")
		fmt.Println("4. Lihat Tempat Wisata")
		fmt.Println("5. Cari Tempat Wisata")
		fmt.Println("6. Keluar")
		fmt.Println("----------------------------------------------------------------")
		fmt.Print("Pilih Opsi: ")

		var option int
		fmt.Scan(&option)

		if option == 1 {
			tambahTempatWisata()
		} else if option == 2 {
			editTempatWisata()
		} else if option == 3 {
			hapusTempatWisata()
		} else if option == 4 {
			menuLihat()
		} else if option == 5 {
			cariTempatWisata()
		} else if option == 6 {
			fmt.Println("Terima kasih telah menggunakan VisitVista.")
			break
		} else {
			fmt.Println("Opsi tidak valid. Silahkan coba lagi.")
		}
	}
}

func tambahTempatWisata() {
	reader := bufio.NewReader(os.Stdin)
	var wisata TempatWisata

	wisata.ID = ID

	fmt.Print("Masukkan nama tempat wisata: ")
	nama, _ := reader.ReadString('\n')
	wisata.Nama = strings.TrimSpace(nama)

	fmt.Print("Masukkan lokasi tempat wisata: ")
	lokasi, _ := reader.ReadString('\n')
	wisata.Lokasi = strings.TrimSpace(lokasi)

	fmt.Print("Masukkan biaya tempat wisata: ")
	biaya, _ := reader.ReadString('\n')
	wisata.Biaya, _ = strconv.Atoi(strings.TrimSpace(biaya))

	fmt.Print("Masukkan jarak tempat wisata (Dalam KM): ")
	jarak, _ := reader.ReadString('\n')
	wisata.Jarak, _ = strconv.ParseFloat(strings.TrimSpace(jarak), 64)

	fmt.Print("Masukkan rating tempat wisata (dalam skala 1.0 - 5.0): ")
	rating, _ := reader.ReadString('\n')
	wisata.Rating, _ = strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if wisata.Rating > 5.0 || wisata.Rating < 1.0 {
		fmt.Println("Rating tidak valid. Silahkan coba lagi.")
		return
	}

	fmt.Print("Masukkan deskripsi tempat wisata: ")
	deskripsi, _ := reader.ReadString('\n')
	wisata.Deskripsi = strings.TrimSpace(deskripsi)

	List[n] = wisata
	n++
	ID++

	fmt.Println("Tempat wisata berhasil ditambahkan.")
}

func editTempatWisata() {
	var id, edit int

	fmt.Print("Masukkan ID tempat wisata yang ingin diedit: ")
	fmt.Scan(&id)

	idx := SearchID(List, n, id)
	if idx == -1 {
		fmt.Println("Tempat wisata dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("1. Nama Tempat wisata")
	fmt.Println("2. Lokasi Tempat wisata")
	fmt.Println("3. Biaya Tempat wisata")
	fmt.Println("4. Deskripsi Tempat wisata")
	fmt.Println("5. Jarak Tempat wisata")
	fmt.Println("6. Rating Tempat wisata")
	fmt.Print("Masukkan nomor pilihan yang ingin diedit: ")
	fmt.Scan(&edit)
	if edit == 1 {
		editNama(idx)
	} else if edit == 2 {
		editLokasi(idx)
	} else if edit == 3 {
		editBiaya(idx)
	} else if edit == 4 {
		editDeskripsi(idx)
	} else if edit == 5 {
		editJarak(idx)
	} else if edit == 6 {
		editRating(idx)
	} else {
		fmt.Println("Input tidak valid. Silahkan coba lagi.")
	}
}

func SearchID(List tempatWisataList, n, idxID int) int {
	for i := 0; i < n; i++ {
		if List[i].ID == idxID {
			return i
		}
	}
	return -1
}

func editNama(idx int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama baru: ")
	gantinama, _ := reader.ReadString('\n')
	gantinama = strings.TrimSpace(gantinama)
	List[idx].Nama = gantinama
	fmt.Println("Nama tempat wisata telah diedit")
}

func editLokasi(idx int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan lokasi baru: ")
	gantilokasi, _ := reader.ReadString('\n')
	gantilokasi = strings.TrimSpace(gantilokasi)
	List[idx].Lokasi = gantilokasi
	fmt.Println("Lokasi tempat wisata telah diedit")
}

func editBiaya(idx int) {
	var gantibiaya int
	fmt.Print("Masukkan biaya baru: ")
	fmt.Scan(&gantibiaya)
	List[idx].Biaya = gantibiaya
	fmt.Println("Biaya tempat wisata telah diedit")
}

func editDeskripsi(idx int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan deskripsi baru: ")
	gantideskripsi, _ := reader.ReadString('\n')
	gantideskripsi = strings.TrimSpace(gantideskripsi)
	List[idx].Deskripsi = gantideskripsi
	fmt.Println("Deskripsi tempat wisata telah diedit")
}

func editJarak(idx int) {
	var gantijarak float64
	fmt.Print("Masukkan jarak baru: ")
	fmt.Scan(&gantijarak)
	List[idx].Jarak = gantijarak
	fmt.Println("Jarak tempat wisata telah diedit")
}

func editRating(idx int) {
	var gantirating float64
	fmt.Print("Masukkan rating baru: ")
	fmt.Scan(&gantirating)
	List[idx].Rating = gantirating
	fmt.Println("Biaya rating wisata telah diedit")
}

func hapusTempatWisata() {
	var id int
	fmt.Print("Masukkan ID tempat wisata yang ingin dihapus : ")
	fmt.Scan(&id)

	hapus := SearchID(List, n, id)
	if hapus == -1 {
		fmt.Println("Tempat wisata dengan ID tersebut tidak ditemukan.")
		return
	}

	for i := hapus; i < n-1; i++ {
		List[i] = List[i+1]
		List[i].ID = i + 1
	}
	n--
	fmt.Println("Tempat wisata berhasil dihapus.")

	for i := 0; i < n; i++ {
		List[i].ID = i + 1
	}
}

func lihatTempatWisata() {
	fmt.Println("Daftar Tempat Wisata:")
	fmt.Printf("%-4s %-20s %-20s %-8s %-15s %-8s\n", "ID", "Nama", "Lokasi", "Biaya", "Jarak (KM)", "Rating")
	for i := 0; i < n; i++ {
		fmt.Printf("%-4d %-20s %-20s %-8d %-15.2f %-8.2f\n", List[i].ID, List[i].Nama, List[i].Lokasi, List[i].Biaya, List[i].Jarak, List[i].Rating)
	}
}

func lihatDeskripsi() {
	var id int
	fmt.Print("Masukkan ID tempat wisata yang ingin dilihat deskripsinya: ")
	fmt.Scan(&id)
	idx := SearchID(List, n, id)
	if idx == -1 {
		fmt.Println("Tempat wisata dengan ID tersebut tidak ditemukan.")
	} else {
		fmt.Println("Deskripsi Tempat Wisata:", List[idx].Deskripsi)
	}
}

func cariTempatWisata() {
	var cari string
	fmt.Println("1. Nama Tempat Wisata")
	fmt.Println("2. Lokasi Tempat Wisata")
	fmt.Println("3. Biaya Tempat Wisata")
	fmt.Println("4. Jarak Tempat Wisata")
	fmt.Println("5. Rating Tempat Wisata")
	fmt.Print("Masukkan nomor pilihan yang ingin dicari: ")
	fmt.Scan(&cari)
	if cari == "1" {
		cariNama()
	} else if cari == "2" {
		cariLokasi()
	} else if cari == "3" {
		cariBiaya()
	} else if cari == "4" {
		cariJarak()
	} else if cari == "5" {
		cariRating()
	} else {
		fmt.Println("Input tidak valid. Silahkan coba lagi.")
	}
}

func cariNama() {
	var cari string
	fmt.Print("Masukkan nama tempat wisata yang ingin dicari: ")
	fmt.Scan(&cari)

	for i := 0; i < n; i++ {
		if List[i].Nama == cari {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
		} else if List[i].Nama != cari {
			fmt.Println("Tempat wisata dengan nama tersebut tidak ditemukan.")
		}
	}
}

func cariLokasi() {
	var cari string
	fmt.Print("Masukkan lokasi tempat wisata yang ingin dicari: ")
	fmt.Scan(&cari)

	for i := 0; i < n; i++ {
		if List[i].Lokasi == cari {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
		} else if List[i].Lokasi != cari {
			fmt.Println("Tempat wisata dengan lokasi tersebut tidak ditemukan.")
		}
	}
}

func cariBiaya() {
	var no int
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Rentang perkiraan biaya tempat wisata")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("1. Rp0 - Rp100000")
	fmt.Println("2. Rp100000 - Rp300000")
	fmt.Println("3. Rp300000 - Rp500000")
	fmt.Println("4. Rp500000 - Rp1000000")
	fmt.Println("5. Rp1000000++")
	fmt.Print("Masukkan nomor pilihan yang ingin dicari: ")
	fmt.Scan(&no)

	found := false

	for i := 0; i < n; i++ {
		if no == 1 && List[i].Biaya >= 0 && List[i].Biaya <= 100000 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		} else if no == 2 && List[i].Biaya > 100000 && List[i].Biaya <= 300000 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		} else if no == 3 && List[i].Biaya > 300000 && List[i].Biaya <= 500000 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		} else if no == 4 && List[i].Biaya > 500000 && List[i].Biaya <= 1000000 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		} else if no == 5 && List[i].Biaya > 1000000 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		}
	}
	if !found {
		fmt.Println("Tempat wisata dengan biaya tersebut tidak ditemukan.")
	} else if no < 1 || no > 5 {
		fmt.Println("Input tidak valid. Silahkan coba lagi.")
	}
}

func cariJarak() {
	var no int
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Rentang perkiraan jarak tempat wisata")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("1. 0 - 5 Km")
	fmt.Println("2. 5 - 10 Km")
	fmt.Println("3. 10 - 20 Km")
	fmt.Println("4. 20 - 50 Km")
	fmt.Println("5. 50 Km++")
	fmt.Print("Masukkan nomor pilihan yang ingin dicari: ")
	fmt.Scan(&no)

	found := false
	for i := 0; i < n; i++ {
		if no == 1 && List[i].Jarak >= 0 && List[i].Jarak <= 5 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		} else if no == 2 && List[i].Jarak >= 5 && List[i].Jarak <= 10 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		} else if no == 3 && List[i].Jarak >= 10 && List[i].Jarak <= 20 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		} else if no == 4 && List[i].Jarak >= 20 && List[i].Jarak <= 50 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		} else if no == 5 && List[i].Jarak > 50 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		}
	}

	if !found {
		fmt.Println("Tempat wisata dengan jarak tersebut tidak ditemukan.")
	} else if no < 1 || no > 5 {
		fmt.Println("Input tidak valid. Silahkan coba lagi.")
	}
}

func cariRating() {
	var no int
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Rentang rating tempat wisata")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("1. 1 - 2")
	fmt.Println("2. 2 - 3")
	fmt.Println("3. 3 - 4")
	fmt.Println("4. 4 - 5")
	fmt.Print("Masukkan nomor pilihan yang ingin dicari: ")
	fmt.Scan(&no)

	found := false
	for i := 0; i < n; i++ {
		if no == 1 && List[i].Rating >= 1 && List[i].Rating <= 2 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		} else if no == 2 && List[i].Rating >= 2 && List[i].Rating <= 3 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		} else if no == 3 && List[i].Rating >= 3 && List[i].Rating <= 4 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		} else if no == 4 && List[i].Rating >= 4 && List[i].Rating <= 5 {
			fmt.Println("No:", List[i].ID, ",", "Nama:", List[i].Nama, ",", "Lokasi:", List[i].Lokasi, ",", "Biaya:", List[i].Biaya, ",", "Jarak:", List[i].Jarak, "Rating: ", List[i].Rating)
			found = true
		}
	}

	if !found {
		fmt.Println("Tempat wisata dengan rating tersebut tidak ditemukan.")
	} else if no < 1 || no > 5 {
		fmt.Println("Input tidak valid. Silahkan coba lagi.")
	}
}

func urutJarakUp() {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if List[i].Jarak > List[j].Jarak {
				temp := List[i]
				List[i] = List[j]
				List[j] = temp
			}
		}
	}
}

func urutBiayaUp() {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if List[i].Biaya > List[j].Biaya {
				temp := List[i]
				List[i] = List[j]
				List[j] = temp
			}
		}
	}
}

func urutRatingUp() {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if List[i].Rating > List[j].Rating {
				temp := List[i]
				List[i] = List[j]
				List[j] = temp
			}
		}
	}
}

func urutJarakDown() {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if List[i].Jarak < List[j].Jarak {
				temp := List[i]
				List[i] = List[j]
				List[j] = temp
			}
		}
	}
}

func urutBiayaDown() {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if List[i].Biaya < List[j].Biaya {
				temp := List[i]
				List[i] = List[j]
				List[j] = temp
			}
		}
	}
}

func urutRatingDown() {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if List[i].Rating < List[j].Rating {
				temp := List[i]
				List[i] = List[j]
				List[j] = temp
			}
		}
	}
}

func menuLihat() {
	var opsi int
	fmt.Println("1. Lihat Daftar Tempat Wisata")
	fmt.Println("2. Lihat Deskripsi Tempat Wisata")
	fmt.Println("3. Lihat Urutan berdasarkan Jarak/Rating/Biaya")
	fmt.Print("Masukan nomor pilihan yang ingin dilihat: ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		lihatTempatWisata()
	} else if opsi == 2 {
		lihatDeskripsi()
	} else if opsi == 3 {
		urutTempatWisata()
	} else {
		fmt.Println("Input tidak valid. Silahkan coba lagi.")
	}
}

func urutTempatWisata() {
	var opsi, opsi2 int
	fmt.Print("Urutan Berdasarkan (1. Jarak, 2.Biaya, 3.Rating): ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		fmt.Print("Masukan nomor pilihan (1.Terdekat, 2.Terjauh): ")
		fmt.Scan(&opsi2)
		if opsi2 == 1 {
			urutJarakUp()
			lihatTempatWisata()
		} else if opsi2 == 2 {
			urutJarakDown()
			lihatTempatWisata()
		} else {
			fmt.Println("Input tidak valid. Silahkan coba lagi.")
		}
	}

	if opsi == 2 {
		fmt.Print("Masukan nomor pilihan (1.Termurah, 2.Termahal):")
		fmt.Scan(&opsi2)
		if opsi2 == 1 {
			urutBiayaUp()
			lihatTempatWisata()
		} else if opsi2 == 2 {
			urutBiayaDown()
			lihatTempatWisata()
		} else {
			fmt.Println("Input tidak valid. Silahkan coba lagi.")
		}
	}

	if opsi == 3 {
		fmt.Print("Masukan nomor pilihan (1.Terendah, 2.Tertinggi):")
		fmt.Scan(&opsi2)
		if opsi2 == 1 {
			urutRatingUp()
			lihatTempatWisata()
		} else if opsi2 == 2 {
			urutRatingDown()
			lihatTempatWisata()
		} else {
			fmt.Println("Input tidak valid. Silahkan coba lagi.")
		}
	}
}
