package day08

import (
	_ "embed"
	"fmt"
	"masonheverett/advent-of-go/util"
	"strings"
)

//go:embed input.txt
var input string

type Entry struct {
	patterns []string
	output   []string
}

func Solve() {
	entries := parseInput()
	part1(entries)
	part2(entries)
}

func parseInput() []Entry {
	lines := strings.Split(input, "\n")
	entries := make([]Entry, len(lines))
	for i, line := range lines {
		sections := strings.Split(line, " | ")
		entries[i].patterns = strings.Split(sections[0], " ")
		entries[i].output = strings.Split(sections[1], " ")
		for j, pattern := range entries[i].patterns {
			entries[i].patterns[j] = util.Sort(pattern)
		}
		for j, digit := range entries[i].output {
			entries[i].output[j] = util.Sort(digit)
		}
	}
	return entries
}

func part1(entries []Entry) {
	count := 0
	for _, entry := range entries {
		for _, digit := range entry.output {
			switch len(digit) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	fmt.Println(count)
}

func part2(entries []Entry) {
	total := 0
	for _, entry := range entries {
		// Identify 1, 4, 7, and 8
		var onePattern, fourPattern string
		digitMap := make(map[string]int)
		for _, pattern := range entry.patterns {
			switch len(pattern) {
			case 2:
				digitMap[pattern] = 1
				onePattern = pattern
			case 3:
				digitMap[pattern] = 7
			case 4:
				digitMap[pattern] = 4
				fourPattern = pattern
			case 7:
				digitMap[pattern] = 8
			}
		}
		// Identify the segment mappings
		segCounts := map[rune]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0}
		for _, pattern := range entry.patterns {
			for _, letter := range pattern {
				segCounts[letter]++
			}
		}
		segMap := make(map[string]rune)
		for letter, count := range segCounts {
			switch count {
			case 4:
				segMap["bottom-left"] = letter
			case 6:
				segMap["top-left"] = letter
			case 7:
				if strings.ContainsRune(fourPattern, letter) {
					segMap["middle"] = letter
				} else {
					segMap["bottom"] = letter
				}
			case 8:
				if strings.ContainsRune(onePattern, letter) {
					segMap["top-right"] = letter
				} else {
					segMap["top"] = letter
				}
			case 9:
				segMap["bottom-right"] = letter
			}
		}
		// Identify 0, 2, 3, 5, 6, and 9
		for _, pattern := range entry.patterns {
			switch len(pattern) {
			case 5:
				switch {
				case !strings.ContainsRune(pattern, segMap["bottom-right"]):
					digitMap[pattern] = 2
				case !strings.ContainsRune(pattern, segMap["top-right"]):
					digitMap[pattern] = 5
				default:
					digitMap[pattern] = 3
				}
			case 6:
				switch {
				case !strings.ContainsRune(pattern, segMap["middle"]):
					digitMap[pattern] = 0
				case !strings.ContainsRune(pattern, segMap["top-right"]):
					digitMap[pattern] = 6
				case !strings.ContainsRune(pattern, segMap["bottom-left"]):
					digitMap[pattern] = 9
				}
			}
		}
		// Determine output
		output := 0
		for _, digit := range entry.output {
			output *= 10
			output += digitMap[digit]
		}
		total += output
	}
	fmt.Println(total)
}
