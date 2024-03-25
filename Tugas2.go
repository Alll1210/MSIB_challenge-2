package main

import "fmt"

func main() {
	var bbMark, tbMark, bbJohn, tbJohn float64

	fmt.Print("Masukkan berat badan Mark (dalam kg): ")
	fmt.Scan(&bbMark)
	fmt.Print("Masukkan tinggi badan Mark (dalam meters): ")
	fmt.Scan(&tbMark)

	fmt.Print("Masukkan berat badan John (dalam kg): ")
	fmt.Scan(&bbJohn)
	fmt.Print("Masukkan tinggi badan John (dalam meters): ")
	fmt.Scan(&tbJohn)

	markBMIPowerFormula := bbMark / (tbMark * tbMark)
	johnBMIPowerFormula := bbJohn / (tbJohn * tbJohn)

	markHigherBMI := markBMIPowerFormula > johnBMIPowerFormula

	fmt.Println("BMI Mark adalah:", markBMIPowerFormula)
	fmt.Println("BMI John adalah:", johnBMIPowerFormula)
	fmt.Println("Apakah Mark memiliki BMI lebih tinggi dari John?", markHigherBMI)
}
