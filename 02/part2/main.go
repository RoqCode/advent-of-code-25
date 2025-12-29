package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"advent-of-code-25/internal/input"
)

func main() {
	lines, err := input.ReadLines("02/input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %s", err)
	}

	for _, line := range lines {

		IDs := strings.Split(line, ",")

		invalidIDsSum := 0

		for _, id := range IDs {

			IDRange := strings.Split(id, "-")
			IDRangeStart, _ := strconv.Atoi(IDRange[0])
			IDRangeEnd, _ := strconv.Atoi(IDRange[1])

			for IDRangeStart <= IDRangeEnd {

				if !validateID(fmt.Sprint(IDRangeStart)) {
					invalidIDsSum += IDRangeStart
				}

				IDRangeStart++
			}

		}

		fmt.Println(invalidIDsSum)
	}
}

func validateID(id string) bool {
	n := len(id)

	if n < 2 {
		return true
	}

	for k := 1; k <= n/2; k++ {
		if n%k != 0 {
			continue
		}

		pattern := id[:k]
		isRepetition := true

		for pos := k; pos < n; pos += k {
			if id[pos:pos+k] != pattern {
				isRepetition = false
				break
			}
		}

		if isRepetition {
			return false
		}
	}

	return true
}
