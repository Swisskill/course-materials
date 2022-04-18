package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

//==========================================================================//

var shalookup map[string]string
var md5lookup map[string]string
var shaGo sync.Map
var md5Go sync.Map
var shaCount = 0
var md5Count = 0

func GuessSingle(sourceHash string, filename string) string {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		// TODO DONE- From the length of the hash you should know which one of these to check ...
		// add a check and logicial structure
		if len(sourceHash) == 32 {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
				return password
			}
		}
		if len(sourceHash) == 64 {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
				return password
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return "nil"
}

func GenHashMapsC(filename string) (int, int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var phore sync.WaitGroup
	for scanner.Scan() {
		password := scanner.Text()
		phore.Add(2)
		go shaHelp(password, &phore)
		go md5Help(password, &phore)

	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return shaCount, md5Count
}
func shaHelp(password string, phore1 *sync.WaitGroup) {
	shaGo.Store(fmt.Sprintf("%x", sha256.Sum256([]byte(password))), password) //[fmt.Sprintf("%x", sha256.Sum256([]byte(password)))] = password //fmt.Sprintf("%x", sha256.Sum256([]byte(password))
	shaCount++
	phore1.Done()
}
func md5Help(password string, phore2 *sync.WaitGroup) {
	md5Go.Store(fmt.Sprintf("%x", md5.Sum([]byte(password))), password) //[fmt.Sprintf("%x", sha256.Sum256([]byte(password)))] = password //fmt.Sprintf("%x", sha256.Sum256([]byte(password))
	md5Count++
	phore2.Done()
}

func GenHashMaps(filename string) (int, int) {

	//TODO DONE
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time
	//OPTIONAL -- Can you use workers to make this even faster

	//TODO DONE create a test in hscan_test.go so that you can time the performance of your implementation
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
		if shalookup == nil {
			shalookup = make(map[string]string)
		}
		if md5lookup == nil {
			md5lookup = make(map[string]string)
		}
		password := scanner.Text()
		shalookup[fmt.Sprintf("%x", sha256.Sum256([]byte(password)))] = password //fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		md5lookup[fmt.Sprintf("%x", md5.Sum([]byte(password)))] = password       //fmt.Sprintf("%x", md5.Sum([]byte(password)))

	}
	r1 := len(md5lookup)
	r2 := len(shalookup)
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return r1, r2
}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		returner := fmt.Sprintf("[+] From the general SHA function: %s", password)
		return returner, errors.New("")

	} else {

		return "", errors.New("password does not exist")

	}
}

//TODO DONE
func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		returner := fmt.Sprintf("[+] From the general MD5 function: %s", password)
		return returner, errors.New("")
	} else {

		return "", errors.New("password does not exist")

	}
}

func GetSHAC(hash string) (string, error) {
	password, ok := shaGo.Load(hash)
	if ok {
		returner := fmt.Sprintf("[+] From the general SHA function: %s", password)
		return returner, errors.New("")

	} else {

		return "", errors.New("password does not exist")

	}
}

//TODO DONE
func GetMD5C(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		returner := fmt.Sprintf("[+] From the general MD5 function: %s", password)
		return returner, errors.New("")
	} else {

		return "", errors.New("password does not exist")

	}
}

/*
type SafeNumbers struct {
	sync.RWMutex
	numbers map[int]int
}
To be able to read and write items concurrently to this structure, we need to create the responsible methods:

func (sn *SafeNumbers) Add(num int) {
	sn.Lock()
	defer sn.Unlock()
	sn.numbers[num] = num
}
Here we are basically telling to lock the numbers map, during adding of the new number to it. Other goroutines will wait until it became unlocked again.

And another method for reading:

func (sn *SafeNumbers) Get(num int) (int, error) {
	sn.RLock()
	defer sn.RUnlock()
	if number, ok := sn.numbers[num]; ok {
		return number, nil
	}
	return 0, errors.New("Number does not exists")
}
Next, let’s refactor our generateNumbersMap() function:

func generateNumbersMap() {
	wg := sync.WaitGroup{}
    // Init our "safe" numbers map struct.
	safeNumbers := &SafeNumbers{
		numbers: map[int]int{},
	}
    // Write.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeNumbers.Add(i)
		}(i)
	}
    // Read.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			number, err := safeNumbers.Get(i)
			if err != nil {
				log.Print(err)
			} else {
				log.Print(number)
			}
		}(i)
	}

	wg.Wait()
}


}
*/
