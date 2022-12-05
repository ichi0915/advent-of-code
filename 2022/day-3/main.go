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

As you finish identifying the misplaced items, the Elves come to you with another issue.

For safety, the Elves are divided into groups of three. Every Elf carries a badge that identifies their group. For efficiency, within each group of three Elves,
the badge is the only item type carried by all three Elves. That is, if a group's badge is item type B, then all three Elves will have item type B somewhere in their rucksack,
and at most two of the Elves will be carrying any other item type.

The problem is that someone forgot to put this year's updated authenticity sticker on the badges.
All of the badges need to be pulled out of the rucksacks so the new authenticity stickers can be attached.

Additionally, nobody wrote down which item type corresponds to each group's badges.
The only way to tell which item type is the right one is by finding the one item type that is common between all three Elves in each group.

Every set of three lines in your list corresponds to a single group, but each group can have a different badge item type.
So, in the above example, the first group's rucksacks are the first three lines:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg

And the second group's rucksacks are the next three lines:

wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw

In the first group, the only item type that appears in all three rucksacks is lowercase r; this must be their badges. In the second group, their badge item type must be Z.

Priorities for these items must still be found to organize the sticker attachment efforts: here, they are 18 (r) for the first group and 52 (Z) for the second group. The sum of these is 70.

Find the item type that corresponds to the badges of each three-Elf group. What is the sum of the priorities of those item types?


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
	rucksackCompartments := readFileData(1)
	// print2DArray(rucksackCompartments)

	compartmentsItem := getItemsInBothCompartments(rucksackCompartments)
	// printArray(compartmentsItem)

	sum := getSumOfPriorities(compartmentsItem)
	fmt.Println("Sum of the priorities: ", sum)

	rucksack3SetsCompartments := readFileData(2)
	// print2DArray2(rucksack3SetsCompartments)

	compartmentsItem = getItemsIn3Compartments(rucksack3SetsCompartments)
	// printArray(compartmentsItem)

	sum = getSumOfPriorities(compartmentsItem)
	fmt.Println("Sum of the 3 priorities: ", sum)
}

func readFileData(typeOfArray int) [][]string {
	var response [][]string
	strategyValues := make([]string, 0)
	counter := 0

	file, err := os.Open("puzzleInput.txt")
	// file, err := os.Open("puzzleInputSimple.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if 1 == typeOfArray {
			len := len(scanner.Text()) / 2
			readValue := scanner.Text()

			strategyValues = []string{readValue[:len], readValue[len:]}
			response = append(response, strategyValues)
		} else {
			if 3 == counter {
				response = append(response, strategyValues)
				strategyValues = make([]string, 0)
				counter = 0
			}
			counter++
			readValue := scanner.Text()
			strategyValues = append(strategyValues, readValue)
		}
	}
	//Verify if we added the last rucksack inventory
	if 0 != counter {
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

func print2DArray2(arr [][]string) {
	for i := range arr {
		fmt.Println(arr[i][0], " ", arr[i][1], " ", arr[i][2])
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

func getItemsIn3Compartments(arr [][]string) []string {
	var items []string

	for i := range arr {
		char := compare3Chars(arr, i)
		items = append(items, char)
	}
	return items
}

func compare3Chars(arr [][]string, i int) string {
	resp := "0"

	for _, char := range arr[i][0] {
		for _, char2 := range arr[i][1] {
			for _, char3 := range arr[i][2] {
				if equals3(string(char), string(char2), string(char3)) {
					resp = string(char)
					break
				}
			}
		}
	}

	return resp
}

func equals3(str string, str2 string, str3 string) bool {
	response := false
	if str == str2 && str2 == str3 {
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
