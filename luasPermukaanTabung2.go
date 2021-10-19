package main

import "fmt"

func main() {
	const phi = 3.14

	fmt.Print("Masukkan tinggi dan jari-jari tabung (pisahkan dg spasi) : ")

	var t, r float64

	fmt.Scanf("%v", &t)
	fmt.Scanf("%v", &r)

	luas := 2 * phi * r * (r + t)

	fmt.Println("Luas permukaan tabung adalah", luas)
}
