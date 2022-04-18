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
	jimothy, ward := GenHashMaps("C:/Users/Will/Desktop/School/COSC/Cyber/course-materials/materials/lab/7/main/toplist.txt")
	if jimothy != 303868 || ward != 303868 {
		t.Errorf("got %d and %d but wanted 303872", jimothy, ward)
	}
}

func TestGenHashMapsC(t *testing.T) {
	jimothy, ward := GenHashMapsC("C:/Users/Will/Desktop/School/COSC/Cyber/course-materials/materials/lab/7/main/toplist.txt")
	if jimothy != 303868 || ward != 303868 {
		t.Errorf("got %d and %d but wanted 303872", jimothy, ward)
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
