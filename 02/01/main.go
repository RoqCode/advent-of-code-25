package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line here

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

		// Check for errors during scanning
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading file: %s", err)
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
