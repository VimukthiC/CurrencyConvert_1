package main

import "fmt"

var valueOfDoller float64 = 180.00

func main() {

	fmt.Println("lkr convert to dollar input number 1")
	fmt.Println("dollar convert to lkr input number 2")

	var input int
	fmt.Print("\nEnter input number :")
	fmt.Scan(&input)

	switch input {
	case 1:
		var i float64
		fmt.Print("Input Rupees : ")
		fmt.Scan(&i)
		fmt.Print("Doller amount : ", i/valueOfDoller)
	case 2:
		var i float64
		fmt.Print("Input Doller : ")
		fmt.Scan(&i)
		fmt.Print("Rupees amount : ", i*valueOfDoller)
	default:
		fmt.Println("Please use 0 or 1 as per your requirement")
	}

}
