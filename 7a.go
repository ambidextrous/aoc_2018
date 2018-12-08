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

//func getCordinatesSlice(rawInput string) []coordinate {
//	coordinates := make([]coordinate, 0)
//	linesArray := strings.Split(rawInput, "\n")
//	for i, line := range linesArray {
//		if len(line) > 0 {
//			index := i + 1
//			x, err := strconv.Atoi(strings.Split(line, ",")[0])
//			check(err)
//			x -= 1
//			y, err := strconv.Atoi(strings.Split(line, " ")[1])
//			check(err)
//			y -= 1
//			coord := coordinate{x: x, y: y, index: index, closestNeighbour: index}
//			neighbours := make([]position, 0)
//			coord.neighbours = neighbours
//			coordinates = append(coordinates, coord)
//		}
//	}
//	return coordinates
//}
//
//func getMaxXAndY(coordiantesSlice []coordinate) (int, int) {
//	highestX := 0
//	highestY := 0
//	for _, coord := range coordiantesSlice {
//		if coord.x > highestX {
//			highestX = coord.x
//		}
//		if coord.y > highestY {
//			highestY = coord.y
//		}
//	}
//	return highestX, highestY
//}
//
//type coordinate struct {
//	index                        int
//	x                            int
//	y                            int
//	infinite                     bool
//	area                         int
//	neighbours                   []position
//	closestNeighbour             int
//	distanceFromClosestNeighbour int
//}
//
//func getEmptyGrid(maxX int, maxY int) [][]coordinate {
//	xGrid := make([][]coordinate, maxX+1)
//	for i, _ := range xGrid {
//		yGrid := make([]coordinate, maxY+1)
//		xGrid[i] = yGrid
//	}
//	return xGrid
//}
//
//func populateGrid(emptyGrid [][]coordinate, coordinates []coordinate) [][]coordinate {
//	for _, coord := range coordinates {
//		emptyGrid[coord.x][coord.y] = coord
//	}
//	return emptyGrid
//}
//
//type position struct {
//	x int
//	y int
//}
//
//func getNeighbouringPositions(initialPos position) []position {
//	neighbouringPositions := make([]position, 0)
//	neighbouringPositions = append(neighbouringPositions, position{x: initialPos.x + 0, y: initialPos.y + 1})
//	neighbouringPositions = append(neighbouringPositions, position{x: initialPos.x + 0, y: initialPos.y - 1})
//	neighbouringPositions = append(neighbouringPositions, position{x: initialPos.x + 1, y: initialPos.y + 0})
//	neighbouringPositions = append(neighbouringPositions, position{x: initialPos.x - 1, y: initialPos.y + 0})
//	return neighbouringPositions
//}
//
//func filterPositionsOutsideWorld(positions []position, maxX int, maxY int) []position {
//	positionsInWorld := make([]position, 0)
//	for _, pos := range positions {
//		if pos.x < maxX && pos.y < maxY && pos.x >= 0 && pos.y >= 0 {
//			positionsInWorld = append(positionsInWorld, pos)
//		}
//	}
//	return positionsInWorld
//}
//
//func markNeighbours(coord coordinate, populatedGrid [][]coordinate, movesFromInitialCoordiate int) ([]position, [][]coordinate) {
//	newNeighbours := make([]position, 0)
//	for _, pos := range coord.neighbours {
//		addNewNeighbours := false
//		gridSquare := (populatedGrid)[pos.x][pos.y]
//		if gridSquare.closestNeighbour == 0 {
//			gridSquare.closestNeighbour = coord.index
//			gridSquare.distanceFromClosestNeighbour = movesFromInitialCoordiate
//			addNewNeighbours = true
//		} else if gridSquare.distanceFromClosestNeighbour == movesFromInitialCoordiate && gridSquare.closestNeighbour != coord.index {
//			gridSquare.closestNeighbour = -1
//			gridSquare.distanceFromClosestNeighbour = movesFromInitialCoordiate
//			addNewNeighbours = true
//		}
//		populatedGrid[pos.x][pos.y] = gridSquare
//		if addNewNeighbours {
//			potentialNeighbours := getNeighbouringPositions(pos)
//			neighboursInWorld := filterPositionsOutsideWorld(potentialNeighbours, len(populatedGrid), len(populatedGrid[0]))
//			newNeighbours = append(newNeighbours, neighboursInWorld...)
//		}
//	}
//	return newNeighbours, populatedGrid
//}
//
//func markGrid(coordinates []coordinate, populatedGrid [][]coordinate) [][]coordinate {
//	initialNeighbours := make([][]position, 0)
//	for _, coord := range coordinates {
//		neighbours := getNeighbouringPositions(position{coord.x, coord.y})
//		neighboursInWorld := filterPositionsOutsideWorld(neighbours, len(populatedGrid), len(populatedGrid[0]))
//		initialNeighbours = append(initialNeighbours, neighboursInWorld)
//	}
//	firstPass := true
//	numCoordinatesStillMarking := len(coordinates)
//	moveCounter := 0
//	coordinateNeighbourMap := make(map[int][]position, 0)
//	for numCoordinatesStillMarking > 0 {
//		moveCounter += 1
//		if moveCounter > 2000 {
//			panic("Too many moves!")
//		}
//		fmt.Printf("moveCounter = %s\n", moveCounter)
//		numCoordinatesStillMarking = 0
//		for i, coordinate := range coordinates {
//			if firstPass {
//				coordinate.neighbours = initialNeighbours[i]
//			} else {
//				coordinate.neighbours = coordinateNeighbourMap[i+1]
//			}
//			coordinateNeighbourMap[i+1], populatedGrid = markNeighbours(coordinate, populatedGrid, moveCounter)
//			if len(coordinateNeighbourMap[i+1]) > 0 {
//				numCoordinatesStillMarking += 1
//			}
//		}
//		firstPass = false
//	}
//	return populatedGrid
//}
//
//func prettyPrintGrid(grid [][]coordinate) {
//	for _, row := range grid {
//		output := ""
//		for _, square := range row {
//			if square.closestNeighbour >= 0 {
//				output += " "
//			}
//			output += strconv.Itoa(square.closestNeighbour) + " "
//		}
//		fmt.Println(output + "\n")
//	}
//}
//
//func countPopulationTotals(grid [][]coordinate) map[int]int {
//	countMap := make(map[int]int, 0)
//	for _, line := range grid {
//		for _, square := range line {
//			if _, ok := countMap[square.closestNeighbour]; ok {
//				countMap[square.closestNeighbour] += 1
//			} else {
//				countMap[square.closestNeighbour] = 1
//			}
//		}
//	}
//	return countMap
//}
//
//func getFiniteAreas(grid [][]coordinate) map[int]bool {
//	infiniteMap := make(map[int]bool, 0)
//	for i, line := range grid {
//		for j, square := range line {
//			if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1 {
//				if _, ok := infiniteMap[square.closestNeighbour]; ok {
//					infiniteMap[square.closestNeighbour] = true
//				} else {
//					infiniteMap[square.closestNeighbour] = true
//				}
//			}
//		}
//	}
//	return infiniteMap
//}
//
//func getLargestFiniteArea(infiniteMap map[int]bool, populationTotalMap map[int]int) int {
//	largestFiniteArea := 0
//	for key, _ := range populationTotalMap {
//		if key != -1 {
//			if _, ok := infiniteMap[key]; ok {
//				// Do nothing
//			} else {
//				if populationTotalMap[key] > largestFiniteArea {
//					largestFiniteArea = populationTotalMap[key]
//				}
//			}
//		}
//	}
//	return largestFiniteArea
//}
//
//type distAndCoord struct {
//	dist  int
//	coord coordinate
//}
//
//func getNearestCoordinateIndex(pos position, coordinates []coordinate) int {
//	minDist := 100000000
//	distanceFrequencyMap := make(map[int]int, 0)
//	closesCoordIndex := -2
//	for i, coord := range coordinates {
//		dist := Abs(pos.x-coord.x) + Abs(pos.y-coord.y)
//		if dist < minDist {
//			minDist = dist
//			closesCoordIndex = i + 1
//		}
//		if _, ok := distanceFrequencyMap[dist]; ok {
//			distanceFrequencyMap[dist] += 1
//		} else {
//			distanceFrequencyMap[dist] = 1
//		}
//	}
//	if distanceFrequencyMap[minDist] > 1 {
//		return -1
//	}
//	return closesCoordIndex
//}
//
//func markDistancesWithoutBlowingStack(coordinates []coordinate, grid [][]coordinate) [][]coordinate {
//	for i, line := range grid {
//		for j, _ := range line {
//			currentPos := position{x: i, y: j}
//			grid[i][j].closestNeighbour = getNearestCoordinateIndex(currentPos, coordinates)
//		}
//	}
//	return grid
//}
//
//func isWithinSafeRange(pos position, coordinates []coordinate) bool {
//	safeRangeLimit := 10000
//	totalDistance := 0
//	for _, coord := range coordinates {
//		dist := Abs(pos.x-coord.x) + Abs(pos.y-coord.y)
//		totalDistance += dist
//	}
//	if totalDistance >= safeRangeLimit {
//		return false
//	}
//	return true
//}
//
//func getSafeAreaSize(coordinates []coordinate, grid [][]coordinate) int {
//	safeAreaSize := 0
//	for i, line := range grid {
//		for j, _ := range line {
//			currentPos := position{x: i, y: j}
//			isSafe := isWithinSafeRange(currentPos, coordinates)
//			if isSafe {
//				safeAreaSize += 1
//			}
//		}
//	}
//	return safeAreaSize
//}

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

func completeTasks(dependencyMap map[string][]string) string {
	tasksToBeCompleted := getTasksToBeCompleted(dependencyMap)
	fmt.Println("tasksToBeCompleted:")
	fmt.Println(tasksToBeCompleted)
	completedTasks := ""
	letterIndex := 0
	for letterIndex < len(tasksToBeCompleted) {
		currentLetter := tasksToBeCompleted[letterIndex]
		if
	}
	return completedTasks
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
	fmt.Println("orderOfTasks:")
	fmt.Println(orderOfTasks)
	//fmt.Println(input)
	//coordinatesSlice := getCordinatesSlice(input)
	//fmt.Println("Coordiante slice:")
	//fmt.Println(coordinatesSlice)
	//maxX, maxY := getMaxXAndY(coordinatesSlice)
	//fmt.Printf("maxX = %d; maxY = %d\n", maxX, maxY)
	//emptyGrid := getEmptyGrid(maxX, maxY)

	//fmt.Println("Marking grid without blowing stack...")
	//markedGrid := markDistancesWithoutBlowingStack(coordinatesSlice, emptyGrid)
	////prettyPrintGrid(markedGrid)

	//populationTotalMap := countPopulationTotals(markedGrid)
	//infiniteMap := getFiniteAreas(markedGrid)
	//largestFiniteArea := getLargestFiniteArea(infiniteMap, populationTotalMap)
	//fmt.Printf("largest finite area = %d\n", largestFiniteArea)

	//safeAreaSize := getSafeAreaSize(coordinatesSlice, emptyGrid)
	//fmt.Printf("Safe area size = %d\n", safeAreaSize)
}
