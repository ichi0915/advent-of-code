/*
--- Day 5: Supply Stacks ---

The expedition can depart as soon as the final supplies have been unloaded from the ships. Supplies are stored in stacks of marked crates,
but because the needed supplies are buried under many other crates, the crates need to be rearranged.

The ship has a giant cargo crane capable of moving crates between stacks. To ensure none of the crates get crushed or fall over,
the crane operator will rearrange them in a series of carefully-planned steps. After the crates are rearranged, the desired crates will be at the top of each stack.

The Elves don't want to interrupt the crane operator during this delicate procedure, but they forgot to ask her which crate will end up where,
and they want to be ready to unload them as soon as possible so they can embark.

They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2

In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates;
from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.

Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack.
In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1, resulting in this configuration:

[D]
[N] [C]
[Z] [M] [P]
 1   2   3

In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time, so the first crate to be moved (D) ends up below the second and third crates:

        [Z]
        [N]
    [C] [D]
    [M] [P]
 1   2   3

Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

        [Z]
        [N]
[M]     [D]
[C]     [P]
 1   2   3

Finally, one crate is moved from stack 1 to stack 2:

        [Z]
        [N]
        [D]
[C] [M] [P]
 1   2   3

The Elves just need to know which crate will end up on top of each stack; in this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.

After the rearrangement procedure completes, what crate ends up on top of each stack?


-- Data ichi --
movementsList[0]	- move Amount
movementsList[1]	- From
movementsList[2]	- To

--- Part Two ---

-- Data ichi --

*/

// package main
package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Main() {
	stackList, movementsList := readFileData()

	fmt.Println("stackList")
	// print2DArray(stackList)

	fmt.Println("movementsList")
	// print2DArray(movementsList)

	fmt.Println("createInvertedStacks")
	// stackObject := createStack(stackList)
	stackListObject := createInvertedStacks(stackList)
	// popStackList(stackListObject)

	stackListObject = rearrangement(stackListObject, movementsList)
	popStackList(stackListObject)

	// cratesOntop := getCrateOnTop(stackListObject)
	// fmt.Println("Creates on top: ", cratesOntop)
}

func readFileData() ([][]string, [][]string) {
	var stackList [][]string
	var movementsList [][]string
	list := make([]string, 0)
	moveSection := false

	file, err := os.Open("./2022/day5/puzzleInput.txt")
	// file, err := os.Open("./2022/day5/puzzleInputSimple.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//((   )|(\[[a-zA-Z]\]))( ((   )|(\[[a-zA-Z]\])) )|(\[[a-zA-Z]\])
	// reg := regexp.MustCompile(`((   )|(\[[a-zA-Z]\]))( ((   )|(\[[a-zA-Z]\])) )|(\[[a-zA-Z]\])`)
	regList := regexp.MustCompile(`(\[[a-zA-Z]\])|(    )`)
	regMove := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readValue := scanner.Text()

		if "" == readValue {
			moveSection = true
		}

		if moveSection {
			// readValue := scanner.Text()
			// fmt.Println(readValue)
			matches := regMove.FindAllStringSubmatch(readValue, -1)
			for i := range matches {
				for j := range matches[i] {
					if j != 0 {
						// fmt.Println(".", matches[i][j], ".")
						list = append(list, matches[i][j])
					}
				}
			}
			if nil != list && 0 < len(list) {
				movementsList = append(movementsList, list)
			}
			list = []string{}
		} else {

			// matches := reg.FindAllStringSubmatch(readValue, -1)
			matches := regList.FindAllString(readValue, -1)

			// fmt.Println(matches)
			for _, match := range matches {
				match = strings.Trim(match, " ")
				match = strings.ReplaceAll(match, "[", "")
				match = strings.ReplaceAll(match, "]", "")
				// fmt.Println("", match, "")
				// fmt.Println(match, "-")
				list = append(list, match)
			}
			stackList = append(stackList, list)
			list = []string{}

			if "" == readValue {
				moveSection = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return stackList, movementsList
}

func print2DArray(arr [][]string) {
	for i := range arr {
		for j := range arr[i] {
			// fmt.Print(arr[i][j], ".")
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

func createInvertedStacks(arr [][]string) []*Stack {
	amountStacks := len(arr)
	stackList := createEmptyStackList(amountStacks) //Create an empty List of Stacks

	//here we iterate the array in reverse
	for i := amountStacks - 1; i >= 0; i-- {
		k := amountStacks - 1
		for j := len(arr[i]) - 1; j >= 0; j-- {
			if "" != arr[i][j] {
				stackList[k].Push(arr[i][j])
				// fmt.Print(arr[i][j], " ")
				// fmt.Print(stackList[k].top.value, " ") //P M Z C N D
			}
			k--
		}
	}

	return stackList
}

func createEmptyStackList(amount int) []*Stack {
	stacks := []*Stack{}

	for i := 0; i <= amount; i++ {
		stack := new(Stack)
		stacks = append(stacks, stack)
	}

	return stacks
}

// This function is to check if the stack is correct
func popStackList(stackList []*Stack) {
	fmt.Println("Popping Stack List")

	for _, val := range stackList {
		stackLen := val.Len()
		for i := 0; i < stackLen; i++ {
			fmt.Print(val.Pop(), " ")
		}
		fmt.Println()
	}
}

/*
movementsList[0]	- move Amount
movementsList[1]	- From
movementsList[2]	- To
*/
func rearrangement(stackList []*Stack, movementsList [][]string) []*Stack {
	for i := range movementsList {
		moveAmount, _ := strconv.Atoi(movementsList[i][0])
		moveFrom, _ := strconv.Atoi(movementsList[i][1])
		moveTo, _ := strconv.Atoi(movementsList[i][2])
		fmt.Println(moveAmount, " ", moveFrom, " ", moveTo)

		// crates_to_move := stackList[moveFrom-1].shift(moveAmount)
		// stackList[moveTo-1] = crates_to_move.reverse.concat(stackList[moveTo-1])
		for j := 0; j < moveAmount; j++ {
			stackList[moveTo].Push(stackList[moveFrom].Pop())
		}
	}

	return stackList
}

// Error - NCSZCJGV - NCVSZFCJGV
func getCrateOnTop(stackList []*Stack) string {
	response := ""
	for _, val := range stackList {
		if nil != val.top && nil != val.top.value {
			response += val.top.value.(string)
			// response += val.Pop().(string)
		}
	}
	return response
}
