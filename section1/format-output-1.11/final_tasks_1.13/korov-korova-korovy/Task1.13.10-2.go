package main

import "fmt"

func main() {
	var cows, ref int
	var suffix = ""
	var prefix = "korov"
	fmt.Scanln(&cows)
	ref = cows
	if cows > 19 {
		ref = cows % 10
	}
	switch ref {
	case 1:
		suffix = "a"
	case 2, 3, 4:
		suffix = "y"
	}
	fmt.Printf("%d %s%s", cows, prefix, suffix)
}
