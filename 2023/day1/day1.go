/*
--- Day 1: Trebuchet?! ---

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky")
and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course,
where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet
("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended
by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading
the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that
the Elves now need to recover. On each line, the calibration value can be found by combining the first digit
and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

--- Part Two ---
Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values?

*/

// package main
package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var debug bool

func Main() {
	debug = false
	// debug = true
	fmt.Println("== Day 1 of 2023 ==")

	calibrationData := readFileData()
	if debug {
		fmt.Println("Printing calibration Data")
		printArray(calibrationData)
	}

	// digits := getDigitsPart1(calibrationData)
	digits := getDigitsPart2(calibrationData)
	if debug {
		fmt.Println("Printing Digits")
		printArray(digits)
	}

	sum := getSum(digits)
	fmt.Println("The sum is: ", sum)
}

func getDigitsPart1(calibrationData []string) []string {
	var response []string
	pos := 0
	regDigits := regexp.MustCompile(`(\d).*(\d)|(\d)`)

	for _, val := range calibrationData {
		// val.fin
		matches := regDigits.FindAllStringSubmatch(val, -1)
		var tmpVal string
		for i := range matches {
			for j := range matches[i] {
				if 1 == len(matches[i][j]) {
					if debug {
						fmt.Println("Matches i:", i, " j:", j)
						fmt.Println(matches[i][j])
					}
					tmpVal += matches[i][j]
				}
			}
			response = append(response, tmpVal)
		}
		pos++
	}

	return response
}

func getDigitsPart2(calibrationData []string) []string {
	var response []string
	pos := 0
	regDigits := regexp.MustCompile(`(?:(?:(\d)|(\Qone\E)|(\Qtwo\E)|(\Qthree\E)|(\Qfour\E)|(\Qfive\E)|(\Qsix\E)|(\Qseven\E)|(\Qeight\E)|(\Qnine\E)).*(?:(\d)|(\Qone\E)|(\Qtwo\E)|(\Qthree\E)|(\Qfour\E)|(\Qfive\E)|(\Qsix\E)|(\Qseven\E)|(\Qeight\E)|(\Qnine\E)))|(?:(?:(\d))|(?:(\Qone\E)|(\Qtwo\E)|(\Qthree\E)|(\Qfour\E)|(\Qfive\E)|(\Qsix\E)|(\Qseven\E)|(\Qeight\E)|(\Qnine\E)))`)

	for _, val := range calibrationData {
		matches := regDigits.FindAllStringSubmatch(val, -1)
		var val string

		if debug {
			print2DArrayFull(matches)
		}

		for i := range matches {
			tmpVal := "0"
			for j := range matches[i] {
				// if debug {
				// 	fmt.Println("\nMatches i:", i, " j:", j)
				// 	fmt.Println(matches[i][j])
				// }

				if 1 == len(matches[i][j]) {
					tmpVal = matches[i][j]
				} else {
					tmpVal = getStrFromStr(matches[i][j])
				}

				// fmt.Println("tmpVal: ", tmpVal)

				if "0" != tmpVal {
					val += tmpVal
				}
				// fmt.Println("tmpVal: ", val)
			}
			response = append(response, val)
		}
		pos++
	}

	return response
}

func getStrFromStr(str string) string {
	response := "0"

	switch str {
	case "one":
		response = "1"
	case "two":
		response = "2"
	case "three":
		response = "3"
	case "four":
		response = "4"
	case "five":
		response = "5"
	case "six":
		response = "6"
	case "seven":
		response = "7"
	case "eight":
		response = "8"
	case "nine":
		response = "9"
	}

	return response
}

func getIntFromStr(str string) int {
	response := 0

	switch str {
	case "one":
		response = 1
	case "two":
		response = 2
	case "three":
		response = 3
	case "four":
		response = 4
	case "five":
		response = 5
	case "six":
		response = 6
	case "seven":
		response = 7
	case "eight":
		response = 8
	case "nine":
		response = 9
	}

	return response
}

func getSum(digits []string) int {
	response := 0

	for _, val := range digits {
		responseTmp, _ := stringToInt(val)
		response += responseTmp
	}

	return response
}

func readFileData() []string {
	var response []string
	var file *os.File
	var err error

	if debug {
		// file, err = os.Open("./2023/day1/puzzleInputSimple.txt")
		file, err = os.Open("./2023/day1/puzzleInputSimplePart2.txt")
	} else {
		file, err = os.Open("./2023/day1/puzzleInput.txt")
	}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readValue := scanner.Text()

		if "" != readValue {
			response = append(response, readValue)
		}
	}

	return response
}

func printArray(arr []string) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

// func print2DArray(arr [][]string) {
// 	fmt.Println("== print2DArray ==")
// 	for i := range arr {
// 		fmt.Println(arr[i][0], " ", arr[i][1])
// 	}
// }

func print2DArrayFull(arr [][]string) {
	fmt.Println("== print2DArrayFull ==")

	for i := range arr {
		for j := range arr[i] {
			if "" != arr[i][j] {
				fmt.Println(arr[i][j])
			}
		}
	}
}

func stringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
