package main

import (
	"bhg-scanner/scanner"
	"fmt"
)

func main() {
	fmt.Println("Provide an address or type 'default' to use the default server")
	var ad string
	fmt.Scanln(ad)
	if ad == "default" {
		ad = "scanme.nmap.org:"
	}
	scanner.PortScanner(portMap.ProtoMap(), ad)
https: //uwyo.zoom.us/j/3077664366?pwd=dWlVTXdtZjcxVEttOWZvT2t1M1hoZz09
}
