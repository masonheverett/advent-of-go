package util

import (
	"sort"
	"strings"
)

func Sort(s string) string {
	slc := strings.Split(s, "")
	sort.Strings(slc)
	return strings.Join(slc, "")
}
