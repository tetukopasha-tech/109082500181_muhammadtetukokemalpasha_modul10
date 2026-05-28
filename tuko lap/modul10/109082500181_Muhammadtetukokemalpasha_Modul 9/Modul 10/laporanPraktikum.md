# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>
<p align="center">[Muhammadtetukokemalpasha] - [109082500181]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var berat [1000]float64
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}

		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Println(min, max)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/tetukopasha-tech/109082500181_muhammadtetukokemalpasha_modul10/blob/main/tuko%20lap/modul10/109082500181_Muhammadtetukokemalpasha_Modul%209/Modul%2010/output/soal1.png)
[Penjelasan:Program Go di atas digunakan untuk mencari nilai berat badan minimum dan maksimum dari sejumlah $n$ data yang dimasukkan oleh pengguna. Pertama, program mendeklarasikan sebuah array bernama `berat` dengan kapasitas hingga 1000 elemen dan membaca total data yang akan diinput (`n`), lalu mengisi array tersebut menggunakan perulangan pertama sesuai dengan jumlah $n$. Selanjutnya, variabel `min` dan `max` diinisialisasi menggunakan nilai dari elemen pertama array (`berat[0]`) sebagai acuan awal. Melalui perulangan kedua yang dimulai dari indeks 1 hingga $n-1$, program membandingkan setiap elemen array; jika ditemukan nilai yang lebih kecil dari `min`, maka `min` akan diperbarui, dan jika ditemukan nilai yang lebih besar dari `max`, maka `max` akan diperbarui. Di akhir program, nilai terkecil dan terbesar yang berhasil ditemukan dicetak ke layar secara berdampingan. ]

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var berat [1000]float64
	var totalWadah [1000]float64
	var x, y int

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	indeks := 0

	for i := 0; i < jumlahWadah; i++ {
		total := 0.0

		for j := 0; j < y && indeks < x; j++ {
			total += berat[indeks]
			indeks++
		}

		totalWadah[i] = total
	}

	totalSemua := 0.0

	for i := 0; i < jumlahWadah; i++ {
		totalSemua += totalWadah[i]
	}

	rata := totalSemua / float64(jumlahWadah)

	for i := 0; i < jumlahWadah; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(totalWadah[i])
	}

	fmt.Println()
	fmt.Println(rata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/tetukopasha-tech/109082500181_muhammadtetukokemalpasha_modul10/blob/main/tuko%20lap/modul10/109082500181_Muhammadtetukokemalpasha_Modul%209/Modul%2010/output/soal2.png)
[Program Go di atas berfungsi untuk mengelompokkan data berat sebanyak $x$ elemen ke dalam beberapa wadah yang masing-masing berkapasitas maksimal $y$ elemen, lalu menghitung total berat tiap wadah serta rata-rata berat keseluruhannya. Pertama, program menerima input nilai $x$ dan $y$, membaca seluruh data berat ke dalam array, lalu menentukan jumlah wadah yang dibutuhkan dengan membagi $x$ dengan $y$ (serta menambah satu wadah jika ada sisa pembagian). Selanjutnya, menggunakan perulangan bertingkat (*nested loop*), program menjumlahkan elemen-elemen berat secara berurutan sesuai kapasitas wadah dan menyimpannya ke dalam array `totalWadah`. Di akhir proses, program menghitung rata-rata berat per wadah dengan membagi total seluruh berat dengan jumlah wadah, lalu mencetak total berat masing-masing wadah dalam satu baris yang dipisahkan spasi diikuti dengan nilai rata-ratanya di baris baru.]

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}

		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max, rata float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)
	rata = rerata(data, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/tetukopasha-tech/109082500181_muhammadtetukokemalpasha_modul10/blob/main/tuko%20lap/modul10/109082500181_Muhammadtetukokemalpasha_Modul%209/Modul%2010/output/soal3.png)
[Program Go tersebut mencatat $n$ data berat balita ke dalam *array*, lalu menghitung nilai minimum, maksimum, dan rata-ratanya menggunakan fungsi terpisah. Fungsi `main` bertugas menerima input data, prosedur `hitungMinMax` mencari berat terkecil dan terbesar dengan membandingkan tiap elemen melalui *pointer*, dan fungsi `rerata` menghitung rata-rata dengan membagi total berat dengan jumlah balita. Hasil akhirnya kemudian dicetak dengan format dua angka di belakang koma.			]

