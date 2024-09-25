package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const NMAX int = 1000

type barang struct {
	nama, kategori                      string
	hargaAwal, hargaJual, stok, terjual int
}

type arrBarang [NMAX]barang

type transaksi struct {
	belibarang                  string
	quantity, total, totalHarga int
}
type arrTran [NMAX]transaksi

func hapusLayar() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func menu(A arrBarang, b arrTran, n int, q int) {
	var pil, pil2, y, cari string
	var pilih int
	for pilih != 8 {
		fmt.Println("*------------------------------------------------*")
		fmt.Println("|             Zonder Computer Shop               |")
		fmt.Println("*------------------------------------------------*")
		fmt.Println("| 1. Mengolah data barang                        |")
		fmt.Println("| 2. Mengolah data transaksi                     |")
		fmt.Println("| 3. Tampilkan data barang                       |")
		fmt.Println("| 4. Cari data barang                            |")
		fmt.Println("| 5. Tampilkan data transaksi                    |")
		fmt.Println("| 6. Tampilkan barang yang paling banyak terjual |")
		fmt.Println("| 7. Tampilkan Modal dan Pendapatan              |")
		fmt.Println("| 8. Keluar                                      |")
		fmt.Println("*------------------------------------------------*")
		fmt.Print("Tentukan pilihan Anda (1/2/3/4/5/6/7): ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			hapusLayar()
			fmt.Print("Apa yang ingin Anda lakukan (tambah/edit/hapus): ")
			fmt.Scan(&pil)
			if pil == "tambah" {
				hapusLayar()
				masukkanBarang(&A, &n)
				fmt.Print("Apakah Anda ingin memasukkan barang lagi? (ya/tidak): ")
				fmt.Scan(&y)
				for y == "ya" {
					masukkanBarang(&A, &n)
					fmt.Print("Apakah Anda ingin memasukkan barang lagi? (ya/tidak): ")
					fmt.Scan(&y)
				}
			} else if pil == "edit" {
				hapusLayar()
				editBarang(&A, n)
				fmt.Print("Apakah Anda ingin melakukan edit lagi? (ya/tidak): ")
				fmt.Scan(&y)
				for y == "ya" {
					hapusLayar()
					editBarang(&A, n)
					fmt.Print("Apakah Anda ingin melakukan edit lagi? (ya/tidak): ")
					fmt.Scan(&y)
				}
			} else if pil == "hapus" {
				hapusLayar()
				hapusBarang(&A, &n)
				fmt.Println("Data berhasil diubah menjadi: ")
				tampilkanBarang(A, n)
				fmt.Print("Apakah Anda ingin menghapus barang lagi? (ya/tidak): ")
				fmt.Scan(&y)
				for y == "ya" {
					hapusLayar()
					hapusBarang(&A, &n)
					fmt.Println("Data berhasil diubah menjadi: ")
					tampilkanBarang(A, n)
					fmt.Print("Apakah Anda ingin menghapus barang lagi? (ya/tidak): ")
					fmt.Scan(&y)
				}
			}
			hapusLayar()
			menu(A, b, n, q)
		case 2:
			hapusLayar()
			fmt.Print("Apa yang ingin Anda lakukan (tambah/edit/hapus): ")
			fmt.Scan(&pil)
			if pil == "tambah" {
				hapusLayar()
				fmt.Println(strings.Repeat("=", 150))
				fmt.Println("                  Daftar Barang                  ")
				tampilkanBarang(A, n)
				tambahTransaksi(&A, &b, n, &q)
				fmt.Print("Apakah Anda menambahkan transaksi lagi? (ya/tidak): ")
				fmt.Scan(&y)
				for y == "ya" {
					hapusLayar()
					tambahTransaksi(&A, &b, n, &q)
					fmt.Print("Apakah Anda menambahkan transaksi lagi? (ya/tidak): ")
					fmt.Scan(&y)
				}
			} else if pil == "edit" {
				hapusLayar()
				editTrans(&A, &b, n, &q)
				fmt.Print("Apakah Anda ingin edit transaksi lagi? (ya/tidak): ")
				fmt.Scan(&y)
				for y == "ya" {
					hapusLayar()
					editTrans(&A, &b, n, &q)
					fmt.Print("Apakah Anda ingin edit transaksi lagi? (ya/tidak): ")
					fmt.Scan(&y)
				}
			} else if pil == "hapus" {
				hapusLayar()
				hapusTrans(&A, &b, n, &q)
				fmt.Print("Apakah Anda ingin hapus transaksi lagi? (ya/tidak): ")
				fmt.Scan(&y)
				for y == "ya" {
					hapusLayar()
					hapusTrans(&A, &b, n, &q)
					fmt.Print("Apakah Anda ingin hapus transaksi lagi? (ya/tidak): ")
					fmt.Scan(&y)
				}
			}
			hapusLayar()
			menu(A, b, n, q)
		case 3:
			hapusLayar()
			tampilkanBarang(A, n)
			fmt.Print("Apakah Anda ingin mengurutkan data? (ya/tidak): ")
			fmt.Scan(&pil)
			for pil == "ya" {
				hapusLayar()
				fmt.Print("Urutkan berdasarkan (nama/hargaAwal/kategori/hargaJual/stok): ")
				fmt.Scan(&pil2)
				sortBarang(A, n, pil2)
				fmt.Print("Apakah Anda ingin mengurutkan data lagi? (ya/tidak): ")
				fmt.Scan(&pil)
			}
			hapusLayar()
			menu(A, b, n, q)
		case 4:
			hapusLayar()
			fmt.Print("Anda ingin mencari data barang berdasarkan apa? (nama/hargaAwal/kategori/hargaJual/stok): ")
			fmt.Scan(&pil)
			if pil == "nama" {
				fmt.Print("Berikan nama barang yang ingin Anda cari: ")
				fmt.Scan(&cari)
				cariBarang(A, n, cari)
			} else if pil == "hargaAwal" || pil == "hargaJual" || pil == "stok" {
				cariHarga(A, n, pil)
			} else if pil == "kategori" {
				fmt.Print("Berikan kategori barang yang ingin Anda cari: ")
				fmt.Scan(&cari)
				cariKategori(A, n, cari)
			}
			fmt.Print("Apakah Anda ingin mencari barang lagi? (ya/tidak): ")
			fmt.Scan(&y)
			for y == "ya" {
				hapusLayar()
				fmt.Print("Anda ingin mencari data barang berdasarkan apa (nama/hargaAwal/kategori/hargaJual/stok): ")
				fmt.Scan(&pil)
				if pil == "nama" {
					fmt.Print("Berikan nama barang yang ingin Anda cari: ")
					fmt.Scan(&cari)
					cariBarang(A, n, cari)
				} else if pil == "hargaAwal" || pil == "hargaJual" || pil == "stok" {
					cariHarga(A, n, pil)
				} else if pil == "kategori" {
					fmt.Print("Berikan kategori barang yang ingin Anda cari: ")
					fmt.Scan(&cari)
					cariKategori(A, n, cari)
				}
				fmt.Print("Apakah Anda ingin mencari barang lagi? (ya/tidak): ")
				fmt.Scan(&y)
			}
			hapusLayar()
			menu(A, b, n, q)
		case 5:
			hapusLayar()
			tampilkanData(b, q)
			fmt.Print("Apakah Anda ingin mengurutkan data? (ya/tidak): ")
			fmt.Scan(&pil)
			for pil == "ya" {
				hapusLayar()
				fmt.Print("Urutkan berdasarkan (nama/quantity/harga): ")
				fmt.Scan(&pil2)
				sortTransaksi(b, q, pil2)
				fmt.Print("Apakah Anda ingin mengurutkan data lagi? (ya/tidak): ")
				fmt.Scan(&pil)
			}
			hapusLayar()
			menu(A, b, n, q)
		case 6:
			hapusLayar()
			terjualTerbanyak(A, n)
			fmt.Print("Apakah Anda ingin menampilkan datanya lagi? (ya/tidak): ")
			fmt.Scan(&y)
			for y == "ya" {
				hapusLayar()
				terjualTerbanyak(A, n)
				fmt.Print("Apakah Anda ingin menampilkan datanya lagi? (ya/tidak): ")
				fmt.Scan(&y)
			}
			hapusLayar()
			menu(A, b, n, q)
		case 7:
			hapusLayar()
			hitung(A, b, n, q)
			fmt.Print("Apakah Anda ingin menampilkan datanya lagi? (ya/tidak): ")
			fmt.Scan(&y)
			for y == "ya" {
				hapusLayar()
				hitung(A, b, n, q)
				fmt.Print("Apakah Anda ingin menampilkan datanya lagi? (ya/tidak): ")
				fmt.Scan(&y)
			}
			hapusLayar()
			menu(A, b, n, q)
		case 8:
			hapusLayar()
			fmt.Println("----------------")
			fmt.Println("| Terima Kasih |")
			fmt.Println("----------------")
			os.Exit(0)
		}
	}
}

// olah data barang //

func masukkanBarang(A *arrBarang, n *int) {
	var i, j, x, modal, harga, stok, idx int
	var nama, kategori string
	fmt.Print("Berapa barang yang ingin Anda masukkan: ")
	fmt.Scan(&x)
	fmt.Println("(nama, hargaAwal, kategori, hargaJual, stok)")
	for i < x {
		fmt.Scan(&nama, &modal, &kategori, &harga, &stok)
		idx = SeqBarang(*A, *n, nama)
		if idx != -1 {
			A[idx].stok += stok
		} else {
			j = *n
			A[j].nama = nama
			A[j].hargaAwal = modal
			A[j].kategori = kategori
			A[j].hargaJual = harga
			A[j].stok = stok
			*n++
		}
		i++
	}
}

func editBarang(A *arrBarang, n int) {
	var index, index2 int
	var x, y string
	tampilkanBarang(*A, n)
	fmt.Print("Masukkan nama barang yang ingin Anda ubah: ")
	fmt.Scan(&x)
	index = SeqBarang(*A, n, x)
	if index != -1 {
		fmt.Print("Ubah data barang: ")
		fmt.Println("(nama)")
		fmt.Scan(&y)
		index2 = SeqBarang(*A, n, y)
		if index2 == -1 || A[index2].nama == A[index].nama {
			A[index].nama = y
			fmt.Println("(hargaAwal, kategori, hargaJual, stok)")
			fmt.Scan(&A[index].hargaAwal, &A[index].kategori, &A[index].hargaJual, &A[index].stok)
			fmt.Println("Data berhasil diubah menjadi: ")
			tampilkanBarang(*A, n)
		} else {
			fmt.Println("Nama barang sudah ada")
		}
	} else {
		fmt.Println("Data barang tidak ditemukan")
	}
}

func hapusBarang(A *arrBarang, n *int) {
	var index, i int
	var x string
	tampilkanBarang(*A, *n)
	fmt.Print("Masukkan nama barang yang ingin Anda hapus: ")
	fmt.Scan(&x)
	index = SeqBarang(*A, *n, x)
	if index != -1 {
		i = index
		for i < *n-1 {
			A[i] = A[i+1]
			i++
		}
		*n--
	} else {
		fmt.Println("Data barang tidak ditemukan")
	}
}

func tambahTransaksi(a *arrBarang, b *arrTran, n int, q *int) {
	var barang, tanya string
	var i, j, jumlah, idx, x int
	fmt.Print("Berapa banyak transaksi yang sudah dilakukan: ")
	fmt.Scan(&x)
	for i < x {
		fmt.Scan(&barang, &jumlah)
		idx = SeqBarang(*a, n, barang)
		for idx == -1 || a[idx].stok-jumlah < 0 {
			fmt.Println("Barangnya tidak ada atau stoknya sudah habis, apakah anda ingin ke (menu/tambah)")
			fmt.Scan(&tanya)
			if tanya == "menu" {
				menu(*a, *b, n, *q)
			} else if tanya == "tambah" {
				fmt.Scan(&barang, &jumlah)
				idx = SeqBarang(*a, n, barang)
			} else {
				fmt.Println("tidak relevan")
				fmt.Scan(&tanya)
			}
		}
		j = *q
		b[j].belibarang = barang
		b[j].quantity = jumlah
		a[idx].stok -= jumlah
		a[idx].terjual += jumlah
		b[j].totalHarga = a[idx].hargaJual * b[j].quantity
		*q++
		i++
	}
}

func editTrans(a *arrBarang, b *arrTran, n int, q *int) {
	var barang, tanya string
	var i, idx, jumlah int
	tampilkanData(*b, *q)
	fmt.Print("Masukkan data ke berapa yang ingin diubah: ")
	fmt.Scan(&i)
	i--
	for i == *q {
		fmt.Print("melebihi")
		fmt.Scan(&i)
		i--
	}

	fmt.Println("Masukan perubahannya:")
	fmt.Scan(&barang, &jumlah)
	idx = SeqBarang(*a, n, barang)
	for idx == -1 || a[idx].stok+b[i].quantity-jumlah < 0 {
		fmt.Print("Barangnya tidak ada atau stoknya sudah habis, apakah anda ingin ke (menu/edit): ")
		fmt.Scan(&tanya)
		if tanya == "menu" {
			menu(*a, *b, n, *q)
		} else if tanya == "edit" {
			fmt.Scan(&barang, &jumlah)
			idx = SeqBarang(*a, n, barang)
		} else {
			fmt.Print("Masukan tidak valid, input lagi: ")
			fmt.Scan(&tanya)
		}
	}
	if a[idx].nama == b[i].belibarang {
		a[idx].stok += b[i].quantity - jumlah
		a[idx].terjual += jumlah - b[i].quantity
	} else {
		idx = SeqBarang(*a, n, b[i].belibarang)
		a[idx].stok += b[i].quantity
		a[idx].terjual -= b[i].quantity
		hapusTransEdit(b, q, b[i].belibarang)
		idx = SeqBarang(*a, n, barang)
		a[idx].stok -= jumlah
		a[idx].terjual += jumlah
		*q++
	}
	b[i].belibarang = barang
	b[i].quantity = jumlah
	b[i].totalHarga = a[idx].hargaJual * b[i].quantity
}

func hapusTrans(a *arrBarang, b *arrTran, n int, q *int) {
	var i, idx int
	tampilkanData(*b, *q)
	fmt.Print("Masukkan data ke berapa yang ingin Anda hapus: ")
	fmt.Scan(&i)
	i--
	idx = SeqBarang(*a, n, b[i].belibarang)
	a[idx].stok += b[i].quantity
	a[idx].terjual -= b[i].quantity
	for i == *q {
		fmt.Print("melebihi")
		fmt.Scan(&i)
		i--
	}
	for i < *q-1 {
		b[i] = b[i+1]
		i++
	}
	*q--

	fmt.Println("==========================")
	fmt.Println("Data Berhasil dihapus: ")
	fmt.Println("--------------------------")
	tampilkanData(*b, *q)
	fmt.Println("==========================")
}

func hapusTransEdit(b *arrTran, n *int, nama string) {
	var index, i int
	index = SeqTransaksi(*b, *n, nama)
	i = index
	for i < *n-1 {
		b[i] = b[i+1]
		i++
	}
	*n = *n - 1
}

// hitung modal dan pendapatan //

func hitung(A arrBarang, b arrTran, n, q int) {
	i := 0
	index := -1
	var x, y, z int
	x = 0
	y = 0
	for i < q {
		index = SeqBarang(A, n, b[i].belibarang)
		x = x + ((A[index].terjual + A[index].stok) * A[index].hargaAwal)
		y = y + (b[i].quantity * A[index].hargaJual)
		i++
	}
	z = y - x
	fmt.Println(strings.Repeat("=", 100))
	fmt.Printf("%-20s%-20s%-20s\n", "Modal", "Hasil", "Pendapatan")
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("%-20d%-20d%-20d\n", x, y, z)
	fmt.Println(strings.Repeat("=", 100))
}

// Sorting //

func sortBarang(A arrBarang, n int, pilih string) {
	var pass, idx, i int
	var temp barang
	var urutan string
	fmt.Print("Anda ingin mengurutkan (ascending/descending): ")
	fmt.Scan(&urutan)
	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		if urutan == "ascending" {
			for i < n {
				if pilih == "nama" {
					if A[idx].nama > A[i].nama {
						idx = i
					}
				} else if pilih == "kategori" {
					if A[idx].kategori > A[i].kategori {
						idx = i
					} else if A[idx].kategori == A[i].kategori {
						if A[idx].nama > A[i].nama {
							idx = i
						}
					}
				} else if pilih == "hargaAwal" {
					if A[idx].hargaAwal > A[i].hargaAwal {
						idx = i
					} else if A[idx].hargaAwal == A[i].hargaAwal {
						if A[idx].nama > A[i].nama {
							idx = i
						}
					}
				} else if pilih == "hargaJual" {
					if A[idx].hargaJual > A[i].hargaJual {
						idx = i
					} else if A[idx].hargaJual == A[i].hargaJual {
						if A[idx].nama > A[i].nama {
							idx = i
						}
					}
				} else if pilih == "stok" {
					if A[idx].stok > A[i].stok {
						idx = i
					} else if A[idx].stok == A[i].stok {
						if A[idx].nama > A[i].nama {
							idx = i
						}
					}
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		} else if urutan == "descending" {
			for i < n {
				if pilih == "nama" {
					if A[idx].nama < A[i].nama {
						idx = i
					}
				} else if pilih == "kategori" {
					if A[idx].kategori < A[i].kategori {
						idx = i
					} else if A[idx].kategori == A[i].kategori {
						if A[idx].nama < A[i].nama {
							idx = i
						}
					}
				} else if pilih == "hargaAwal" {
					if A[idx].hargaAwal < A[i].hargaAwal {
						idx = i
					} else if A[idx].hargaAwal == A[i].hargaAwal {
						if A[idx].nama < A[i].nama {
							idx = i
						}
					}
				} else if pilih == "hargaJual" {
					if A[idx].hargaJual < A[i].hargaJual {
						idx = i
					} else if A[idx].hargaJual == A[i].hargaJual {
						if A[idx].nama < A[i].nama {
							idx = i
						}
					}
				} else if pilih == "stok" {
					if A[idx].stok < A[i].stok {
						idx = i
					} else if A[idx].stok == A[i].stok {
						if A[idx].nama < A[i].nama {
							idx = i
						}
					}
				}
				i++
			}
			temp = A[pass-1]
			A[pass-1] = A[idx]
			A[idx] = temp
			pass++
		} else {
			fmt.Print("Masukan tidak valid, input lagi (ascending/descending): ")
			fmt.Scan(&urutan)
		}
	}
	fmt.Println(strings.Repeat("=", 150))
	fmt.Println("           Data barang setelah sorting           ")
	tampilkanBarang(A, n)
}

func sortTransaksi(b arrTran, n int, sort string) {
	var pass, i int
	var temp transaksi
	var tanya string
	pass = 1
	fmt.Print("Anda ingin mengurutkan (ascending/descending): ")
	fmt.Scan(&tanya)
	if tanya == "descending" {
		for pass < n {
			i = pass
			temp = b[pass]
			if sort == "nama" {
				for i > 0 && temp.belibarang > b[i-1].belibarang {
					b[i] = b[i-1]
					i--
				}
			} else if sort == "quantity" {
				for i > 0 && temp.quantity > b[i-1].quantity {
					b[i] = b[i-1]
					i--
				}
			} else if sort == "harga" {
				for i > 0 && temp.totalHarga > b[i-1].totalHarga {
					b[i] = b[i-1]
					i--
				}
			}
			b[i] = temp
			pass++
		}
	} else if tanya == "ascending" {
		for pass < n {
			i = pass
			temp = b[pass]
			if sort == "nama" {
				for i > 0 && temp.belibarang < b[i-1].belibarang {
					b[i] = b[i-1]
					i--
				}
			} else if sort == "quantity" {
				for i > 0 && temp.quantity < b[i-1].quantity {
					b[i] = b[i-1]
					i--
				}
			} else if sort == "harga" {
				for i > 0 && temp.totalHarga < b[i-1].totalHarga {
					b[i] = b[i-1]
					i--
				}
			}
			b[i] = temp
			pass++
		}
	} else {
		fmt.Print("Masukan tidak valid, input lagi (ascending/descending): ")
		fmt.Scan(&tanya)
	}
	fmt.Println(strings.Repeat("=", 150))
	fmt.Println("           Data barang setelah sorting           ")
	tampilkanData(b, n)
}

// Search //

func SeqTransaksi(b arrTran, n int, x string) int {
	var i, found int
	found = -1
	i = 0
	for i < n && found == -1 {
		if b[i].belibarang == x {
			found = i
		}
		i++
	}
	return found
}

func SeqBarang(A arrBarang, n int, x string) int {
	var i, found int
	found = -1
	i = 0
	for i < n && found == -1 {
		if A[i].nama == x {
			found = i
		}
		i++
	}
	return found
}

func cariKategori(A arrBarang, n int, kategori string) {
	i := 0
	fmt.Println("================================================")
	for i < n {
		if A[i].kategori == kategori {
			fmt.Println(A[i].nama, A[i].hargaAwal, A[i].kategori, A[i].hargaJual, A[i].stok)
		}
		i++
	}
	fmt.Println("================================================")
}

func cariBarang(A arrBarang, n int, nama string) {
	var i, found int
	found = -1
	i = 0
	for i < n && found == -1 {
		if A[i].nama == nama {
			found = i
		}
		i++
	}
	if found != -1 {
		fmt.Println("================================================")
		fmt.Println(A[found].nama, A[found].hargaAwal, A[found].kategori, A[found].hargaJual, A[found].stok)
		fmt.Println("================================================")
	} else {
		fmt.Println("Data barang tidak ditemukan")
	}
}

func cariHarga(A arrBarang, n int, z string) {
	var x, y int
	i := 0
	if z == "hargaJual" {
		fmt.Print("Range hargaJual: ")
		fmt.Scan(&x, &y)
		fmt.Println("================================================")
		for i < n {
			if A[i].hargaJual >= x && A[i].hargaJual <= y {
				fmt.Println(A[i].nama, A[i].hargaAwal, A[i].kategori, A[i].hargaJual, A[i].stok)
			}
			i++
		}
		fmt.Println("================================================")
	} else if z == "hargaAwal" {
		fmt.Print("Range hargaBeli: ")
		fmt.Scan(&x, &y)
		fmt.Println("================================================")
		for i < n {
			if A[i].hargaAwal >= x && A[i].hargaAwal <= y {
				fmt.Println(A[i].nama, A[i].hargaAwal, A[i].kategori, A[i].hargaJual, A[i].stok)
			}
			i++
		}
		fmt.Println("================================================")
	} else if z == "stok" {
		fmt.Print("Range stok: ")
		fmt.Scan(&x, &y)
		fmt.Println("================================================")
		for i < n {
			if A[i].stok >= x && A[i].stok <= y {
				fmt.Println(A[i].nama, A[i].hargaAwal, A[i].kategori, A[i].hargaJual, A[i].stok)
			}
			i++
		}
		fmt.Println("================================================")
	}
}

// function untuk menampilkan data //

func tampilkanBarang(A arrBarang, n int) {
	fmt.Println(strings.Repeat("=", 150))
	fmt.Printf("%-30s%-25s%-25s%-25s%-10s\n", "Nama Barang", "Harga Beli", "Kategori", "Harga Jual", "Stok")
	fmt.Println(strings.Repeat("=", 150))
	for i := 0; i < n; i++ {
		fmt.Printf("%-30s%-25d%-25s%-25d%-10d\n", A[i].nama, A[i].hargaAwal, A[i].kategori, A[i].hargaJual, A[i].stok)
	}
	fmt.Println(strings.Repeat("=", 150))
}

func tampilkanData(b arrTran, n int) {
	var barang, jumlah int
	for i := 0; i < n; i++ {
		b[0].total += b[i].totalHarga
		barang += 1
		jumlah += b[i].quantity
	}
	fmt.Println(strings.Repeat("=", 150))
	fmt.Printf("%-10s%-40s%-15s%-30s\n", "Urutan", "Nama Barang ", "Quantity ", "Bayar ")
	fmt.Println(strings.Repeat("=", 150))
	for i := 0; i < n; i++ {
		fmt.Printf("%-10d%-40s%-15d%-30d\n", i+1, b[i].belibarang, b[i].quantity, b[i].totalHarga)
	}
	fmt.Println(strings.Repeat("=", 150))
	fmt.Printf("%-50d%-15d%-30d\n", barang, jumlah, b[0].total)
	fmt.Println(strings.Repeat("=", 150))
}

// tampilkan barang terjual terbanyak //

func terjualTerbanyak(a arrBarang, n int) {
	var pass, i int
	var temp barang
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = a[pass]
		for i > 0 && temp.terjual > a[i-1].terjual {
			a[i] = a[i-1]
			i = i - 1
		}
		a[i] = temp
		pass++
	}
	if n > 5 {
		n = 5
	}
	fmt.Println(strings.Repeat("=", 100))
	fmt.Println("        Barang yang paling banyak terjual        ")
	fmt.Println(strings.Repeat("-", 100))
	for k := 0; k < n; k++ {
		fmt.Print(k+1, ". ", a[k].nama, " dengan kategori ", a[k].kategori, " terjual sebanyak ", a[k].terjual, "\n")
	}
	fmt.Println(strings.Repeat("=", 100))
}

func main() {
	var n, q int
	var A arrBarang
	var b arrTran

	menu(A, b, n, q)
}
