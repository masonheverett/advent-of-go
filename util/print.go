package util

import "fmt"

func PrintHeader(year, day, part int) {
	fmt.Printf("\n*----------------------------*\n")
	fmt.Printf("| Year %4d, Day %02d, Part %02d |\n", year, day, part)
	fmt.Printf("*----------------------------*\n\n")
}
