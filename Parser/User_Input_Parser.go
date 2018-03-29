package parser

import (
	"fmt"
	"bufio"
	"errors"
	"os"
	"strconv"
)


// func Blah(int) uint
// if int = 1
// 	PromptForSingleInput

const MAX_NUM_MENU_CHOICE = 2;
const MIN_NUM_MENU_CHOICE = 0;

func MainMenuPrompt(scanner *bufio.Scanner) (uint64, error) {
	fmt.Println("\n-----MAIN MENU-----")
	fmt.Println("1.) Single Input")
	fmt.Println("2.) Dual Input (Balls & Mins)")
	fmt.Println("0.) Exit")
	fmt.Println("\n-----PLEASE ENTER YOUR CHOICE BELOW-----")

	for scanner.Scan() {
		var selection uint64
		var err error
		input := scanner.Text()
		if selection, err = strconv.ParseUint(input, 10, 8); err != nil {
			msg := fmt.Sprintf("Failed to parse input, \"%s\", as uint8. Please, select a proper choice", input)
			fmt.Fprintf(os.Stderr, msg)
		}

		if selection > MAX_NUM_MENU_CHOICE || selection < MIN_NUM_MENU_CHOICE {
			msg := fmt.Sprintf("Choice %d does not exist in menu, please select a proper choice. \n", selection)
			fmt.Fprintf(os.Stderr, msg)
		} else {
			return selection, nil
		}

	}

	return 0, errors.New("Unexpected Main Menu Error")

}

//Prompts user for a single input, and parses accordingly, returning ints.
func PromptForSingleInput(scanner *bufio.Scanner, min uint64, max uint64) (uint64, error) {

	fmt.Println("\n---WELCOME TO THE SINGLE INPUT MODE---")
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


