package day16

import (
	_ "embed"
	"fmt"
	"masonheverett/advent-of-go/util"
)

//go:embed input.txt
var input string

var versionSum int

func readPacket(s string) (int, string) {
	versionSum += util.BinStringToInt(s[:3])
	typeId := util.BinStringToInt(s[3:6])
	if typeId == 4 {
		return readLiteral(s[6:])
	}
	var subVals []int
	var leftover string
	if s[6] == '0' {
		subVals, leftover = readByLength(s[22:], util.BinStringToInt(s[7:22]))
	} else {
		subVals, leftover = readByCount(s[18:], util.BinStringToInt(s[7:18]))
	}
	value := 0
	switch typeId {
	case 0:
		value = util.Sum(subVals)
	case 1:
		value = util.Product(subVals)
	case 2:
		value = util.MinInt(subVals)
	case 3:
		value = util.MaxInt(subVals)
	case 5:
		if subVals[0] > subVals[1] {
			value = 1
		}
	case 6:
		if subVals[0] < subVals[1] {
			value = 1
		}
	case 7:
		if subVals[0] == subVals[1] {
			value = 1
		}
	}
	return value, leftover
}

func readByLength(s string, length int) ([]int, string) {
	var result int
	values := make([]int, 0)
	leftover := s
	for length != len(s)-len(leftover) {
		result, leftover = readPacket(leftover)
		values = append(values, result)
	}
	return values, leftover
}

func readByCount(s string, count int) ([]int, string) {
	var result int
	values := make([]int, 0)
	leftover := s
	for i := 0; i < count; i++ {
		result, leftover = readPacket(leftover)
		values = append(values, result)
	}
	return values, leftover
}

func readLiteral(s string) (int, string) {
	literal := ""
	for i := 0; ; i += 5 {
		literal += s[i+1 : i+5]
		if s[i] == '0' {
			return util.BinStringToInt(literal), s[i+5:]
		}
	}
}

func Solve() {
	result, _ := readPacket(parseInput())
	fmt.Println(versionSum)
	fmt.Println(result)
}

func parseInput() string {
	return util.HexStringToBitString(input)
}
