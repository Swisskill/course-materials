package main

import (
	"fmt"
	"hscan/hscan"
)

func main() {

	//To test this with other password files youre going to have to hash
	var md5hash = "5f4dcc3b5aa765d61d8327deb882cf99"                                    //"77f62e3524cd583d698d51fa24fdff4f"
	var sha256hash = "5E884898DA28047151D0E56F8DC6292773603D0D6AABBDD62A11EF721D1542D8" //"95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced"

	//TODO - Try to find these (you may or may not based on your password lists)
	//var drmike1 = "90f2c9c53f66540e67349e0ab83d8cd0"
	//var drmike2 = "1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032"

	// NON CODE - TODO
	// Download and use bigger password file from: https://weakpass.com/wordlist/tiny  (want to push yourself try /small ; to easy? /big )

	//TODO Grab the file to use from the command line instead; look at previous lab (e.g., #3 ) for examples of grabbing info from command line
	var file string
	fmt.Println("Tread Cautiously:\nEnter in the file")
	fmt.Scanln(&file)
	fmt.Println()
	fmt.Print(" single md5  ")
	fmt.Println()
	hscan.GuessSingle(md5hash, file)
	fmt.Print(" single sha  ")
	fmt.Println()
	hscan.GuessSingle(sha256hash, file)
	fmt.Println()
	fmt.Print(" GenHashMaps  ")
	hscan.GenHashMaps(file)
	fmt.Println(hscan.GetSHA(sha256hash))
	fmt.Println(hscan.GetMD5(md5hash))
}
