package util

import (
	"log"
	"strconv"
	"strings"
)

func BinStringToInt(s string) int {
	n, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(n)
}

func DecStringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func HexStringToBitString(s string) string {
	total := ""
	for _, digit := range strings.Split(s, "") {
		n, err := strconv.ParseUint(digit, 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		bs := strconv.FormatUint(n, 2)
		for len(bs) < 4 {
			bs = "0" + bs
		}
		total += bs
	}
	return total
}
