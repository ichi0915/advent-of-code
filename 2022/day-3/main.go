/*
--- Day 3: Rucksack Reorganization ---

One Elf has the important job of loading all of the rucksacks with supplies for the jungle journey. Unfortunately, that Elf didn't quite follow the packing instructions, and so a few items now need to be rearranged.

Each rucksack has two large compartments. All items of a given type are meant to go into exactly one of the two compartments. The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.

The Elves have made a list of all of the items currently in each rucksack (your puzzle input), but they need your help finding the errors. Every item type is identified by a single lowercase or uppercase letter (that is, a and A refer to different types of items).

The list of items for each rucksack is given as characters all on a single line. A given rucksack always has the same number of items in each of its two compartments, so the first half of the characters represent items in the first compartment, while the second half of the characters represent items in the second compartment.

For example, suppose you have the following list of contents from six rucksacks:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw

    The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp. The only item type that appears in both compartments is lowercase p.
    The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item type that appears in both compartments is uppercase L.
    The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is uppercase P.
    The fourth rucksack's compartments only share item type v.
    The fifth rucksack's compartments only share item type t.
    The sixth rucksack's compartments only share item type s.

To help prioritize item rearrangement, every item type can be converted to a priority:

    Lowercase item types a through z have priorities 1 through 26.
    Uppercase item types A through Z have priorities 27 through 52.

In the above example, the priority of the item type that appears in both compartments of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.

Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?


-- Data ichi --
	a through z have priorities 1 through 26
	A through Z have priorities 27 through 52.


--- Part Two ---

-- Data ichi --

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	rucksackCompartments := readFileData()
	// print2DArray(rucksackCompartments)

	compartmentsItem := getItemsInBothCompartments(rucksackCompartments)
	// printArray(compartmentsItem)

	sum := getSumOfPriorities(compartmentsItem)
	fmt.Println("Sum of the priorities: ", sum)
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
		len := len(scanner.Text()) / 2
		readValue := scanner.Text()

		strategyValues = []string{readValue[:len], readValue[len:]}
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

func printArray(arr []string) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func getItemsInBothCompartments(arr [][]string) []string {
	var items []string

	for i := range arr {
		char := compareChars(arr, i)
		items = append(items, char)
	}
	return items
}

func compareChars(arr [][]string, i int) string {
	resp := "0"

	for _, char := range arr[i][0] {
		for _, char2 := range arr[i][1] {
			if equals(string(char), string(char2)) {
				resp = string(char)
				break
			}
		}
	}

	return resp
}

func equals(str string, str2 string) bool {
	response := false
	if str == str2 {
		response = true
	}
	return response
}

// a through z have priorities 1 through 26
// A through Z have priorities 27 through 52.
// 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s);
func getSumOfPriorities(priorities []string) int {
	sum := 0

	for _, val := range priorities {
		amount := charToIntVal([]rune(val))
		// fmt.Println("Priority: ", amount)
		sum += amount
	}

	return sum
}

func charToIntVal(r []rune) int {
	resp := 0
	//Here we check if the letter is upper case or not to rest the Ascii value of the A or a
	if unicode.IsUpper(r[0]) {
		resp = int(r[0] - 64 + 26) // rune - (Ascii of A-1) - 26 minusculas --- A = 65
	} else {
		resp = int(r[0] - 96) // rune - (Ascii of a-1) --- a = 97
	}

	return resp
}
