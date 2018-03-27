package driver

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"github.com/CGillenwater/BallClock_GoLang/Clock"
	"github.com/CGillenwater/BallClock_GoLang/Parser"
)

const NUM_ARGS = 0
const MAX_BALLS = 127
const MIN_BALLS = 27
const END_OF_INPUT_VAL = 0

func Usage() {
	name := path.Base(os.Args[0])
	msg := fmt.Sprintf("Usage: %s\n\n"+"%s takes no arguments and accepts input from stdin.\n", name, name)
	fmt.Fprintf(os.Stderr, msg)
}

//Parse the scanned input from bufio. Otherwise, throw an error
func RunSingleInput(scanner *bufio.Scanner) error {
	var numBalls uint64 //uint64 instead of uint8 due to strconv.ParseUint returning a uint64
	var err error
	numBalls, err = parser.PromptForSingleInput(scanner, MIN_BALLS, MAX_BALLS)
	if(err != nil) {
		return err
	} 
	fmt.Fprintf(os.Stderr, "%d balls cycle after %d days. \n", numBalls, clock.CalcNumDaysInCycle(uint8(numBalls)))
	return nil
}