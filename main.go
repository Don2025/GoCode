package main

import "fmt"

func main() {
	s := "aqlipcrwjcrkmp]{mw]imv]oc"
	for i, ch := range s {
		if i&1 == 0 {
			fmt.Printf("%c", ch+2)
		} else {
			fmt.Printf("%c", ch-2)
		}
	}
}
