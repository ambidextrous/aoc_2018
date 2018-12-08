package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

type worker struct {
	id                   int
	currentTask          string
	secondsWorkingOnTask int
	totalTaskDuration    int
}

func generateWorkers(numWorkers int) []worker {
	workers := make([]worker, 0)
	for i := 0; i < numWorkers; i++ {
		newWorker := worker{id: i}
		workers = append(workers, newWorker)
	}
	return workers
}

func getReadyToGoTasks(dependencyMap map[string][]string, completedTasks []string) []string {
	readyToGoTasks := make([]string, 0)
	tasksToBeCompleted := getTasksToBeCompleted(dependencyMap)
	fmt.Println("tasksToBeCompleted:")
	fmt.Println(tasksToBeCompleted)
	for _, task := range tasksToBeCompleted {
		readyToGo := false
		if !containsString(completedTasks, task) {
			if len(dependencyMap[task]) == 0 {
				readyToGo = true
			} else {
				dependenciesMet := true
				for _, dependency := range dependencyMap[task] {
					if !containsString(completedTasks, dependency) {
						dependenciesMet = false
					}
					readyToGo = dependenciesMet
				}
			}
		}
		if readyToGo {
			readyToGoTasks = append(readyToGoTasks, task)
		}
	}
	sort.Strings(readyToGoTasks)
	return readyToGoTasks
}

func getReadyToGoWorkers(workers []worker) []worker {
	readyToGoWorkers := make([]worker, 0)
	for _, w := range workers {
		if len(w.currentTask) == 0 {
			readyToGoWorkers = append(readyToGoWorkers, w)
		}
	}
	return readyToGoWorkers
}

func workOnTasks(workers []worker, completedTasks []string) ([]worker, []string) {
	for i := 0; i < len(workers); i++ {
		if len(workers[i].currentTask) > 0 {
			workers[i].secondsWorkingOnTask++
			if workers[i].secondsWorkingOnTask == workers[i].totalTaskDuration {
				completedTasks = append(completedTasks, workers[i].currentTask)
				workers[i].currentTask = ""
				workers[i].secondsWorkingOnTask = 0
				workers[i].totalTaskDuration = 0
			}
		}
	}
	return workers, completedTasks
}

func filterAssignedTasks(readyToGoTasks []string, workers []worker) []string {
	unnassignedTasks := make([]string, 0)
	for _, task := range readyToGoTasks {
		assigned := false
		for _, w := range workers {
			if w.currentTask == task {
				assigned = true
			}
		}
		if !assigned {
			unnassignedTasks = append(unnassignedTasks, task)
		}
	}
	return unnassignedTasks
}

func getAlphabetIndex(letter string) int {
	alphabet := getAlphabet()
	for i, char := range alphabet {
		if letter == char {
			return i + 1
		}
	}
	panic("letter " + letter + " not found in alphabet")
	return 0
}

func assignTasks(workers []worker, unnassignedReadyToGoTasks []string, offset int) []worker {
	for _, task := range unnassignedReadyToGoTasks {
		unnassigned := true
		for i := 0; i < len(workers); i++ {
			if len(workers[i].currentTask) == 0 && unnassigned {
				workers[i].currentTask = task
				workers[i].totalTaskDuration = offset + getAlphabetIndex(task)
				workers[i].secondsWorkingOnTask = 0
				unnassigned = false
			}
		}
	}
	return workers
}

func prettyPrintWorldState(workers []worker, completedTasks []string, numSeconds int) {
	output := strconv.Itoa(numSeconds) + " "
	for _, w := range workers {
		if len(w.currentTask) > 0 {
			output += w.currentTask
		} else {
			output += "_"
		}
	}
	output += " "
	output += strings.Join(completedTasks, ", ")
	fmt.Println(output)
}

func getNumSecondsToComplete(dependencyMap map[string][]string, numWorkers int, offset int) int {
	numSecondsToComplete := 0
	workers := generateWorkers(numWorkers)
	fmt.Println(workers)
	completedTasks := make([]string, 0)
	readyToGoTasks := make([]string, 0)
	//for len(completedTasks) < len(dependencyMap) {
	readyToGoTasks = getReadyToGoTasks(dependencyMap, completedTasks)
	fmt.Println("readyToGoTasks:")
	fmt.Println(readyToGoTasks)
	unnassignedReadyToGoTasks := filterAssignedTasks(readyToGoTasks, workers)
	fmt.Println("unnassignedReadyToGoTasks:")
	fmt.Println(unnassignedReadyToGoTasks)
	sort.Strings(unnassignedReadyToGoTasks)
	readyToGoWorkers := getReadyToGoWorkers(workers)
	fmt.Println("readyToGoWorkers:")
	fmt.Println(readyToGoWorkers)
	workers, completedTasks = workOnTasks(readyToGoWorkers, completedTasks)
	workers = assignTasks(workers, unnassignedReadyToGoTasks, offset)
	prettyPrintWorldState(workers, completedTasks, numSecondsToComplete)
	numSecondsToComplete++
	fmt.Println()
	//}
	return numSecondsToComplete
}

func main() {
	filename := "test"
	//filename = "input7a"
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
	fmt.Println("Answer to part 1: orderOfTasks:")
	fmt.Println(orderOfTasks)
	numSecondsToComplete := getNumSecondsToComplete(dependencyMap, 2, 0)
	fmt.Println("numSecondsToComplete:")
	fmt.Println(numSecondsToComplete)
}
