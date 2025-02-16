package scanner

import (
	"fmt"
	"testing"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T) {
	fmt.Println("Alert: Tests use default condtions only. Edge cases must be tested in main")
	get, _ := PortScanner("scanme.nmap.org:", 1, 1024) // Currently function returns only number of open ports
	want := 2                                          // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns
	//consider what would happen if you parameterize the portscanner address and ports to scan

	if get != want {
		t.Errorf("got %d, wanted %d", get, want)
	}
}

func TestTotalPortsScanned(t *testing.T) {
	// THIS TEST WILL FAIL - YOU MUST MODIFY THE OUTPUT OF PortScanner()
	fmt.Println("Alert: Tests use default condtions only. Edge cases must be tested in main")
	get, got := PortScanner("scanme.nmap.org:", 1, 1024) // Currently function returns only number of open ports
	want := 1024                                         // default value; consider what would happen if you parameterize the portscanner ports to scan

	if get+got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
