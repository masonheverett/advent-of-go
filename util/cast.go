package util

import (
	"log"
	"strconv"
)

func DecStringToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func BinStringToUint64(s string) uint64 {
	n, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
