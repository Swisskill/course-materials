package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
)

//==========================================================================\\

var shalookup map[string]string
var md5lookup map[string]string

func GuessSingle(sourceHash string, filename string) string {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		// TODO - From the length of the hash you should know which one of these to check ...
		// add a check and logicial structure

		hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if hash == sourceHash {
			fmt.Printf("[+] Password found (MD5): %s\n", password)
			return password
		}

		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		if hash == sourceHash {
			fmt.Printf("[+] Password found (SHA-256): %s\n", password)
			return password
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return "nil"
}

func GenHashMaps(filename string) {

	//TODO
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time
	//OPTIONAL -- Can you use workers to make this even faster

	//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()
		go hash1 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		go hash2 = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		

	}

}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}
}

//TODO done?
func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}
}


func generateNumbersMap() {
	wg := sync.WaitGroup{}
    // Write.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			numbers[i] = i
		}(i)
	}
    // Read.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Print(numbers[i])
		}(i)
	}

	wg.Wait()
}