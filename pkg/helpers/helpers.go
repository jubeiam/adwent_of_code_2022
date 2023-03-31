package helpers

import (
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadFile(filePath string) string {
	cwd, _ := os.Getwd()
	dat, err := os.ReadFile(cwd + filePath)
	Check(err)
	return string(dat)
}

func SplitByNewline(data string) []string {
	return strings.Split(data, "\n")
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
