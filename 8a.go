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

func main() {
	filename := "test"
	//filename = "inputr87a"
	input := readFile(filename)
	input = strings.Replace(input, "\n", "", -1)
	fmt.Println("input:")
	fmt.Println(input)
	//input = "0 1 99"
	//input = "1 1 0 1 99 2"
	intSlice := convertInputToIntSlice(input)
	fmt.Println("intSlice:")
	fmt.Println(intSlice)
	tree, err := parseTree(intSlice)
	check(err)
	fmt.Println("tree:")
	fmt.Printf("%+v\n", tree)
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
