package parser

import (
	"fmt"
	"bufio"
	"errors"
	"os"
	"strconv"
)

//Prompts user for a single input, and parses accordingly, returning ints.
func PromptForSingleInput(scanner *bufio.Scanner, min uint64, max uint64) (uint64, error) {

	fmt.Println("\n--------------------\nENTER '0' TO EXIT")
	fmt.Println("\nPlease enter the number of balls in the clock (27-127):")

	for scanner.Scan() {
		var numBalls uint64 //uint64 instead of uint8 due to strconv.ParseUint returning a uint64
		var err error
		input := scanner.Text()
		if numBalls, err = strconv.ParseUint(input, 10, 8); err != nil {
			msg := fmt.Sprintf("Failed to parse input, \"%s\", as uint8", input)
			fmt.Fprintf(os.Stderr, msg)
			return 0, errors.New(msg)
		}

		if numBalls > max {
			msg := fmt.Sprintf("Too many balls declared, %d > %d", numBalls, max)
			fmt.Fprintln(os.Stderr, msg)
			return 0, errors.New(msg)
		} else if numBalls < min {
			msg := fmt.Sprintf("Too few balls declared, %d < %d", numBalls, min)
			fmt.Fprintln(os.Stderr, msg)
			return 0, errors.New(msg)
		} else {
				return numBalls, nil
		}
	}
	
	return 0, errors.New("Unexpected scanner leakage. Scanner was not detected.")
}


