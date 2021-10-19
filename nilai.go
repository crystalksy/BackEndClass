package main

import "fmt"

func main() {
	var studentScore int = 34

	switch {
	case studentScore >= 80:
		fmt.Println("Nilai A")
	case studentScore >= 65:
		fmt.Println("Nilai B")
	case studentScore >= 50:
		fmt.Println("Nilai C")
	case studentScore >= 35:
		fmt.Println("Nilai D")
	case studentScore >= 0:
		fmt.Println("Nilai E")
	case studentScore > 100 || studentScore < 0:
		fmt.Println("Nilai Invalid")
	}
}
