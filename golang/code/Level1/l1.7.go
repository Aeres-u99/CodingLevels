package main

import "fmt"

func main() {
	var C float64
	var F float64
	var choice rune
	fmt.Println("Select Conversion you want to Perform: ")
	fmt.Println("C. Celsius to Fahrenheit")
	fmt.Println("F. Fahrenheit to Celsius")
	fmt.Scanf("%c",&choice)

	switch choice {
		case 'C', 'c':
			fmt.Println("Converting Celsius to fahrenheit")
			fmt.Println("Enter temperature as Celsius: ")
			fmt.Scan(&C)
			F = (C * 9 / 5.0) + 32
			fmt.Printf("%.2f\n", F)

		case 'F', 'f':
			fmt.Println("Converting Fahrenheit to Celsius")
			fmt.Println("Enter temperatures in Fahrenheit: ")
			fmt.Scan(&F)
			C = (F - 32) * (5 / 9.0)
			fmt.Printf("%.2f\n", C)
	}
	
}
