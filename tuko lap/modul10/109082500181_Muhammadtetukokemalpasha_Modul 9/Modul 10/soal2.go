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