package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"github.com/CGillenwater/BallClock_GoLang/BallClockRunner"
)

func parseCommandLine() {
	flag.Parse()
}

func main() {
	parseCommandLine()
	fmt.Println("\n--------------------\nENTER '0' TO EXIT")
	fmt.Println("\nPlease enter the number of balls in the clock (27-127):")

	if flag.NArg() != driver.NUM_ARGS {
		driver.Usage()
		os.Exit(1)
	}

	if err := driver.RunSingleInput(bufio.NewScanner(os.Stdin), false); err != nil {
		os.Exit(1)
	}
}