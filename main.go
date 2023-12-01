package main

import (
	"flag"

	y22d5 "github.com/ichi0915/advent-of-code/2022/day5"
	y23d1 "github.com/ichi0915/advent-of-code/2023/day1"
)

func main() {
	var programToRun int
	var yearToRun int

	// flags declaration using flag package
	flag.IntVar(&programToRun, "p", 1, "Specify program to run. Default is 1")
	flag.IntVar(&yearToRun, "y", 2023, "Specify year to run. Default is 2023")

	flag.Parse()

	// check if cli params match
	if yearToRun == 2023 {
		if programToRun == 1 {
			y23d1.Main()
		} else {
			y23d1.Main()
		}
	} else if yearToRun == 2022 {
		if programToRun == 5 {
			y22d5.Main()
		} else {
			y22d5.Main()
		}
	}
}
