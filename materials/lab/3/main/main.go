// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"shodan/shodan"
)

func main() {
	os.Setenv("SHODAN_API_KEY", "gIiBqdKhe0Pl6yUkmplFnnp2HxK2befL")
	var yes string
	fmt.Println("Type Y for documentation on usage or press enter to continue")
	fmt.Scanln(&yes)
	if yes == "Y" {
		fmt.Println("-------------------------Welcome to Shodan, Terminal Edition-------------------------")
		fmt.Println("You will be asked go supply filters to search by. The available filters are")
		fmt.Println("os:\ntimestamp:\nisp:\nasn:\nhostnames:\nlocation:\nip:\ndomains:\norg:\ndata:\nport:\nipstring\n")
		fmt.Println("Furthermore, location can have a few different ways to search by. These include:")
		fmt.Println("city\nregion_code:\narea_code:\nlongitude:\ncountry_code3:\ncountry_name:\npostal_code:\ndma_code:\ncountry_code:\nlatitude:\n")
		fmt.Println("Now notice the colon after every filter. You'll need to make sure you have those. An example:")
		fmt.Println("city:Chicago or port:80\n------------------------------------------------------------------------\nIf you have any questions, please contact support at wbrant@uwyo.edu")
		fmt.Println("------------------------------------------------------------------------\n")
	}
	var ad int
	for i := 0; i < 1; i++ {
		fmt.Println("Please enter the number of filters you would like to use")
		fmt.Scanln(&ad)
		if ad < 1 || ad > 12 {
			fmt.Println("The minimum is 1 filter and the maximum is 12 filters.")
			i--
		}
		fmt.Println("--------------------------------------------------")
	}
	var mySlic []string
	var slicer string
	//fmt.Println("Please enter your filters one at a time\n")
	for i := 0; i < ad; i++ {
		fmt.Printf("Enter filter number %d:\n", i+1)
		fmt.Scanln(&slicer)
		mySlic = append(mySlic, slicer)
	}
	var pag int
	for i := 0; i < 1; i++ {
		fmt.Println("\nPlease enter the number of pages you would like to view")
		fmt.Scanln(&pag)
		if ad < 1 {
			fmt.Println("The minimum is 1 page")
			i--
		}
	}

	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hostSearch, err := s.HostSearch(mySlic)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Host Data Dump\n")
	ind := 0
	for _, host := range hostSearch.Matches {
		ind++
		fmt.Println("==== start ", host.IPString, "====")
		h, _ := json.Marshal(host)
		fmt.Println(string(h))
		fmt.Println("==== end ", host.IPString, "====")
		if ind != pag {
			fmt.Printf("Press the Enter Key to continue to page %d\n", ind+1)
			fmt.Scanln()
		} else {
			fmt.Println("Press the Enter Key to exit")
			fmt.Scanln()
			break
		}
	}

	fmt.Printf("IP, Port\n")

	for _, host := range hostSearch.Matches {
		fmt.Printf("%s, %d\n", host.IPString, host.Port)
	}

}
