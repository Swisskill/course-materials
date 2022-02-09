// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage: This is a port scanner.
/*
To run, you can either do
        cd C:\Users\wrbra\Desktop\COSC\Cyber\course-materials\materials\lab\2\bhg-scanner\main\
        go build
        ./main
    And then follow the prompts or
        cd C:\Users\wrbra\Desktop\COSC\Cyber\course-materials\materials\lab\2\bhg-scanner\scanner
        go test
*/
// Will Brant, 2/3/22
// {TODO 1: FILL IN}

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

//TODO 3 : ADD closed ports; currently code only tracks open ports

func worker(ports, results chan int, ad string) {
	for p := range ports {
		address := fmt.Sprintf("%s%d", ad, p)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second) // TODO 2 : REPLACE THIS WITH DialTimeout (before testing!)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func PortScanner(ad string, range1 int, range2 int) (int, int) {

	var openports []int // notice the capitalization here. access limited!
	var closports int
	ports := make(chan int, 100) // TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE (am i doing this right?)
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, ad)
	}

	go func() {
		for i := range1; i <= range2; i++ {
			ports <- i
		}
	}()

	for i := range1; i <= range2; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
			//fmt.Printf("Port number %d is open\n", port)
		} else {
			closports++ // = append(closports, port)
			//fmt.Printf("Port number %d is closed\n", i)
		}
	}
	if len(openports) == 0 {
		fmt.Println("Sorry. There are no open ports")
	}

	close(ports)
	close(results)
	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("%d %s\n", port, protoMap[port])

	}
	switch {
	case len(openports) == 1:
		fmt.Printf("There was %d open ports and %d closed ports\n", len(openports), closports)
	case len(openports) == 0:
		fmt.Printf("There were no open ports, but %d closed ports\n", closports)
	case len(openports) > 1:
		fmt.Printf("There were %d open ports and %d closed ports\n", len(openports), closports)
	}

	return len(openports), closports
}

//TODO 5 : Enhance the output for easier consumption, include closed ports
//return len(openports)
//return len(closports)
// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
//does this change this

//let's take in a parameter for what ports we are scanning for
//add a data structure that tells what each port does
//
// TODO 6 : Return total number of ports scanned (number open, number closed);
//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
