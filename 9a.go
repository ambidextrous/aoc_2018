package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func containsString(slice []string, s string) bool {
	for _, element := range slice {
		if element == s {
			return true
		}
	}
	return false
}

func readFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	check(err)
	contents := string(dat)
	return contents
}

func convertInputToIntSlice(input string) []int {
	stringSlice := strings.Split(input, " ")
	intSlice := make([]int, 0)
	for _, s := range stringSlice {
		i, err := strconv.Atoi(s)
		check(err)
		intSlice = append(intSlice, i)
	}
	return intSlice
}

func sumSlice(s []int) int {
	tot := 0
	for _, i := range s {
		tot += i
	}
	return tot
}

type tree struct {
	meta     []int
	children []tree
}

func parseTree(intSlice []int) (tree, error) {
	//fmt.Println("parseTree:")
	//fmt.Println("intSlice:")
	fmt.Println(intSlice)
	t := tree{}
	children := make([]tree, 0)
	newErr := errors.New("too short")
	if len(intSlice) < 2 {
		return t, newErr
	}
	numChild := intSlice[0]
	metaLen := intSlice[1]
	if numChild == 0 {
		if metaLen == len(intSlice)-2 {
			m := intSlice[2:]
			t.meta = m
			t.children = children
			return t, nil
		} else {
			t := tree{}
			return t, newErr
		}
	}
	childrenLeftToCount := numChild
	startPointer := 2
	endPointer := startPointer
	for childrenLeftToCount > 0 {
		//fmt.Println("startPointer:")
		//fmt.Println(startPointer)
		//fmt.Println("endPointer:")
		//fmt.Println(endPointer)
		minViableLength := ((childrenLeftToCount - 1) * 2) + metaLen
		if len(intSlice)-endPointer < minViableLength {
			return t, newErr
		}
		//fmt.Println("intSlice[startPointer:endPointer]")
		//fmt.Println(intSlice[startPointer:endPointer])
		childTree, err := parseTree(intSlice[startPointer:endPointer])
		if err != nil {
			//fmt.Println("caught error")
			endPointer++
		} else {
			children = append(children, childTree)
			startPointer = endPointer //+ 1
			endPointer = startPointer
			childrenLeftToCount--
		}
	}
	t.children = children
	meta := intSlice[len(intSlice)-metaLen:]
	t.meta = meta
	return t, nil
}

func countTreeMetadata(t tree) int {
	tot := 0
	if len(t.meta) > 0 {
		tot += sumSlice(t.meta)
	}
	if len(t.children) > 0 {
		for _, child := range t.children {
			tot += countTreeMetadata(child)
		}
	}
	return tot
}

func parseTreeNew(data []int) (int, int, []int) {
	children := data[0]
	metas := data[1]
	data = data[2:]
	scores := make([]int, 0)
	totals := 0

	total := 0
	score := 0
	for i := 0; i < children; i++ {
		total, score, data = parseTreeNew(data)
		totals += total
		scores = append(scores, score)
	}

	totals += sumSlice(data[:metas])

	if children == 0 {
		return totals, sumSlice(data[:metas]), data[metas:]
	} else {
		numsToSum := make([]int, 0)
		for i, _ := range data[:metas] {
			if i > 0 && i <= len(scores) {
				numsToSum = append(numsToSum, scores[i-1])
			}
		}
		return totals, sumSlice(numsToSum), data[metas:]
	}
}

func insertMarbleAtIndex(marble int, index int, circle []int) []int {
	if index == len(circle) {
		circle = append(circle, marble)
		return circle
	}
	middlePart := make([]int, 0)
	middlePart = append(middlePart, marble)
	secondPart := append(middlePart, circle[index:]...)
	firstPart := circle[:index]
	circle = append(firstPart, secondPart...)
	return circle
}

func removeMarbleAtIndex(index int, circle []int) ([]int, int) {
	retVal := circle[index]
	if index == len(circle) {
		circle = circle[:len(circle)-1]
		return circle, retVal
	}
	firstPart := make([]int, 0)
	firstPart = append(firstPart, circle[:index]...)
	secondPart := make([]int, 0)
	secondPart = append(secondPart, circle[index+1:]...)
	circle = append(firstPart, secondPart...)
	return circle, retVal
}

func playMarbleGame(numPlayers int, finalLastMarbleScore int) int {
	circle := make([]int, 0)
	players := make(map[int][]int)
	for i := 1; i <= numPlayers; i++ {
		players[i] = make([]int, 0)
	}
	currentMarbleScore := 0
	roundCounter := 0
	currentMarblePos := 0
	nextMarblePos := 0
	currentPlayer := 0
	for currentMarbleScore < finalLastMarbleScore {
		if roundCounter%23 == 0 && roundCounter != 0 {
			currentPlayer++
			if currentPlayer > numPlayers {
				currentPlayer = 1
			}
			removalPosition := currentMarblePos - 9
			if removalPosition < 0 {
				removalPosition = len(circle) - removalPosition
			}
			nextMarblePos = currentMarblePos - 7
			if nextMarblePos < 0 {
				nextMarblePos = len(circle) - nextMarblePos
			}
			circle, _ = removeMarbleAtIndex(removalPosition, circle)
		} else if roundCounter == 0 {
			circle = insertMarbleAtIndex(roundCounter, currentMarblePos, circle)
			nextMarblePos = 1
		} else if roundCounter == 1 {
			circle = insertMarbleAtIndex(roundCounter, currentMarblePos, circle)
			nextMarblePos = 1
			currentPlayer = 1
		} else if roundCounter == 2 {
			circle = insertMarbleAtIndex(roundCounter, currentMarblePos, circle)
			nextMarblePos = 3
			currentPlayer = 2
		} else {
			currentPlayer++
			if currentPlayer > numPlayers {
				currentPlayer = 1
			}
			circle = insertMarbleAtIndex(roundCounter, currentMarblePos, circle)
			nextMarblePos = currentMarblePos + 2
			if nextMarblePos > len(circle) {
				nextMarblePos = 1
			}
		}
		fmt.Printf("round %d: player %d: %v\n", roundCounter, currentPlayer, circle)
		roundCounter++
		currentMarbleScore++
		currentMarblePos = nextMarblePos
	}
	return currentMarbleScore
}

func main() {
	//filename := "test"
	//filename = "input8a"
	//input := readFile(filename)
	//input = strings.Replace(input, "\n", "", -1)
	//fmt.Println("input:")
	//fmt.Println(input)
	playMarbleGame(9, 26)
}
