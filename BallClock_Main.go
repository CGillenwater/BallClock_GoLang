package main

import (
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

	if flag.NArg() != driver.NUM_ARGS {
		driver.Usage()
		os.Exit(1)
	}

	if err := driver.RunSingleInput(bufio.NewScanner(os.Stdin)); err != nil {
		os.Exit(1)
	}
}