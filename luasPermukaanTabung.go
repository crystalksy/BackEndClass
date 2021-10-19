package main

import "fmt"

func main() {
	const phi = 3.14
	var t, r float32

	fmt.Print("Masukkan tinggi tabung : ")
	fmt.Scan(&t)
	fmt.Print("Masukkan jari-jari tabung : ")
	fmt.Scan(&r)

	luas := 2 * phi * r * (r + t)

	fmt.Println("Luas permukaan tabung adalah", luas)
}
