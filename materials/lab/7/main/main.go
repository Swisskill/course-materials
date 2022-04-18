package main

import (
	"fmt"
	"hscan/hscan"
	"time"
)

func main() {

	//To test this with other password files youre going to have to hash
	//var md5hash = "5ee3ca63ca4240d1bac8921c50221f94" //"5f4dcc3b5aa765d61d8327deb882cf99"
	//"77f62e3524cd583d698d51fa24fdff4f"
	//var sha256hash = "8b7d3d77384a9756e3596dd3e81bca2e370c0890556f2f6d7a9cbf60493e68e4" //"5E884898DA28047151D0E56F8DC6292773603D0D6AABBDD62A11EF721D1542D8" //"95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced"
	//DONE
	//TODO - Try to find these (you may or may not based on your password lists)
	var drmike1 = "90f2c9c53f66540e67349e0ab83d8cd0"
	var drmike2 = "1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032"

	// NON CODE - TODO DONE
	// Download and use bigger password file from: https://weakpass.com/wordlist/tiny  (want to push yourself try /small ; to easy? /big )
	//DONE
	//TODO Grab the file to use from the command line instead; look at previous lab (e.g., #3 ) for examples of grabbing info from command line
	var file string
	var preFile string
	var tuiHash1 string
	var tuiHash2 string
	var preHash string
	fmt.Println("Tread Cautiously:\nEnter in the file or type \"default\"")
	fmt.Scanln(&preFile)
	if preFile == "default" {
		file = "toplist.txt"
	} else {
		file = preFile
	}
	fmt.Println("\nEnter in hash to scan or type \"default\"")
	fmt.Scanln(&preHash)
	fmt.Println()
	if preHash == "default" {
		start := time.Now()
		tuiHash1 = drmike1
		tuiHash2 = drmike2
		hscan.GuessSingle(tuiHash1, file)
		hscan.GuessSingle(tuiHash2, file)
		//hscan.GenHashMaps(file)
		hscan.GenHashMapsC(file)
		fmt.Println(hscan.GetSHAC(drmike2))
		//fmt.Println(hscan.GetSHA(drmike2))
		//fmt.Println(hscan.GetMD5(drmike1))
		fmt.Println("-----------------------")
		fmt.Printf("Time elapsed: %s\n", time.Since(start))

	} else {
		start := time.Now()
		hscan.GuessSingle(preHash, file)
		hscan.GenHashMaps(file)
		if len(preHash) == 64 {
			fmt.Println(hscan.GetSHA(preHash))
			fmt.Println("-----------------------")
			fmt.Printf("Time elapsed: %s\n", time.Since(start))

		} else if len(preHash) == 32 {
			fmt.Println(hscan.GetMD5(preHash))
			fmt.Println("-----------------------")
			fmt.Printf("Time elapsed: %s\n", time.Since(start))

		} else {
			fmt.Println("This was not a valid hash")
			fmt.Println("-----------------------")
			fmt.Printf("Time elapsed: %s\n", time.Since(start))

		}
	}

}

//start := time.Now()

//... do something
