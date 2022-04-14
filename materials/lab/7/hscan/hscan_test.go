// Optional Todo

package hscan

import (
	"testing"
)

func TestGuessSingle(t *testing.T) {
	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "C:/Users/Will/Desktop/School/COSC/Cyber/course-materials/materials/lab/7/main/wordlist.txt") // Currently function returns only number of open ports
	want := "password"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}

func TestGenHashMaps(t *testing.T) {
	GenHashMaps("C:/Users/Will/Desktop/School/COSC/Cyber/course-materials/materials/lab/7/main/wordlist.txt")

}
