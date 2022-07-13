package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLinesAsInts(fname string) []int {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var ints []int
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		i, err := strconv.Atoi(scan.Text())
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func ReadLinesAsStrings(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var strs []string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		str := scan.Text()
		strs = append(strs, str)
	}
	return strs
}

func ReadLinesAsStringSlices(fname string) [][]string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var stringSlices [][]string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		stringSlices = append(stringSlices, strings.Split(scan.Text(), " "))
	}
	return stringSlices
}
