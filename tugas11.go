package main

import "fmt"
import "strconv"

func main() {
	// var nilai1 = "15"
	// var nilai2 = "5"

	var nilai1 string
	var nilai2 string
	fmt.Print("masukan nilai 1 :")
	fmt.Scanln(&nilai1)
	fmt.Print("masukan nilai 2 :")
	fmt.Scanln(&nilai2)

	var num1, err = strconv.Atoi(nilai1)
	var num2, err2 = strconv.Atoi(nilai2)

	if err == nil {
		fmt.Println("nilai 1 :", num1)
	}
	if err2 == nil {
		fmt.Println("nilai 2 :", num2)
	}
	fmt.Println("hasil penjumlahan :", num1+num2)
	fmt.Println("hasil pengurangan :", num1-num2)
	fmt.Println("hasil perkalian :", num1*num2)
	fmt.Println("hasil pembagian :", num1/num2)
}
