/*
--- Day 4: Scratchcards ---
The gondola takes you up. Strangely, though, the ground doesn't seem to be coming with you; you're not climbing a mountain.
As the circle of Snow Island recedes below you, an entire new landmass suddenly appears above you!
The gondola carries you to the surface of the new island and lurches into the station.

As you exit the gondola, the first thing you notice is that the air here is much warmer than it was on Snow Island.
It's also quite humid. Is this where the water source is?

The next thing you notice is an Elf sitting on the floor across the station in what seems to be a pile of colorful square cards.

"Oh! Hello!" The Elf excitedly runs over to you. "How may I be of service?" You ask about water sources.

"I'm not sure; I just operate the gondola lift. That does sound like something we'd have, though - this is Island Island,
after all! I bet the gardener would know. He's on a different island, though - er, the small kind surrounded by water, not the floating kind.
We really need to come up with a better naming scheme. Tell you what: if you can help me with something quick,
I'll let you borrow my boat and you can go visit the gardener. I got all these scratchcards as a gift, but I can't figure out what I've won."

The Elf leads you over to the pile of colorful cards. There, you discover dozens of scratchcards, all with their opaque covering already scratched off.
Picking one up, it looks like each card has two lists of numbers separated by a vertical bar (|): a list of winning numbers and then a list of numbers you have.
You organize the information into a table (your puzzle input).

As far as the Elf has been able to figure out, you have to figure out which of the numbers you have appear in the list of winning numbers.
The first match makes the card worth one point and each match after the first doubles the point value of that card.

For example:

Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
In the above example,
card 1 has five winning numbers (41, 48, 83, 86, and 17) and eight numbers you have (83, 86, 6, 31, 17, 9, 48, and 53).
Of the numbers you have, four of them (48, 83, 17, and 86) are winning numbers! That means card 1 is worth 8 points (1 for the first match,
then doubled three times for each of the three matches after the first).

Card 2 has two winning numbers (32 and 61), so it is worth 2 points.
Card 3 has two winning numbers (1 and 21), so it is worth 2 points.
Card 4 has one winning number (84), so it is worth 1 point.
Card 5 has no winning numbers, so it is worth no points.
Card 6 has no winning numbers, so it is worth no points.
So, in this example, the Elf's pile of scratchcards is worth 13 points.

Take a seat in the large pile of colorful cards. How many points are they worth in total?

*/

package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Card struct {
	cardNumber     int
	winningNumbers []int
	actualNumbers  []int
	matchesAmount  int
	matchesPoints  int
}

var debug bool
var trace bool

func Main() {
	trace = false
	debug = false
	// debug = true
	// trace = true
	fmt.Println("== Day 4 of 2023 ==")

	readData := readFileData()
	if trace {
		fmt.Println("\nPrinting Read Data")
		printArray(readData)
	}

	cardList := getCards(readData)
	if trace {
		fmt.Println("Printing Card List")
		printCards(cardList)
	}

	setMatchesAmount(cardList)
	if trace {
		fmt.Println("Printing Card List Matches Amount")
		printCards(cardList)
	}

	setMatchesPoints(cardList)
	if debug {
		fmt.Println("Printing Card List Matches Points")
		printCards(cardList)
	}

	pointsSum := getPointsSum(cardList)
	fmt.Println("pointsSum: ", pointsSum)

}

func getPointsSum(cards []Card) int {
	response := 0
	for _, card := range cards {
		response += card.matchesPoints
	}
	return response
}

func setMatchesPoints(cards []Card) {
	for i := range cards {
		matchesPoints := 0
		for j := 1; j <= cards[i].matchesAmount; j++ {
			if 0 == matchesPoints {
				matchesPoints = 1
			} else {
				matchesPoints = matchesPoints * 2
			}
		}
		cards[i].matchesPoints = matchesPoints
	}
}

func setMatchesAmount(cards []Card) {
	for i := range cards {
		for _, winningNum := range cards[i].winningNumbers {
			for _, cardNum := range cards[i].actualNumbers {
				if winningNum == cardNum {
					cards[i].matchesAmount++
					break
				}
			}
		}
	}
}

func getCards(readData []string) []Card {
	var response []Card
	regExp := regexp.MustCompile(`\QCard\E +(\d+): +(.*)\| +(.*)`)

	for _, val := range readData {
		matches := regExp.FindAllStringSubmatch(val, -1)

		for i := range matches {
			var card Card
			for j := range matches[i] {
				if trace {
					fmt.Println("Matches [", i, "][", j, "]: ", matches[i][j], " ")
				}

				// We skip the first one because it is the full string
				if j == 1 {
					card.cardNumber = stringToIntIgnoreError(matches[i][j])
				} else if j == 2 || j == 3 {
					var numbers []int
					cleanGame := strings.Split(matches[i][j], " ")
					for _, val := range cleanGame {
						if " " != val && "" != val {
							numbers = append(numbers, stringToIntIgnoreError(val))
						}
					}
					if j == 2 {
						card.winningNumbers = numbers
					}
					if j == 3 {
						card.actualNumbers = numbers
					}
				}
			}
			response = append(response, card)
		}
		if trace {
			fmt.Println("\n== Card Ended ==")
		}
	}
	return response
}

// Returns every digit in a different positions of a 2d array
func readFileData() []string {
	var response []string
	var file *os.File
	var err error

	if debug {
		file, err = os.Open("./2023/day4/puzzleInputSimple.txt")
		// file, err = os.Open("./2023/day4/puzzleInputSimplePart2.txt")
	} else {
		file, err = os.Open("./2023/day4/puzzleInput.txt")
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

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func printArray(arr []string) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func printCards(cards []Card) {
	for _, card := range cards {
		fmt.Println("Number:", card.cardNumber,
			" Winning:", card.winningNumbers,
			" Actual:", card.actualNumbers,
			" Matches Amount:", card.matchesAmount,
			" Matches Points:", card.matchesPoints)
	}
}

func stringToIntIgnoreError(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}
