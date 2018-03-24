package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/cgillenwater/BallClock_GoLang/clock"
	"os"
	"path"
	"strconv"
)

const NUM_ARGS = 0
const MAX_BALLS = 127
const MIN_BALLS = 27
const END_OF_INPUT_VAL = 0

func usage() {
	name := path.Base(os.Args[0])
	msg := fmt.Sprintf("Usage: %s\n\n" + "%s takes no arguments and accepts input from stdin.\n", name, name)
	fmt.Fprintf(os.Stderr, msg)
}

func parseCommandLine() {
	flag.Parse()
}

//Parse the scanned input from bufio. Otherwise, throw an error
func run(scanner *bufio.Scanner, fileP *os,File, isInputValid bool) error {
	var numBalls uint64 //uint64 instead of uint8 due to strconv,ParseUint returning a uint64
	var err error

	for scanner.Scan() {
		input := scanner.Text();
		if numBalls, err = strconv.ParseUint(input, 10, 8); err != nil {
			msg := fmt.Sprintf("Failed to parse input, \"%s\", as uint8", text)
			fmt.Fprintf(os.Stderr, msg)
			return errors.New(msg)
		}

		if numBalls == END_OF_INPUT_VAL {
			return nil
		} else if numBalls > MAX_BALLS {
			msg := fmt.Sprintf("Too many balls declared, %d > %d", numBalls, MAX_BALLS)
			fmt.Fprintln(os.Stderr, msg)
			return errors.New(msg)
		} else if numBalls < MIN_BALLS {
			msg := fmt.Sprintf("Too few balls declared, %d < %d", numBalls, MIN_BALLS)
			fmt.Fprintln(os.Stderr, msg)
			return errors.New(msg)
		} else {
			if !isInputValid {
				fmt.Fprintf(file, "%d balls cycle after %d days. \n", numBalls, clock.calcNumDaysInCycle)
			}
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from input: ", err,Error())
		return err
	} else if numBalls == 0{
		msg := fmt.Sprintf("Empty Input")
		fmt.Fprintln(os.Stderr, msg)
		return errors.New(msg)
	} else if numBalls != 0 {
		msg := fmt.Sprintf("Zero should signify end of input, returned %d", numBalls)
		fmt.Fprintln(os.Stderr, msg)
		return errors.New(msg)
	}
	return nil
}

func main() {
	flag.Usage = usage
	if flag.NArg() != NUM_ARGS {
		usage()
		os.Exit(1)
	}

	if err := run(bufio.NewScanner(os.Stdin), os.Stdout, false); err != nil {
		os.Exit(1)
	}
}