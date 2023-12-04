/*
--- Day 2: Cube Conundrum ---
You're launched high into the atmosphere! The apex of your trajectory just barely reaches the surface of a large island floating in the sky.
You gently land in a fluffy pile of leaves. It's quite cold, but you don't see much snow. An Elf runs over to greet you.

The Elf explains that you've arrived at Snow Island and apologizes for the lack of snow. He'll be happy to explain the situation, but it's a bit of a walk,
so you have some time. They don't get many visitors up here; would you like to play a game in the meantime?

As you walk, the Elf shows you a small bag and some cubes which are either red, green, or blue. Each time you play this game,
he will hide a secret number of cubes of each color in the bag, and your goal is to figure out information about the number of cubes.

To get information, once a bag has been loaded with cubes, the Elf will reach into the bag, grab a handful of random cubes, show them to you,
and then put them back in the bag. He'll do this a few times per game.

You play several games and record the information from each game (your puzzle input). Each game is listed with its ID number (like the 11 in Game 11: ...)
followed by a semicolon-separated list of subsets of cubes that were revealed from the bag (like 3 red, 5 green, 4 blue).

For example, the record of a few games might look like this:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

In game 1, three sets of cubes are revealed from the bag (and then put back again). The first set is 3 blue cubes and 4 red cubes;
the second set is 1 red cube, 2 green cubes, and 6 blue cubes; the third set is only 2 green cubes.

The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration.
However,game 3 would have been impossible because at one point the Elf showed you 20 red cubes at once; similarly,
game 4 would also have been impossible because the Elf showed you 15 blue cubes at once. If you add up the IDs of the games that would have been possible, you get 8.

Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes.
What is the sum of the IDs of those games?

--- Part Two ---
The Elf says they've stopped producing snow because they aren't getting any water! He isn't sure why the water stopped;
however, he can show you how to get to the water source to check it out for yourself. It's just up ahead!

As you continue your walk, the Elf poses a second question: in each game you played, what is the fewest number of cubes of each color
that could have been in the bag to make the game possible?

Again consider the example games from earlier:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

In game 1, the game could have been played with as few as 4 red, 2 green, and 6 blue cubes.
If any color had even one fewer cube, the game would have been impossible.
Game 2 could have been played with a minimum of 1 red, 3 green, and 4 blue cubes.
Game 3 must have been played with at least 20 red, 13 green, and 6 blue cubes.
Game 4 required at least 14 red, 3 green, and 15 blue cubes.
Game 5 needed no fewer than 6 red, 3 green, and 2 blue cubes in the bag.

The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
The power of the minimum set of cubes in game 1 is 48. In games 2-5 it was 12, 1560, 630, and 36, respectively.
Adding up these five powers produces the sum 2286.

For each game, find the minimum set of cubes that must have been present. What is the sum of the power of these sets?
*/

package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var debug bool
var trace bool

/* -- Part 1 --
Bag loaded with:
	12 red cubes
	13 green cubes
	14 blue cubes
*/

func Main() {
	trace = false
	debug = false
	// debug = true
	fmt.Println("== Day 2 of 2023 ==")

	// Red, Green, Blue
	colorsLoaded := []int{12, 13, 14}

	readData := readFileData()
	if debug {
		fmt.Println("Printing Read Data")
		printArray(readData)
	}

	gamesList := getGames(readData)
	if debug {
		fmt.Println("Printing Games List")
		print2DArrayInt(gamesList)
	}

	posibleGames := getPosibleGames(gamesList, colorsLoaded)
	if debug {
		fmt.Println("Printing Posible Games")
		printArrayInt(posibleGames)
	}

	posibleGamesSum := sumValues(posibleGames)
	fmt.Println("posibleGamesSum: ", posibleGamesSum)

	// --- Part Two ---

	multiplicationList := multiplyValues(gamesList)
	if debug {
		fmt.Println("Printing multiplication List")
		printArrayInt(multiplicationList)
	}

	minValuesSum := sumValues(multiplicationList)
	fmt.Println("minValuesSum: ", minValuesSum)

}

func getPosibleGames(gamesList [][]int, colorsLoaded []int) []int {
	var response []int

	for gameId, val := range gamesList {
		// Red, Green, Blue
		if val[0] <= colorsLoaded[0] &&
			val[1] <= colorsLoaded[1] &&
			val[2] <= colorsLoaded[2] {
			response = append(response, gameId+1)
		}
	}

	return response
}

func getGames(readData []string) [][]int {
	var response [][]int
	gamePos := 0
	// regSets := regexp.MustCompile(`(\d+ (?:\Qred\E|\Qgreen\E|\Qblue\E))`)
	regSets := regexp.MustCompile(`(\d+)(?: \Qred\E| \Qgreen\E| \Qblue\E)`)

	for _, val := range readData {
		cleanGame := strings.Split(val, ":")
		gameSets := strings.Split(cleanGame[1], ";")

		// Red, Green, Blue
		currentMaxColors := []int{0, 0, 0}

		for _, valSets := range gameSets {
			matches := regSets.FindAllStringSubmatch(valSets, -1)
			for i := range matches {
				for j := range matches[i] {
					// fmt.Print("Matches [", i, "][", j, "]: ", matches[i][j], " ")

					// getColor
					if strings.Contains(matches[i][j], "red") {
						amountStr := strings.Split(matches[i][j], " ")
						amount, _ := stringToInt(amountStr[0])
						// fmt.Println("-red match-")
						if currentMaxColors[0] < amount {
							// fmt.Println("red amount: ", amount)
							currentMaxColors[0] = amount
						}
						break
					} else if strings.Contains(matches[i][j], "green") {
						amountStr := strings.Split(matches[i][j], " ")
						amount, _ := stringToInt(amountStr[0])
						if currentMaxColors[1] < amount {
							currentMaxColors[1] = amount
						}
						break
					} else if strings.Contains(matches[i][j], "blue") {
						amountStr := strings.Split(matches[i][j], " ")
						amount, _ := stringToInt(amountStr[0])
						if currentMaxColors[2] < amount {
							currentMaxColors[2] = amount
						}
						break
					}
				}
				if trace {
					fmt.Print(" --- ")
				}
			}
			// Acaba un set
			if trace {
				fmt.Println("\n== Set Ended ==")
				fmt.Println("currentMaxColors of Set:", currentMaxColors)
			}
		}

		if debug {
			fmt.Println("\n== Game Ended ==")
			fmt.Println("currentMaxColors:", currentMaxColors)
		}

		// Go lang is funny, first we create a an array with 3 slots and then we append it to the response array
		response = append(response, make([]int, 3))
		response[gamePos] = currentMaxColors
		gamePos++
	}

	return response
}

func sumValues(posibleGames []int) int {
	response := 0

	for _, val := range posibleGames {
		response += val
	}

	return response
}

func multiplyValues(gamesList [][]int) []int {
	var response []int

	for _, val := range gamesList {
		multiplication := val[0] * val[1] * val[2]
		response = append(response, multiplication)
	}

	return response
}

func readFileData() []string {
	var response []string
	var file *os.File
	var err error

	if debug {
		file, err = os.Open("./2023/day2/puzzleInputSimple.txt")
		// file, err = os.Open("./2023/day2/puzzleInputSimplePart2.txt")
	} else {
		file, err = os.Open("./2023/day2/puzzleInput.txt")
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
				fmt.Println(arr[i][j])
			}
		}
	}
}

func stringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
