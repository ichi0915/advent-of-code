/*
--- Day 4: Camp Cleanup ---

Space needs to be cleared before the last supplies can be unloaded from the ships, and so several Elves have been assigned the job of cleaning up sections of the camp.
Every section has a unique ID number, and each Elf is assigned a range of section IDs.

However, as some of the Elves compare their section assignments with each other, they've noticed that many of the assignments overlap.
To try to quickly find overlaps and reduce duplicated effort, the Elves pair up and make a big list of the section assignments for each pair (your puzzle input).

For example, consider the following list of section assignment pairs:

2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8

For the first few pairs, this list means:

    Within the first pair of Elves, the first Elf was assigned sections 2-4 (sections 2, 3, and 4), while the second Elf was assigned sections 6-8 (sections 6, 7, 8).
    The Elves in the second pair were each assigned two sections.
    The Elves in the third pair were each assigned three sections: one got sections 5, 6, and 7, while the other also got 7, plus 8 and 9.

This example list uses single-digit section IDs to make it easier to draw; your actual list might contain larger numbers. Visually, these pairs of section assignments look like this:

.234.....  2-4
.....678.  6-8

.23......  2-3
...45....  4-5

....567..  5-7
......789  7-9

.2345678.  2-8
..34567..  3-7

.....6...  6-6
...456...  4-6

.23456...  2-6
...45678.  4-8

Some of the pairs have noticed that one of their assignments fully contains the other. For example, 2-8 fully contains 3-7, and 6-6 is fully contained by 4-6. In pairs where one assignment fully contains the other, one Elf in the pair would be exclusively cleaning sections their partner will already be cleaning, so these seem like the most in need of reconsideration. In this example, there are 2 such pairs.

In how many assignment pairs does one range fully contain the other?


-- Data ichi --


--- Part Two ---

It seems like there is still quite a bit of duplicate work planned. Instead, the Elves would like to know the number of pairs that overlap at all.

In the above example, the first two pairs (2-4,6-8 and 2-3,4-5) don't overlap, while the remaining four pairs (5-7,7-9, 2-8,3-7, 6-6,4-6, and 2-6,4-8) do overlap:

    5-7,7-9 overlaps in a single section, 7.
    2-8,3-7 overlaps all of the sections 3 through 7.
    6-6,4-6 overlaps in a single section, 6.
    2-6,4-8 overlaps in sections 4, 5, and 6.

So, in this example, the number of overlapping assignment pairs is 4.

In how many assignment pairs do the ranges overlap?


-- Data ichi --

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	cleaningSection := readFileData(1)
	// print2DArray(cleaningSection)

	sum := getSumOfAssignmentPairsThatFullyContainTheOtherWrap(cleaningSection)
	fmt.Println("Sum of the contains: ", sum)

	sum = getSumOfAssignmentPairsThatOverlapTheOtherWrap(cleaningSection)
	fmt.Println("Sum of the overlaps: ", sum)

	/*cleaningSectionValues := unwrapCleaningSections(cleaningSection)
	// print2DArray(cleaningSectionValues)

	sum = getSumOfAssignmentPairsThatFullyContainTheOther(cleaningSectionValues)
	fmt.Println("Sum of the contains: ", sum)*/
}

func readFileData(typeOfArray int) [][]string {
	var response [][]string
	pairOfElfs := make([]string, 0)

	file, err := os.Open("puzzleInput.txt")
	// file, err := os.Open("puzzleInputSimple.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readValue := strings.Split(scanner.Text(), ",")

		pairOfElfs = []string{readValue[0], readValue[1]}
		response = append(response, pairOfElfs)
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

func getSumOfAssignmentPairsThatFullyContainTheOtherWrap(arr [][]string) int {
	sum := 0

	for i := range arr {
		leftBottom, _ := strconv.Atoi(strings.Split(arr[i][0], "-")[0])
		leftTop, _ := strconv.Atoi(strings.Split(arr[i][0], "-")[1])
		rightBottom, _ := strconv.Atoi(strings.Split(arr[i][1], "-")[0])
		rightTop, _ := strconv.Atoi(strings.Split(arr[i][1], "-")[1])

		if leftBottom <= rightBottom && leftTop >= rightTop {
			sum++
		} else if rightBottom <= leftBottom && rightTop >= leftTop {
			sum++
		}
	}

	return sum
}

func getSumOfAssignmentPairsThatOverlapTheOtherWrap(arr [][]string) int {
	sum := 0

	for i := range arr {
		leftBottom, _ := strconv.Atoi(strings.Split(arr[i][0], "-")[0])
		leftTop, _ := strconv.Atoi(strings.Split(arr[i][0], "-")[1])
		rightBottom, _ := strconv.Atoi(strings.Split(arr[i][1], "-")[0])
		rightTop, _ := strconv.Atoi(strings.Split(arr[i][1], "-")[1])

		if leftBottom <= rightBottom && leftTop >= rightTop {
			sum++
		} else if rightBottom <= leftBottom && rightTop >= leftTop {
			sum++
		} else if leftTop >= rightBottom && rightTop >= leftBottom {
			sum++
		}
	}

	return sum
}

/*func unwrapCleaningSections(arr [][]string) [][]string {
	var response [][]string
	pairOfElfs := make([]string, 0)

	for i := range arr {
		pairOfElfs = []string{unwrapValues(arr[i][0]), unwrapValues(arr[i][1])}
		response = append(response, pairOfElfs)
	}

	return response
}

func unwrapValues(val string) string {
	response := ""
	iter := strings.Split(val, "-")
	bottom, _ := strconv.Atoi(iter[0])
	top, _ := strconv.Atoi(iter[1])
	// fmt.Println("bottom:", bottom, " top:", top)

	for bottom <= top {
		response += strconv.Itoa(bottom) + ","
		bottom++
	}

	return response
}

func getSumOfAssignmentPairsThatFullyContainTheOther(arr [][]string) int {
	sum := 0

	for i := range arr {
		if strings.Contains(arr[i][0], arr[i][1]) || strings.Contains(arr[i][1], arr[i][0]) {
			// fmt.Println(arr[i][0], " ", arr[i][1])
			sum++
		}
	}

	return sum
}*/
