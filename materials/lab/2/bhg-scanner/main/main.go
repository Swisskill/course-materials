package main

import (
	"bhg-scanner/scanner"
	"fmt"
)

func main() {
	var ad string
	var range1, range2 int
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("* * * * * * * * * * * Tread -- Carefully * * * * * * * * * * *")
	fmt.Println("Provide an address or type 'default' to use the default server")
	fmt.Println("--------------------------------------------------------------")
	fmt.Scanln(&ad)
	fmt.Println("")
	if ad == "default" {
		ad = "scanme.nmap.org:"
	}

	for i := 0; i < 1; i++ {
		fmt.Println("--------------------------------------------------------------")
		fmt.Println(" Specify the first digit of range")
		fmt.Println("--------------------------------------------------------------")
		fmt.Scanln(&range1)
		fmt.Println("")
		switch {
		case range1 <= 0:
			fmt.Printf("The first digit can't be %d, it needs to be positive\n\n", range1)
			i--
		case range1 > 65506:
			fmt.Printf("The first digit can't be %d, it needs to be less than 65506\n\n", range1)
			i--
		}
	}

	for i := 0; i < 1; i++ {
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("Specify the second digit of range")
		fmt.Println("--------------------------------------------------------------")
		fmt.Scanln(&range2)
		fmt.Println("")
		switch {
		case range2 <= 0:
			fmt.Printf("The second digit can't be %d, it needs to be positive and greater than 0\n\n", range2)
			i--
		case range2 > 65506:
			fmt.Printf("The second digit can't be %d, it needs to be less than 65506\n\n", range2)
			i--
		case range2 < range1:
			fmt.Printf("The second digit can't be %d, it needs to be bigger than %d\n\n", range2, range1)
			i--
		}
	}
	scanner.PortScanner(ad, range1, range2)
	//scanner.PortScanner(ad)

}
