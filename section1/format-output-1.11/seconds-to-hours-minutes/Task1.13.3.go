package main

import "fmt"

// сколько целых часов h и целых минут m прошло с начала суток. Например, если
// k=13257=3*3600+40*60+57,
// то h=3 и m=40.
func main() {
	var seconds, hours, minutes int
	fmt.Scan(&seconds)
	hours = seconds / 3600
	minutes = (seconds - hours*3600) / 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
