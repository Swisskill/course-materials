package main

import (
	"bhg-scanner/scanner"
	"fmt"
)

func main() {
	fmt.Println("Provide an address or type 'default' to use the default server")
	var ad string
	fmt.Scanln(&ad)
	if ad == "default" {
		ad = "scanme.nmap.org:"
	}
	scanner.PortScanner(ad)

}
