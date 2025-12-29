package main

// 1. read from input.txt
// implement "dialVal"
// "R" = +
// "L" = -
// check if dialVal shows 0 -> increment counter
// return counter

import (
	"fmt"
	"log"
	"strconv"

	"advent-of-code-25/internal/input"
)

var (
	counter = 0
	dialVal = 50
)

func main() {
	lines, err := input.ReadLines("01/input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %s", err)
	}

	for _, line := range lines {
		direction := (Direction)(line[:1])

		val, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Can't convert this to an int!")
		}

		err = turnDialVal(direction, val)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Printf("Password: %d\n", counter)
}

type Direction string

const (
	DirectionR Direction = "R"
	DirectionL Direction = "L"
)

func turnDialVal(d Direction, val int) error {
	// Validate direction
	if d != DirectionR && d != DirectionL {
		return fmt.Errorf("invalid direction: %s", d)
	}

	switch d {
	case DirectionR:

		for i := val; i > 0; i-- {
			dialVal++

			if dialVal == 100 {
				dialVal = 0
			}

			if dialVal == 0 {
				counter++
			}
		}
	case DirectionL:
		for i := val; i > 0; i-- {
			dialVal--

			if dialVal == -1 {
				dialVal = 99
			}

			if dialVal == 0 {
				counter++
			}
		}
	}

	fmt.Printf("Turning Dial by %d in Direction %s, Dial shows %d\n", val, d, dialVal)

	return nil
}
