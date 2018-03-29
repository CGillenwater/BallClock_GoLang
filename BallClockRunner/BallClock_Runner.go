package driver

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	//"path"
	"time"
	"github.com/CGillenwater/BallClock_GoLang/Clock"
	"github.com/CGillenwater/BallClock_GoLang/Parser"
)

const NUM_ARGS = 0
const MAX_BALLS = 127
const MIN_BALLS = 27
const END_OF_INPUT_VAL = 0
const TIME_BEFORE_SCREEN_CLEAR = 2

//Clears Screen
func clearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func MainMenuLogic(scanner *bufio.Scanner) {
	var selection uint64
	var err error
	duration := time.Duration(TIME_BEFORE_SCREEN_CLEAR)*time.Second
	
	for true{
		time.Sleep(duration)
		clearScreen()
		selection, err = parser.MainMenuPrompt(scanner)
		switch selection {
		case 1:
			clearScreen()
			RunSingleInput(scanner)
		case 2:
			clearScreen()
			fmt.Println("This is the second choice.\n Returning to Main Menu...")
		case 0:
			clearScreen()
			fmt.Println("\n-----EXITING-----")
			os.Exit(1)
		default:
			fmt.Println(err)
			parser.MainMenuPrompt(scanner)
		}
	}
}

//Parse the scanned input from bufio. Otherwise, throw an error
func RunSingleInput(scanner *bufio.Scanner) error {
	var numBalls uint64 //uint64 instead of uint8 due to strconv.ParseUint returning a uint64
	var err error

	numBalls, err = parser.PromptForSingleInput(scanner, MIN_BALLS, MAX_BALLS)
	if err != nil {
		return err
	} 
	fmt.Fprintf(os.Stderr, "%d balls cycle after %d days.\n Returning to Main Menu... \n", numBalls, clock.CalcNumDaysInCycle(uint8(numBalls)))
	return nil
}