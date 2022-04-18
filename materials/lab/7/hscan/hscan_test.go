// Optional Todo

package hscan

import (
	"testing"
)

func TestGuessSingle(t *testing.T) {
	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "C:/Users/Will/Desktop/School/COSC/Cyber/course-materials/materials/lab/7/main/wordlist.txt") // Currently function returns only number of open ports
	want := "Nickelback4life"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}

func TestGenHashMaps(t *testing.T) {
	GenHashMaps("C:/Users/Will/Desktop/School/COSC/Cyber/course-materials/materials/lab/7/main/wordlist.txt")
	got, _ := GetMD5("9701a1c165dd9420816bfec5edd6c2b1")
	dontWant := ""
	if got == dontWant {
		t.Errorf("you failed the test for md5")
	}

}

func TestSuite(t *testing.T) {
	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "C:/Users/Will/Desktop/School/COSC/Cyber/course-materials/materials/lab/7/main/wordlist.txt") // Currently function returns only number of open ports
	want := "Nickelback4life"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
	got2 := GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "C:/Users/Will/Desktop/School/COSC/Cyber/course-materials/materials/lab/7/main/wordlist.txt") // Currently function returns only number of open ports
	want2 := "Nickelback4life"
	if got2 != want {
		t.Errorf("got %s, wanted %s", got2, want2)
	}

}
