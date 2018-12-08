package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

func splitStringByNewlines(string_of_strings string) []string {
	string_array := strings.Split(string_of_strings, "\n")
	var no_empties []string
	for _, item := range string_array {
		if len(item) > 0 {
			no_empties = append(no_empties, item)
		}
	}
	return no_empties
}

func getAlphabet() []string {
	return []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
}

func getDependencyMap(alphabet []string, lineSlice []string) map[string][]string {
	dependencyMap := make(map[string][]string, 0)
	for _, line := range lineSlice {
		wordSlice := strings.Split(line, " ")
		owner := wordSlice[7]
		dependency := wordSlice[1]
		if letterSlice, ok := dependencyMap[owner]; ok {
			dependencyMap[owner] = append(letterSlice, dependency)
		} else {
			dependencyMap[owner] = make([]string, 0)
			dependencyMap[owner] = append(dependencyMap[owner], dependency)
		}
	}
	return dependencyMap
}

func getTasksToBeCompleted(dependencyMap map[string][]string) []string {
	tasksToBeCompletedMap := make(map[string]bool)
	for key, _ := range dependencyMap {
		tasksToBeCompletedMap[key] = true
		for _, letter := range dependencyMap[key] {
			tasksToBeCompletedMap[letter] = true
		}
	}
	tasksToBeCompletedSlice := make([]string, 0)
	for key, _ := range tasksToBeCompletedMap {
		tasksToBeCompletedSlice = append(tasksToBeCompletedSlice, key)
	}
	sort.Strings(tasksToBeCompletedSlice)
	return tasksToBeCompletedSlice
}

func dependenciesMet(letter string, completedTasks []string, dependencyMap map[string][]string) bool {
	if len(dependencyMap[letter]) == 0 {
		return true
	}
	for _, dependency := range dependencyMap[letter] {
		if !containsString(completedTasks, dependency) {
			return false
		}
	}
	return true
}

func completeTasks(dependencyMap map[string][]string) string {
	tasksToBeCompleted := getTasksToBeCompleted(dependencyMap)
	fmt.Println("tasksToBeCompleted:")
	fmt.Println(tasksToBeCompleted)
	completedTasks := make([]string, 0)
	letterIndex := 0
	for letterIndex < len(tasksToBeCompleted) {
		currentLetter := tasksToBeCompleted[letterIndex]
		if dependenciesMet(currentLetter, completedTasks, dependencyMap) {
			if !containsString(completedTasks, currentLetter) {
				completedTasks = append(completedTasks, currentLetter)
				letterIndex = -1
			}
		}
		letterIndex = letterIndex + 1
	}
	return strings.Join(completedTasks, "")
}

func main() {
	filename := "test"
	filename = "input7a"
	input := readFile(filename)
	fmt.Println("input:")
	fmt.Println(input)
	lineSlice := splitStringByNewlines(input)
	fmt.Println("lineSlice:")
	fmt.Println(lineSlice)
	alphabet := getAlphabet()
	dependencyMap := getDependencyMap(alphabet, lineSlice)
	fmt.Println("dependencyMap:")
	for _, letter := range alphabet {
		fmt.Printf("%s requires ", letter)
		fmt.Println(dependencyMap[letter])

	}
	orderOfTasks := completeTasks(dependencyMap)
	fmt.Println("orderOfTasks:")
	fmt.Println(orderOfTasks)
}
