package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"strconv"
	"time"
)

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
	id int
	x int
	y int
	width int
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
		if rect.x + rect.width > maxX {
			maxX = rect.x + rect.width
		}
		if rect.y < minY {
			minY = rect.y
		}
		if rect.y + rect.height > maxY {
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
		for i := rect.x; i < rect.x + rect.width; i++ {
			for j := rect.y; j < rect.y + rect.height; j++ {
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
		for i := rect.x; i < rect.x + rect.width; i++ {
			for j := rect.y; j < rect.y + rect.height; j++ {
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
	start time.Time
	rawString string
	isoTimeString string
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
		fmt.Printf("isoTimeString = %s\n", isoTimeString)
		t, err := time.Parse(time.RFC3339, isoTimeString)
		check(err)
		fmt.Printf("t = %s\n", t.String())
		e := event{start: t, rawString: s, isoTimeString: isoTimeString}
		events = append(events, e)
	}
	return events
}

//func sortByTime(events []event) []event {
//	sort.Slice(events, func(i, j int) bool { return events[i].start.Before(events[j].start)})
//}

type timeSlice []event

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

func main() {
	filename := "test"
	input := read_file(filename)
	stringArray := split_string_of_strings(input)
	fmt.Println(stringArray)
	events := createEvents(stringArray)
	fmt.Println(events)
	timeSortedEvents := sortEventsByTime(events)
	sort.Sort(timeSortedEvents)
	fmt.Println(timeSortedEvents)
	fmt.Println(timeSortedEvents[0])
	//rectangles := getRectanglesFromStrings(string_array)
	//minX, maxX, minY, maxY := getFabricDimensions(rectangles)
	//fmt.Printf("minX=%d, maxX=%d, minY=%d, maxY=%d\n", minX, maxX, minY, maxY)
	//if (minX != 0) || (minY != 0) {
	//	panic("non-zero min values for fabric matrix")
	//}
	//fabricMatrix := generateMatrix(maxX + 1, maxY + 1)
	//markedFabric := markFabric(fabricMatrix, rectangles)
	//overlappingInstances := countOverlapping(markedFabric, rectangles, maxX + 1, maxY +1)
	//fmt.Printf("Overlapping square inches = %d\n", overlappingInstances) 
	//nonOverlapping := findNonOverlapping(markedFabric, rectangles)
	//fmt.Printf("Id of non-overlapping piece = #%d\n", nonOverlapping[0].id)
}
