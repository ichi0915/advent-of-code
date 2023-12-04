/*
--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source,
but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone!
The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one.
If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine.
There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally,
is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right).
Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?

*/

package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Numbers struct {
	value   int
	yPos    int
	xStart  int
	xFinish int
	len     int
}

type Symbols struct {
	symbol string
	yPos   int
	xPos   int
}

type Coordinates struct {
	yPos int
	xPos int
}

var debug bool
var trace bool
var columnAmount int
var rowAmount int

func Main() {
	trace = false
	debug = false
	debug = true
	fmt.Println("== Day 3 of 2023 ==")

	readData := readFileData2D()
	if debug {
		fmt.Println("\nPrinting Read Data")
		print2DArrayFull(readData)
	}

	numbers, symbols := getNumbersAndSymbolsPos(readData)
	if debug {
		fmt.Println("\nPrinting Numbers")
		printNumbersArray(numbers)
		fmt.Println("\nPrinting Symbols")
		printSymbolsArray(symbols)
		columnAmount--
		fmt.Println("\nColumn Amount: ", columnAmount, "\n")
		rowAmount--
		fmt.Println("\nRow Amount: ", rowAmount, "\n")
	}

	partNumbers := getAdjacentPartNumbers(numbers, symbols)
	if debug {
		fmt.Println("Printing Part Numbers")
		printArrayInt(partNumbers)
	}

	partNumbersSum := sumValues(partNumbers)
	fmt.Println("partNumbersSum: ", partNumbersSum)

}

func getNumbersAndSymbolsPos(readData [][]string) ([]Numbers, []Symbols) {
	var responseNumbers []Numbers
	var responseSymbols []Symbols
	var globalJ int
	notNumberFlag := false
	startingPos := 0

	for i, line := range readData {
		currentNumber := ""
		notNumberFlag = true
		for j, val := range line {
			if "." != val {
				if isInt(val) {
					currentNumber += val
					if notNumberFlag {
						startingPos = j
					}
					notNumberFlag = false
				} else { // Is a symbol
					// pos := fmt.Sprint(i, ",", j)
					symbol := Symbols{val, i, j}
					responseSymbols = append(responseSymbols, symbol)
					notNumberFlag = true
				}
			} else {
				notNumberFlag = true
			}

			if notNumberFlag {
				if "" != currentNumber {
					number := Numbers{value: stringToIntIgnoreError(currentNumber), yPos: i, xStart: startingPos, xFinish: j - 1, len: (j) - startingPos}
					currentNumber = ""

					responseNumbers = append(responseNumbers, number)
				}
			}
			// fmt.Println("val:", val)
			globalJ = j + 1
		}

		if !notNumberFlag && "" != currentNumber {
			number := Numbers{value: stringToIntIgnoreError(currentNumber), yPos: i, xStart: startingPos, xFinish: globalJ - 1, len: (globalJ) - startingPos}
			currentNumber = ""

			responseNumbers = append(responseNumbers, number)
		}
	}

	return responseNumbers, responseSymbols
}

func getAdjacentPartNumbers(numbers []Numbers, symbols []Symbols) []int {
	var response []int

	for _, numberVal := range numbers {

		coordinates := getNumberCoordinates(numberVal)
		if trace {
			printNumbers(numberVal)
			printCoordinatesArray(coordinates)
		}

		for _, symbolVal := range symbols {
			// isAdjacent := checkIfAdjacent(numberVal, symbolVal)
			isAdjacent := checkIfAdjacent(coordinates, symbolVal)
			if isAdjacent {
				if debug {
					fmt.Println("isAdjacent", isAdjacent)
				}
				response = append(response, numberVal.value)
				break
			}
		}

		// fmt.Println("value:", value, " yPos:", yPos, " xPosArr:", xPosArr)
	}

	return response
}

// Return all the posible adjacent coordinates for the number
func getNumberCoordinates(number Numbers) []Coordinates {
	var coordinates []Coordinates
	var coordinate Coordinates

	// Los valores de arriba
	if 0 != number.yPos {
		if 0 != number.xStart {
			coordinate = Coordinates{yPos: number.yPos - 1, xPos: number.xStart - 1}
			coordinates = append(coordinates, coordinate)
		}

		for i := number.xStart + 1; i < (number.xStart + number.len + 1); i++ {
			coordinate = Coordinates{yPos: number.yPos - 1, xPos: i - 1}
			coordinates = append(coordinates, coordinate)
		}
		if columnAmount >= (number.xFinish + 1) {
			coordinate = Coordinates{yPos: number.yPos - 1, xPos: number.xFinish + 1}
			coordinates = append(coordinates, coordinate)
		}

	}

	// Los valores de en medio
	if 0 < number.xStart {
		coordinate = Coordinates{yPos: number.yPos, xPos: number.xStart - 1}
		coordinates = append(coordinates, coordinate)
	}
	if columnAmount >= (number.xFinish + 1) {
		coordinate = Coordinates{yPos: number.yPos, xPos: number.xFinish + 1}
		coordinates = append(coordinates, coordinate)
	}

	// Los valores de abajo
	if rowAmount > number.yPos {
		if 0 != number.xStart {
			coordinate = Coordinates{yPos: number.yPos + 1, xPos: number.xStart - 1}
			coordinates = append(coordinates, coordinate)
		}

		for i := number.xStart + 1; i < (number.xStart + number.len + 1); i++ {
			coordinate = Coordinates{yPos: number.yPos + 1, xPos: i - 1}
			coordinates = append(coordinates, coordinate)
		}
		if columnAmount >= (number.xFinish + 1) {
			coordinate = Coordinates{yPos: number.yPos + 1, xPos: number.xFinish + 1}
			coordinates = append(coordinates, coordinate)
		}
	}

	return coordinates
}

func checkIfAdjacent(coordinates []Coordinates, symbol Symbols) bool {
	response := false

	for _, coordinate := range coordinates {
		if coordinate.xPos == symbol.xPos &&
			coordinate.yPos == symbol.yPos {
			response = true
			break
		}
	}

	return response
}

// Returns every digit in a different positions of a 2d array
func readFileData2D() [][]string {
	var response [][]string
	var file *os.File
	var err error

	if debug {
		file, err = os.Open("./2023/day3/puzzleInputSimple.txt")
		// file, err = os.Open("./2023/day3/puzzleInputSimplePart2.txt")
	} else {
		file, err = os.Open("./2023/day3/puzzleInput.txt")
	}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readValue := scanner.Text()
		rowAmount++

		if "" != readValue {
			val := strings.Split(readValue, "")
			response = append(response, val)
			columnAmount = len(readValue)
		}
	}

	return response
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func sumValues(arr []int) int {
	response := 0

	for _, val := range arr {
		response += val
	}

	return response
}

func printArray(arr []string) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func printArrayInt(arr []int) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func print2DArray(arr [][]string) {
	fmt.Println("== print2DArray ==")
	for i := range arr {
		fmt.Println(arr[i][0], " ", arr[i][1], " ", arr[i][2])
	}
}

func print2DArrayInt(arr [][]int) {
	fmt.Println("== print2DArrayInt ==")
	for i := range arr {
		fmt.Println(arr[i][0], " ", arr[i][1], " ", arr[i][2])
	}
}

func print2DArrayFull(arr [][]string) {
	fmt.Println("== print2DArrayFull ==")

	for i := range arr {
		for j := range arr[i] {
			if "" != arr[i][j] {
				fmt.Print(arr[i][j])
			}
		}
		fmt.Println()
	}
}

func print2DArrayFullV2(arr [][]string) {
	fmt.Println("== print2DArrayFullV2 ==")

	fmt.Println("y-x,x value")
	for i := range arr {
		for j := range arr[i] {
			if "" != arr[i][j] {
				fmt.Print(arr[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func printNumbers(number Numbers) {
	fmt.Println(number.yPos, "-", number.xStart, ",", number.xFinish, " ", number.value, "  ", number.len)
}

func printNumbersArray(numbers []Numbers) {
	fmt.Println("== printNumbersArray ==")

	fmt.Println("y - x , x  value    len")
	for i := range numbers {
		fmt.Println(numbers[i].yPos, "-", numbers[i].xStart, ",", numbers[i].xFinish, " ", numbers[i].value, "  ", numbers[i].len)
	}
}

func printSymbols(symbol Symbols) {
	fmt.Println(symbol.yPos, "-", symbol.xPos, " ", symbol.symbol)
}

func printSymbolsArray(symbols []Symbols) {
	fmt.Println("== printSymbolsArray ==")

	fmt.Println("y - x  symbol")
	for i := range symbols {
		fmt.Println(symbols[i].yPos, "-", symbols[i].xPos, " ", symbols[i].symbol)
	}
}

func printCoordinatesArray(coordinates []Coordinates) {
	fmt.Println("== printCoordinatesArray ==")

	fmt.Println("y - x")
	for i := range coordinates {
		fmt.Println(coordinates[i].yPos, "-", coordinates[i].xPos)
	}
}

func stringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func stringToIntIgnoreError(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}
