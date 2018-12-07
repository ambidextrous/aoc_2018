package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func Abs(x int) int {
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

func read_file(path string) string {
	dat, err := ioutil.ReadFile(path)
	check(err)
	contents := string(dat)
	return contents
}

func split_string_of_strings(string_of_strings string) []string {
	string_array := strings.Split(string_of_strings, "\n")
	var no_empties []string
	for _, item := range string_array {
		if len(item) > 0 {
			no_empties = append(no_empties, item)
		}
	}
	return no_empties
}

func gen_letter_freq_count_map() map[string]int {
	p := make([]byte, 26)
	for i := range p {
		p[i] = 'a' + byte(i)
	}
	alph_array := make(map[string]int)
	for _, letter := range p {
		alph_array[string(letter)] = 0
	}
	return alph_array
}

func generate_letter_counts_for_strings(string_array []string) []map[string]int {
	var letter_counts []map[string]int
	for _, str := range string_array {
		letter_count := gen_letter_freq_count_map()
		for _, char := range str {
			letter_count[string(char)] += 1
		}
		letter_counts = append(letter_counts, letter_count)
	}
	return letter_counts
}

func count_instances_with_n_repetitions(letter_counts []map[string]int, n int) int {
	total_instances := 0
	for _, letter_count := range letter_counts {
		instance := false
		for _, value := range letter_count {
			if value == n {
				instance = true
			}
		}
		if instance {
			total_instances += 1
		}
	}
	return total_instances
}

func get_similar_ids(string_array []string) map[string]bool {
	set := make(map[string]bool, 0)
	for _, s_1 := range string_array {
		for _, s_2 := range string_array {
			diff_count := 0
			for i, char_1 := range s_1 {
				if rune(s_2[i]) != char_1 {
					diff_count += 1
				}
			}
			if diff_count == 1 {
				set[s_1] = true
				set[s_2] = true
			}
		}
	}
	return set
}

func get_composite_id(set map[string]bool) string {
	var ids [2]string
	counter := 0
	for key, _ := range set {
		ids[counter] = key
		counter += 1
	}
	composite_string := ""
	first := ids[0]
	second := ids[1]
	for i, char := range first {
		if rune(char) == rune(second[i]) {
			composite_string += string(rune(char))
		}
	}
	return composite_string
}

//#1229 @ 441,869: 8x20
type rectangle struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func trimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

//#1229 @ 441,869: 8x20
func getRectangleFromString(s string) rectangle {
	parts := strings.Split(s, " ")
	id, err := strconv.Atoi(parts[0][1:])
	check(err)
	xAndY := strings.Split(parts[2], ",")
	x, err := strconv.Atoi(xAndY[0])
	check(err)
	yWithColon := xAndY[1]
	yWithoutColon := strings.Split(yWithColon, ":")
	y, err := strconv.Atoi(yWithoutColon[0])
	check(err)
	widthAndHeight := strings.Split(parts[3], "x")
	width, err := strconv.Atoi(widthAndHeight[0])
	check(err)
	height, err := strconv.Atoi(widthAndHeight[1])
	check(err)
	rect := rectangle{id: id, x: x, y: y, width: width, height: height}
	return rect
}

func getRectanglesFromStrings(lines []string) []rectangle {
	rectangles := make([]rectangle, 0)
	for _, line := range lines {
		rect := getRectangleFromString(line)
		rectangles = append(rectangles, rect)
	}
	return rectangles
}

func getFabricDimensions(rectangles []rectangle) (int, int, int, int) {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	for _, rect := range rectangles {
		if rect.x < minX {
			minX = rect.x
		}
		if rect.x+rect.width > maxX {
			maxX = rect.x + rect.width
		}
		if rect.y < minY {
			minY = rect.y
		}
		if rect.y+rect.height > maxY {
			maxY = rect.y + rect.height
		}
	}
	return minX, maxX, minY, maxY
}

func generateMatrix(maxX int, maxY int) [][]int {
	matrix := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		matrix[i] = make([]int, maxY)
	}
	return matrix
}

func markFabric(fabric [][]int, rectangles []rectangle) [][]int {
	for _, rect := range rectangles {
		for i := rect.x; i < rect.x+rect.width; i++ {
			for j := rect.y; j < rect.y+rect.height; j++ {
				fabric[i][j] += 1
			}
		}
	}
	return fabric
}

func findNonOverlapping(fabric [][]int, rectangles []rectangle) []rectangle {
	nonOverlapping := make([]rectangle, 0)
	for _, rect := range rectangles {
		overlapping := false
		for i := rect.x; i < rect.x+rect.width; i++ {
			for j := rect.y; j < rect.y+rect.height; j++ {
				if fabric[i][j] != 1 {
					overlapping = true
				}
			}
		}
		if !(overlapping) {
			nonOverlapping = append(nonOverlapping, rect)
		}
	}
	return nonOverlapping
}

func countOverlapping(fabric [][]int, rectangles []rectangle, maxX int, maxY int) int {
	overlappingCounter := 0
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if fabric[i][j] > 1 {
				overlappingCounter += 1
			}
		}
	}
	return overlappingCounter
}

type event struct {
	start         time.Time
	rawString     string
	isoTimeString string
	kind          string
	guardId       int
}

func getIsoStringFromRawString(raw string) string {
	s0 := raw
	sa1 := strings.Split(s0, "]")
	sa2 := strings.Split(sa1[0], "[")
	s1 := sa2[1]
	stringSplitBySpace := strings.Split(s1, " ")
	isoString := stringSplitBySpace[0] + "T" + stringSplitBySpace[1] + ":00Z"
	return isoString
}

func createEvents(rawStrings []string) []event {
	events := make([]event, 0)
	for _, s := range rawStrings {
		isoTimeString := getIsoStringFromRawString(s)
		//fmt.Printf("isoTimeString = %s\n", isoTimeString)
		t, err := time.Parse(time.RFC3339, isoTimeString)
		check(err)
		//fmt.Printf("t = %s\n", t.String())
		kind := ""
		guardId := 0
		if strings.Contains(s, "begins shift") {
			kind = "arrives"
			stringArray1 := strings.Split(s, "#")
			stringArray2 := strings.Split(stringArray1[1], " ")
			guardId, err = strconv.Atoi(stringArray2[0])
			check(err)
		} else if strings.Contains(s, "falls asleep") {
			kind = "fallsAsleep"
		} else if strings.Contains(s, "wakes up") {
			kind = "wakesUp"
		} else {
			panic(s)
		}
		e := event{start: t, rawString: s, isoTimeString: isoTimeString, kind: kind, guardId: guardId}
		fmt.Println(e)
		events = append(events, e)
	}
	return events
}

type timeSlice []event

type guard []watchPeriod

func generateWatchPeriods(events []event) map[int][]watchPeriod {
	endTime := events[len(events)-1].start.Add(time.Hour * 2)
	startTime := events[0].start.Add(time.Hour * -2)
	events = append(events, event{start: endTime})
	prependEvent := make([]event, 0)
	prependEvent = append(prependEvent, event{start: startTime})
	events = append(prependEvent, events...)
	watchPeriods := make(map[int][]watchPeriod)
	currentTime := startTime
	hr, min, _ := currentTime.Clock()
	currentGuardId := 0
	currentWatchPeriod := watchPeriod{}
	guardPresent := true
	guardAsleep := false
	currentlyInGuardPeriod := false
	currentEvent := events[0]
	nextEvent := events[1]
	eventCounter := 0
	//for eventCounter < len(events) {
	for currentTime.Before(endTime) {
		if !currentTime.Before(nextEvent.start) {
			eventCounter = eventCounter + 1
			currentEvent = nextEvent
			nextEvent = events[eventCounter+1]
			if currentEvent.guardId > 0 {
				currentGuardId = currentEvent.guardId
			}
			if currentEvent.kind == "fallsAsleep" {
				guardAsleep = true
			} else if currentEvent.kind == "wakesUp" {
				guardAsleep = false
			} else if currentEvent.kind == "arrives" {
				guardPresent = true
				guardAsleep = false
			} else {
				panic(currentEvent.kind)
			}
		}
		if currentlyInGuardPeriod {
			fmt.Printf("guardAsleep = %t\n", guardAsleep)
			fmt.Printf("guardPresent = %t\n", guardPresent)
			fmt.Printf("currentlyInGuardPeriod = %t\n", currentlyInGuardPeriod)
			fmt.Printf("currentTime = %s\n", currentTime)
			fmt.Println(currentWatchPeriod)
		}
		hr, min, _ = currentTime.Clock()
		if hr == 0 && min == 0 {
			currentWatchPeriod = watchPeriod{}
			currentlyInGuardPeriod = true
		} else if hr == 1 && min == 0 {
			guardPresent = false
			currentlyInGuardPeriod = false
			guardAsleep = false
			if _, ok := watchPeriods[currentGuardId]; ok {
				watchPeriods[currentGuardId] = append(watchPeriods[currentGuardId], currentWatchPeriod)
			} else {
				watchPeriods[currentGuardId] = make([]watchPeriod, 0)
				watchPeriods[currentGuardId] = append(watchPeriods[currentGuardId], currentWatchPeriod)
			}
		}
		if guardPresent && currentlyInGuardPeriod && guardAsleep {
			currentWatchPeriod[min] += 1
		}
		currentTime = currentTime.Add(time.Minute * 1)
	}
	return watchPeriods
}

type watchPeriod [60]int

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].start.Before(p[j].start)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortEventsByTime(events timeSlice) timeSlice {
	timeSortedEvents := make(timeSlice, 0, len(events))
	for _, d := range events {
		timeSortedEvents = append(timeSortedEvents, d)
	}
	return timeSortedEvents
}

func getSleepiestGuard(watchPeriods map[int][]watchPeriod) (int, int) {
	maxMinutesSlept := 0
	sleepiestGuardId := 0
	for guardId := range watchPeriods {
		sleepMinutesTotal := 0
		watches := watchPeriods[guardId]
		for _, watch := range watches {
			for _, minuteVal := range watch {
				sleepMinutesTotal += minuteVal
			}
		}
		if sleepMinutesTotal > maxMinutesSlept {
			maxMinutesSlept = sleepMinutesTotal
			sleepiestGuardId = guardId
		}
	}
	return sleepiestGuardId, maxMinutesSlept
}

func getSleepiestMinute(watches []watchPeriod) (int, int) {
	counter := watchPeriod{}
	for _, watch := range watches {
		for i, minuteVal := range watch {
			counter[i] += minuteVal
		}
	}
	sleepiestMinute := 0
	sleepMax := 0
	for j, val := range counter {
		if val > sleepMax {
			sleepMax = val
			sleepiestMinute = j
		}
	}
	return sleepiestMinute, sleepMax
}

func getSleepiestMinuteByGuardId(watchPeriods map[int][]watchPeriod) int {
	sleepiestGuardId := 0
	maxNumSleeps := 0
	minuteWithMostSleeps := 0
	for guardId := range watchPeriods {
		sleepiestMinute, sleepMax := getSleepiestMinute(watchPeriods[guardId])
		if sleepMax > maxNumSleeps {
			maxNumSleeps = sleepMax
			sleepiestGuardId = guardId
			minuteWithMostSleeps = sleepiestMinute
		}
	}
	return minuteWithMostSleeps * sleepiestGuardId
}

func sameTimeAndOppositePolarity(first rune, second rune) bool {
	if unicode.ToLower(first) == unicode.ToLower(second) {
		if unicode.IsLower(first) && unicode.IsUpper(second) || (unicode.IsUpper(first) && unicode.IsLower(second)) {
			return true
		}
	}
	return false
}

func reducePolymer(polymer string, cur rune) string {
	newPolymer := ""
	for {
		newPolymer = ""
		for i := 0; i < len(polymer); i++ {
			if i < len(polymer)-1 {
				if sameTimeAndOppositePolarity(rune(polymer[i]), rune(polymer[i+1])) {
					i += 1
				} else {
					newPolymer += string(rune(polymer[i]))
				}
			} else {
				newPolymer += string(rune(polymer[i]))
			}
		}
		if len(polymer) == len(newPolymer) {
			return newPolymer
		} else {
			polymer = newPolymer
		}
		fmt.Printf("curr = %s; len(polymer) = %d\n", string(cur), len(polymer))
	}
}

func getShortestPolymer(unalteredPolymer string) int {
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	shortest := len(unalteredPolymer)
	for i := 0; i < len(alphabetLower); i++ {
		newPolymer := strings.Replace(unalteredPolymer, string(alphabetLower[i]), "", -1)
		newPolymer = strings.Replace(newPolymer, string(alphabetUpper[i]), "", -1)
		reducedPolymer := reducePolymer(newPolymer, rune(alphabetLower[i]))
		if len(reducedPolymer) < shortest {
			shortest = len(reducedPolymer)
		}
	}
	return shortest
}

func getCordinatesSlice(rawInput string) []coordinate {
	coordinates := make([]coordinate, 0)
	linesArray := strings.Split(rawInput, "\n")
	for i, line := range linesArray {
		if len(line) > 0 {
			index := i + 1
			x, err := strconv.Atoi(strings.Split(line, ",")[0])
			check(err)
			x -= 1
			y, err := strconv.Atoi(strings.Split(line, " ")[1])
			check(err)
			y -= 1
			coord := coordinate{x: x, y: y, index: index, closestNeighbour: index}
			neighbours := make([]position, 0)
			coord.neighbours = neighbours
			coordinates = append(coordinates, coord)
		}
	}
	return coordinates
}

func getMaxXAndY(coordiantesSlice []coordinate) (int, int) {
	highestX := 0
	highestY := 0
	for _, coord := range coordiantesSlice {
		if coord.x > highestX {
			highestX = coord.x
		}
		if coord.y > highestY {
			highestY = coord.y
		}
	}
	return highestX, highestY
}

type coordinate struct {
	index                        int
	x                            int
	y                            int
	infinite                     bool
	area                         int
	neighbours                   []position
	closestNeighbour             int
	distanceFromClosestNeighbour int
}

func getEmptyGrid(maxX int, maxY int) [][]coordinate {
	xGrid := make([][]coordinate, maxX+1)
	for i, _ := range xGrid {
		yGrid := make([]coordinate, maxY+1)
		xGrid[i] = yGrid
	}
	return xGrid
}

func populateGrid(emptyGrid [][]coordinate, coordinates []coordinate) [][]coordinate {
	for _, coord := range coordinates {
		emptyGrid[coord.x][coord.y] = coord
	}
	return emptyGrid
}

type position struct {
	x int
	y int
}

func getNeighbouringPositions(initialPos position) []position {
	neighbouringPositions := make([]position, 0)
	neighbouringPositions = append(neighbouringPositions, position{x: initialPos.x + 0, y: initialPos.y + 1})
	neighbouringPositions = append(neighbouringPositions, position{x: initialPos.x + 0, y: initialPos.y - 1})
	neighbouringPositions = append(neighbouringPositions, position{x: initialPos.x + 1, y: initialPos.y + 0})
	neighbouringPositions = append(neighbouringPositions, position{x: initialPos.x - 1, y: initialPos.y + 0})
	return neighbouringPositions
}

func filterPositionsOutsideWorld(positions []position, maxX int, maxY int) []position {
	positionsInWorld := make([]position, 0)
	for _, pos := range positions {
		if pos.x < maxX && pos.y < maxY && pos.x >= 0 && pos.y >= 0 {
			positionsInWorld = append(positionsInWorld, pos)
		}
	}
	return positionsInWorld
}

func markNeighbours(coord coordinate, populatedGrid [][]coordinate, movesFromInitialCoordiate int) ([]position, [][]coordinate) {
	newNeighbours := make([]position, 0)
	for _, pos := range coord.neighbours {
		addNewNeighbours := false
		gridSquare := (populatedGrid)[pos.x][pos.y]
		if gridSquare.closestNeighbour == 0 {
			gridSquare.closestNeighbour = coord.index
			gridSquare.distanceFromClosestNeighbour = movesFromInitialCoordiate
			addNewNeighbours = true
		} else if gridSquare.distanceFromClosestNeighbour == movesFromInitialCoordiate && gridSquare.closestNeighbour != coord.index {
			gridSquare.closestNeighbour = -1
			gridSquare.distanceFromClosestNeighbour = movesFromInitialCoordiate
			addNewNeighbours = true
		}
		populatedGrid[pos.x][pos.y] = gridSquare
		if addNewNeighbours {
			potentialNeighbours := getNeighbouringPositions(pos)
			neighboursInWorld := filterPositionsOutsideWorld(potentialNeighbours, len(populatedGrid), len(populatedGrid[0]))
			newNeighbours = append(newNeighbours, neighboursInWorld...)
		}
	}
	return newNeighbours, populatedGrid
}

func markGrid(coordinates []coordinate, populatedGrid [][]coordinate) [][]coordinate {
	initialNeighbours := make([][]position, 0)
	for _, coord := range coordinates {
		neighbours := getNeighbouringPositions(position{coord.x, coord.y})
		neighboursInWorld := filterPositionsOutsideWorld(neighbours, len(populatedGrid), len(populatedGrid[0]))
		initialNeighbours = append(initialNeighbours, neighboursInWorld)
	}
	firstPass := true
	numCoordinatesStillMarking := len(coordinates)
	moveCounter := 0
	coordinateNeighbourMap := make(map[int][]position, 0)
	for numCoordinatesStillMarking > 0 {
		moveCounter += 1
		if moveCounter > 2000 {
			panic("Too many moves!")
		}
		fmt.Printf("moveCounter = %s\n", moveCounter)
		numCoordinatesStillMarking = 0
		for i, coordinate := range coordinates {
			if firstPass {
				coordinate.neighbours = initialNeighbours[i]
			} else {
				coordinate.neighbours = coordinateNeighbourMap[i+1]
			}
			coordinateNeighbourMap[i+1], populatedGrid = markNeighbours(coordinate, populatedGrid, moveCounter)
			if len(coordinateNeighbourMap[i+1]) > 0 {
				numCoordinatesStillMarking += 1
			}
		}
		firstPass = false
	}
	return populatedGrid
}

func prettyPrintGrid(grid [][]coordinate) {
	for _, row := range grid {
		output := ""
		for _, square := range row {
			if square.closestNeighbour >= 0 {
				output += " "
			}
			output += strconv.Itoa(square.closestNeighbour) + " "
		}
		fmt.Println(output + "\n")
	}
}

func countPopulationTotals(grid [][]coordinate) map[int]int {
	countMap := make(map[int]int, 0)
	for _, line := range grid {
		for _, square := range line {
			if _, ok := countMap[square.closestNeighbour]; ok {
				countMap[square.closestNeighbour] += 1
			} else {
				countMap[square.closestNeighbour] = 1
			}
		}
	}
	return countMap
}

func getFiniteAreas(grid [][]coordinate) map[int]bool {
	infiniteMap := make(map[int]bool, 0)
	for i, line := range grid {
		for j, square := range line {
			if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1 {
				if _, ok := infiniteMap[square.closestNeighbour]; ok {
					infiniteMap[square.closestNeighbour] = true
				} else {
					infiniteMap[square.closestNeighbour] = true
				}
			}
		}
	}
	return infiniteMap
}

func getLargestFiniteArea(infiniteMap map[int]bool, populationTotalMap map[int]int) int {
	largestFiniteArea := 0
	for key, _ := range populationTotalMap {
		if key != -1 {
			if _, ok := infiniteMap[key]; ok {
				// Do nothing
			} else {
				if populationTotalMap[key] > largestFiniteArea {
					largestFiniteArea = populationTotalMap[key]
				}
			}
		}
	}
	return largestFiniteArea
}

type distAndCoord struct {
	dist  int
	coord coordinate
}

func getNearestCoordinateIndex(pos position, coordinates []coordinate) int {
	minDist := 100000000
	distanceFrequencyMap := make(map[int]int, 0)
	closesCoordIndex := -2
	for i, coord := range coordinates {
		dist := Abs(pos.x-coord.x) + Abs(pos.y-coord.y)
		if dist < minDist {
			minDist = dist
			closesCoordIndex = i + 1
		}
		if _, ok := distanceFrequencyMap[dist]; ok {
			distanceFrequencyMap[dist] += 1
		} else {
			distanceFrequencyMap[dist] = 1
		}
	}
	if distanceFrequencyMap[minDist] > 1 {
		return -1
	}
	return closesCoordIndex
}

func markDistancesWithoutBlowingStack(coordinates []coordinate, grid [][]coordinate) [][]coordinate {
	for i, line := range grid {
		for j, _ := range line {
			currentPos := position{x: i, y: j}
			grid[i][j].closestNeighbour = getNearestCoordinateIndex(currentPos, coordinates)
		}
	}
	return grid
}

func isWithinSafeRange(pos position, coordinates []coordinate) bool {
	safeRangeLimit := 10000
	totalDistance := 0
	for _, coord := range coordinates {
		dist := Abs(pos.x-coord.x) + Abs(pos.y-coord.y)
		totalDistance += dist
	}
	if totalDistance >= safeRangeLimit {
		return false
	}
	return true
}

func getSafeAreaSize(coordinates []coordinate, grid [][]coordinate) int {
	safeAreaSize := 0
	for i, line := range grid {
		for j, _ := range line {
			currentPos := position{x: i, y: j}
			isSafe := isWithinSafeRange(currentPos, coordinates)
			if isSafe {
				safeAreaSize += 1
			}
		}
	}
	return safeAreaSize
}

func main() {
	filename := "test"
	filename = "input6a"
	input := read_file(filename)
	fmt.Println("input:")
	fmt.Println(input)
	coordinatesSlice := getCordinatesSlice(input)
	fmt.Println("Coordiante slice:")
	fmt.Println(coordinatesSlice)
	maxX, maxY := getMaxXAndY(coordinatesSlice)
	fmt.Printf("maxX = %d; maxY = %d\n", maxX, maxY)
	emptyGrid := getEmptyGrid(maxX, maxY)

	fmt.Println("Marking grid without blowing stack...")
	markedGrid := markDistancesWithoutBlowingStack(coordinatesSlice, emptyGrid)
	//prettyPrintGrid(markedGrid)

	populationTotalMap := countPopulationTotals(markedGrid)
	infiniteMap := getFiniteAreas(markedGrid)
	largestFiniteArea := getLargestFiniteArea(infiniteMap, populationTotalMap)
	fmt.Printf("largest finite area = %d\n", largestFiniteArea)

	safeAreaSize := getSafeAreaSize(coordinatesSlice, emptyGrid)
	fmt.Printf("Safe area size = %d\n", safeAreaSize)
}
