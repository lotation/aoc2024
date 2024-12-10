package utils

import (
	"log"
	"os"
	"strconv"
)

func Fopen(filepath string) *os.File {
	fp, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Error opening input file %s: %v", filepath, err)
	}
	defer func() {
		if err := fp.Close(); err != nil {
			log.Panic(err)
		}
	}()
	return fp
}

func ToInt(val string) int {
	num, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		log.Fatalf("Error converting value %s to uint32: %v.", val, err)
	}
	return int(num)
}

func ToIntSlice(slice []string) []int {
	ints := make([]int, len(slice))
	for i, s := range slice {
		ints[i] = ToInt(s)
	}
	return ints
}
