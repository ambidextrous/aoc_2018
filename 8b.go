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

func main() {
	filename := "test"
	//filename = "input8a"
	input := readFile(filename)
	input = strings.Replace(input, "\n", "", -1)
	fmt.Println("input:")
	fmt.Println(input)
	//input = "0 1 99"
	//input = "1 1 0 1 99 2"
	intSlice := convertInputToIntSlice(input)
	fmt.Println("intSlice:")
	fmt.Println(intSlice)

	//tree, err := parseTree(intSlice)
	//check(err)
	//fmt.Println("tree:")
	//fmt.Printf("%+v\n", tree)
	//metadataTotal := countTreeMetadata(tree)
	//fmt.Println("metadataTotal:")
	//fmt.Println(metadataTotal)

	total, value, remaining := parseTreeNew(intSlice)
	fmt.Println("total:")
	fmt.Println(total)
	fmt.Println("value:")
	fmt.Println(value)
	fmt.Println("remaining:")
	fmt.Println(remaining)
	//lineSlice := splitStringByNewlines(input)
	//fmt.Println("lineSlice:")
	//fmt.Println(lineSlice)
	//alphabet := getAlphabet()
	//dependencyMap := getDependencyMap(alphabet, lineSlice)
	//fmt.Println("dependencyMap:")
	//for _, letter := range alphabet {
	//	fmt.Printf("%s requires ", letter)
	//	fmt.Println(dependencyMap[letter])

	//}
	//orderOfTasks := completeTasks(dependencyMap)
	//fmt.Println("Answer to part 1: orderOfTasks:")
	//fmt.Println(orderOfTasks)
	//numSecondsToComplete := getNumSecondsToComplete(dependencyMap, 5, 60)
	//fmt.Println("Answer to part 2: numSecondsToComplete:")
	//fmt.Println(numSecondsToComplete)
}
