package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lSkor1, lSkor2, lSkor3, lumbaSkor int
	var kSkor1, kSkor2, kSkor3, koalaSkor int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan skor Pertama Tim Lumba-lumba: ")
	lSkor1 = getInput(reader)

	fmt.Print("Masukkan skor Pertama Tim Koala: ")
	kSkor1 = getInput(reader)

	fmt.Print("Masukkan skor Kedua Tim Lumba-lumba: ")
	lSkor2 = getInput(reader)

	fmt.Print("Masukkan skor Kedua Tim Koala: ")
	kSkor2 = getInput(reader)

	fmt.Print("Masukkan skor Ketiga Tim Lumba-lumba: ")
	lSkor3 = getInput(reader)

	fmt.Print("Masukkan skor Ketiga Tim Koala: ")
	kSkor3 = getInput(reader)

	lumbaSkor = lSkor1 + lSkor2 + lSkor3
	fmt.Printf("\nTotal Skor Tim Lumba-lumba: %d\n", lumbaSkor)
	koalaSkor = kSkor1 + kSkor2 + kSkor3
	fmt.Printf("Total Skor Tim Koala: %d\n", koalaSkor)

	lumbaLumbaRataRata := float64(lumbaSkor) / 3
	koalaRataRata := float64(koalaSkor) / 3
	fmt.Printf("\nSkor rata-rata tim Lumba-lumba: %f\n", lumbaLumbaRataRata)
	fmt.Printf("Skor rata-rata tim Koala: %f\n", koalaRataRata)

	// Data 1
	if lumbaLumbaRataRata > koalaRataRata {
		fmt.Println("\nTim Lumba-lumba Menang!")
	} else if koalaRataRata > lumbaLumbaRataRata {
		fmt.Println("\nTim Koala Menang!")
	} else {
		fmt.Println("\nPertandingan Seri")
	}

	// Data Bonus 1
	//if lumbaLumbaRataRata > koalaRataRata && lumbaLumbaRataRata >= 100 {
	//	fmt.Println("\nTim Lumba-lumba menang!")
	//} else if koalaRataRata > lumbaLumbaRataRata && koalaRataRata >= 100 {
	//	fmt.Println("\nTim Koala menang!")
	//} else {
	//	fmt.Println("\nTidak ada pemenang karena skor di bawah 100.")
	//}

	// Data Bonus 2
	//if lumbaLumbaRataRata > koalaRataRata && lumbaLumbaRataRata >= 100 {
	//	fmt.Println("\nTim Lumba-lumba menang!")
	//} else if koalaRataRata > lumbaLumbaRataRata && koalaRataRata >= 100 {
	//	fmt.Println("\nTim Koala menang!")
	//} else if (lumbaLumbaRataRata == koalaRataRata && lumbaLumbaRataRata >= 100) ||
	//	(koalaRataRata == lumbaLumbaRataRata && koalaRataRata >= 100) {
	//	fmt.Println("\nPertandingan Seri")
	//} else {
	//	fmt.Println("\nTidak ada pemenang karena skor di bawah 100.")
	//}
}

func getInput(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	value, _ := strconv.Atoi(input)
	return value
}
