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
	if len(id)%2 != 0 {
		return true
	}

	segLength := len(id) / 2

	start := id[segLength:]
	end := id[:segLength]

	return start != end
}
