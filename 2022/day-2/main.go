/*
--- Day 2: Rock Paper Scissors ---

The Elves begin to set up camp on the beach. To decide whose tent gets to be closest to the snack storage, a giant Rock Paper Scissors tournament is already in progress.

Rock Paper Scissors is a game between two players. Each game contains many rounds; in each round, the players each simultaneously choose one of Rock, Paper, or Scissors using a hand shape.
Then, a winner for that round is selected: Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock. If both players choose the same shape, the round instead ends in a draw.

Appreciative of your help yesterday, one Elf gives you an encrypted strategy guide (your puzzle input) that they say will be sure to help you win.
"The first column is what your opponent is going to play: A for Rock, B for Paper, and C for Scissors. The second column--" Suddenly, the Elf is called away to help with someone's tent.

The second column, you reason, must be what you should play in response: X for Rock, Y for Paper, and Z for Scissors. Winning every time would be suspicious,
so the responses must have been carefully chosen.

The winner of the whole tournament is the player with the highest score. Your total score is the sum of your scores for each round.
The score for a single roundis the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round
(0 if you lost, 3 if the round was a draw, and 6 if you won).

Since you can't be sure if the Elf is trying to help you or trick you, you should calculate the score you would get if you were to follow the strategy guide.

For example, suppose you were given the following strategy guide:

A Y
B X
C Z

This strategy guide predicts and recommends the following:

    In the first round, your opponent will choose Rock (A), and you should choose Paper (Y). This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
    In the second round, your opponent will choose Paper (B), and you should choose Rock (X). This ends in a loss for you with a score of 1 (1 + 0).
    The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.

In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).

What would your total score be if everything goes exactly according to your strategy guide?

-- Data ichi --
Column 1:
	A for Rock,
	B for Paper
	C for Scissors.
Column 2:
	X for Rock
	Y for Paper
	Z for Scissors

Scores
	1 for Rock
	2 for Paper
	3 for Scissors
	+
	0 if you lost
	3 if the round was a draw
	6 if you won
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	encryptedStrategy := readFileData()
	// print2DArray(encryptedStrategy)

	score := calculateScore(encryptedStrategy)
	fmt.Println("My score will be: ", score)
}

func readFileData() [][]string {
	var response [][]string
	strategyValues := make([]string, 0)

	file, err := os.Open("puzzleInput.txt")
	// file, err := os.Open("puzzleInputSimple.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readValue := strings.Split(scanner.Text(), " ")

		strategyValues = []string{readValue[0], readValue[1]}
		response = append(response, strategyValues)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return response
}

func print2DArray(arr [][]string) {
	for i := range arr {
		fmt.Println(arr[i][0], " ", arr[i][1])
	}
}

func calculateScore(arr [][]string) int {
	score := 0

	for i := range arr {
		score += getRoundScore(arr[i][0], arr[i][1])
	}
	return score
}

func getRoundScore(opponentChoice string, myChoice string) int {
	score := 0

	score = getShapeScore(myChoice) + getOutcomeScore(opponentChoice, myChoice)

	return score
}

func getShapeScore(myChoice string) int {
	score := 0

	switch {
	case "X" == myChoice:
		score = 1
	case "Y" == myChoice:
		score = 2
	case "Z" == myChoice:
		score = 3
	}

	return score
}

func getOutcomeScore(opponentChoice string, myChoice string) int {
	score := 0

	convertMyChoice(&myChoice)
	// fmt.Println("opponentChoice: ", opponentChoice, " myChoice: ", myChoice)

	/*
		A for Rock,
		B for Paper
		C for Scissors
		--
		0 lost
		3 draw
		6 won
	*/
	switch {
	//Draw
	case opponentChoice == myChoice:
		score = 3
	//I lose
	case opponentChoice == "A" && myChoice == "C" || //He as Rock and I have Scissors
		opponentChoice == "B" && myChoice == "A" || //He as Paper and I have Rock
		opponentChoice == "C" && myChoice == "B": //He as Scissors and I have Paper
		score = 0
	default:
		score = 6
	}

	return score
}

func convertMyChoice(myChoice *string) {
	switch {
	case "X" == *myChoice:
		*myChoice = "A"
	case "Y" == *myChoice:
		*myChoice = "B"
	case "Z" == *myChoice:
		*myChoice = "C"
	}
}
