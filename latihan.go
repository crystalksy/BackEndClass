package main

import "fmt"

func main() {
	//data type
	var isLogin bool = false
	fmt.Println(isLogin)

	var prime int = 7
	fmt.Println(prime)

	var decimal float64 = 45.6
	fmt.Println(decimal)

	var hello string = "Hello World"
	fmt.Println(hello)

	const pi = 3.14
	fmt.Println(pi)

	//operator and operand
	x := 1 + 2
	fmt.Println(x)

	//expression
	a := 5
	b := 6
	c := a + b
	fmt.Println(c)

	alas := 2
	tinggi := 3
	fmt.Println((alas * tinggi) / 2)

	helloworld := "hello" + " " + "world"
	fmt.Println(helloworld)

	//branching
	hour := 15
	if hour < 12 {
		fmt.Println("pagi")
	} else if hour < 18 {
		fmt.Println("sore")
	} else {
		fmt.Println("malam")
	}

	//switch
	var today int = 4
	switch today {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Invalid")
	}

	//looping
	for i := 2; i < 5; i++ {
		fmt.Println(i)
	}

	//looping w continue n break
	for i := 0; i < 5; i++ {
		if i == 1 {
			continue
		}
		if i > 3 {
			break
		}
		fmt.Println(i)
	}

	//looping over string
	// sentence := "Hello"
	// for i := 0; i <len(sentence); i++{
	// 	fmt.Printf(string(sentence[i]) + "-")
	// }
	// fmt.Println("    ----->>>>")
	// for pos, char := range sentence {
	// 	fmt.Printf("character %")
	// }
}
