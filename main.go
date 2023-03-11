// Example application that uses all of the available API options.
package main

import (
	"time"
	"a71.su/gmms/tui"
	"fmt"
)

func main() {
	fmt.Println(tui.AsciiLogo())
	s := tui.GetSpinner()
	s.Suffix = " Loading the flux capacitor"
	s.Start()
	
	time.Sleep(1 * time.Second)
	s.Prefix += "Among us\n"
	time.Sleep(1 * time.Second)
	s.Prefix += "Doctorin' the TARDIS\n"
	time.Sleep(1 * time.Second)
	s.Prefix += "Trolling\n"
	time.Sleep(1 * time.Second)
	s.Prefix += "gmms is a strange name\n"
	time.Sleep(1 * time.Second)
	s.Stop()
}