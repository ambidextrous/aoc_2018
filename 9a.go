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

func removeMarbleAtIndex(index int, circle []int) []int {
	if index == len(circle) {
		circle = circle[:len(circle)-1]
		return circle
	}
	firstPart := make([]int, 0)
	firstPart = append(firstPart, circle[:index]...)
	secondPart := make([]int, 0)
	secondPart = append(secondPart, circle[index+1:]...)
	circle = append(firstPart, secondPart...)
	return circle
}

func getHighestPlayerScore(players map[int][]int) (int, int) {
	fmt.Printf("players: %+v\n", players)
	highestTotalScore := 0
	highestScore := 0
	for key := range players {
		totalScore := 0
		for _, score := range players[key] {
			totalScore += score
			if score > highestScore {
				highestScore = score
			}
		}
		if totalScore > highestTotalScore {
			highestTotalScore = totalScore
		}

	}
	return highestScore, highestTotalScore
}

func playMarbleGame(numPlayers int, finalLastMarbleScore int) (int, int) {
	REGULAR_JUMP_DIST := 2
	SPECIAL_JUMP_DIST := 7
	SPECIAL_DIVISOR := 23
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
	//for currentMarbleScore < 27 {
	for currentMarbleScore < finalLastMarbleScore {
		if roundCounter%SPECIAL_DIVISOR == 0 && roundCounter != 0 {
			currentPlayer++
			if currentPlayer > numPlayers {
				currentPlayer = 1
			}
			removalPosition := currentMarblePos - SPECIAL_JUMP_DIST - REGULAR_JUMP_DIST
			if removalPosition < 0 {
				removalPosition = len(circle) + removalPosition
			}
			nextMarblePos = currentMarblePos - SPECIAL_JUMP_DIST
			if nextMarblePos < 0 {
				nextMarblePos = len(circle) + nextMarblePos
			}
			marbleScore := circle[removalPosition]
			circle = removeMarbleAtIndex(removalPosition, circle)
			playerRoundScore := marbleScore + roundCounter
			players[currentPlayer] = append(players[currentPlayer], playerRoundScore)
			if playerRoundScore == finalLastMarbleScore {
				//fmt.Printf("lastMarbleWorth = %d \n", playerRoundScore)
				return getHighestPlayerScore(players)
			}
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
			nextMarblePos = currentMarblePos + REGULAR_JUMP_DIST
			if nextMarblePos > len(circle) {
				nextMarblePos = 1
			}
		}
		//fmt.Println("currentMarblePos:")
		//fmt.Println(currentMarblePos)
		//fmt.Printf("round %d: player %d: %v\n", roundCounter, currentPlayer, circle)
		roundCounter++
		currentMarbleScore++
		currentMarblePos = nextMarblePos
		//fmt.Println()
		if roundCounter%1000 == 0 {
			fmt.Printf("roundCounter = %d\n", roundCounter)
		}
	}
	return getHighestPlayerScore(players)
}

type node struct {
	next     *node
	nextName int
	prev     *node
	prevName int
	name     int
}

func jumpBackDist(currentNode *node, dist int) (node, int) {
	for i := 0; i < dist; i++ {
		if currentNode.name == 0 {
			fmt.Println()
			fmt.Println("WRAPPING AROUND!!")
			fmt.Println()
		}
		currentNode = currentNode.prev
	}
	replacementNode := node{name: currentNode.next.name, prev: currentNode.prev, prevName: currentNode.prevName, next: currentNode.next.next, nextName: currentNode.next.nextName}
	replacementNode.next.prev = &replacementNode
	replacementNode.next.prevName = replacementNode.name
	replacementNode.prev.next = &replacementNode
	replacementNode.prev.nextName = replacementNode.name
	//fmt.Printf("!!!prevNode: %+v\n", replacementNode.prev)
	//fmt.Printf("!!!nextNode: %+v\n", replacementNode.next)
	return replacementNode, currentNode.name
}

func removeNode(currentNode *node) {
	currentNode.prev.next = currentNode.next
	currentNode.prev.nextName = currentNode.nextName
	currentNode.next.prev = currentNode.prev
	currentNode.next.prevName = currentNode.prevName
	currentNode = currentNode.next
}

func playLinkedGame(numPlayers int, finalLastMarbleScore int) (int, int) {
	players := make(map[int][]int)
	for i := 1; i <= numPlayers; i++ {
		players[i] = make([]int, 0)
	}
	// Set-up
	zeroNode := node{name: 0}
	oneNode := node{name: 1, next: &zeroNode, nextName: 0}
	twoNode := node{name: 2, next: &oneNode, prev: &zeroNode, nextName: 1, prevName: 0}
	zeroNode.prev = &oneNode
	zeroNode.prevName = 1
	zeroNode.next = &twoNode
	zeroNode.nextName = 2
	oneNode.prev = &twoNode
	oneNode.prevName = 2
	fmt.Printf("%+v\n", zeroNode)
	fmt.Printf("%+v\n", oneNode)
	fmt.Printf("%+v\n", twoNode)
	currentNode := twoNode
	roundCounter := 3
	playerCounter := 3
	for roundCounter < finalLastMarbleScore {
		nodeString := ""
		startingNode := &currentNode
		for i := 0; i < roundCounter+4; i++ {
			nextNode := startingNode.next
			nodeString += strconv.Itoa(nextNode.name) + ", "
			startingNode = nextNode
		}
		nodeString += "\n"
		fmt.Println("nodeString:")
		fmt.Println(nodeString)
		if roundCounter%23 == 0 {
			newScore := 0
			currentNode, newScore = jumpBackDist(&currentNode, 7)
			//fmt.Printf("!!!currentNode: %+v\n", currentNode)
			//fmt.Printf("!!!newScore: %d\n", newScore)
			//jumpedToNodeScore := currentNode.name

			playerRoundScore := newScore + roundCounter
			players[playerCounter] = append(players[playerCounter], playerRoundScore)
			//if playerRoundScore == finalLastMarbleScore {
			//	//fmt.Println("GOT HERE!")
			//	return getHighestPlayerScore(players)
			//}
			removeNode(&currentNode)

		} else {
			followingNode := currentNode.next
			newNode := node{name: roundCounter, next: followingNode.next, nextName: followingNode.nextName, prev: followingNode, prevName: followingNode.name}
			newNode.next.prev = &newNode
			newNode.next.prevName = newNode.name
			followingNode.next = &newNode
			followingNode.nextName = newNode.name
			currentNode = newNode
		}
		//fmt.Printf("roundCounter: %d\n", roundCounter)
		//fmt.Printf("playerCounter: %d\n", playerCounter)

		//fmt.Printf("currentNode: %+v\n", currentNode)
		fmt.Printf("prevNode:%d; currentNode:%d; nextNode:%d\n", currentNode.prev.name, currentNode.name, currentNode.next.name)
		//fmt.Println()
		roundCounter++
		playerCounter++
		if playerCounter > numPlayers {
			playerCounter = 1
		}
		if roundCounter%1000 == 0 {
			fmt.Printf("roundCounter = %d\n", roundCounter)
		}
	}
	for i, player := range players {
		fmt.Printf("Player %d:\n", i)
		fmt.Println(player)
	}
	return getHighestPlayerScore(players)
}

func main() {
	//highestPlayerScore, highestTotalScore := playLinkedGame(10, 1618)
	highestPlayerScore, highestTotalScore := playLinkedGame(9, 25)
	fmt.Println()
	fmt.Println("highestPlayerScore:")
	fmt.Println(highestPlayerScore)
	fmt.Println("highestTotalScore:")
	fmt.Println(highestTotalScore)

	//highestPlayerScore := playLinkedGame(10, 1618)
	//highestPlayerScore := playLinkedGame(493, 71863)
	//fmt.Println("highestPlayerScore:")
	//fmt.Println(highestPlayerScore)
	//filename := "test"
	//filename = "input8a"
	//input := readFile(filename)
	//input = strings.Replace(input, "\n", "", -1)
	//fmt.Println("input:")
	//fmt.Println(input)

	//highestPlayerScore, highestTotalScore := playMarbleGame(13, 8001)
	//fmt.Println("highestPlayerScore:")
	//fmt.Println(highestPlayerScore)
	//fmt.Println("highestTotalScore:")
	//fmt.Println(highestTotalScore)

	//highestScore := playMarbleGame(10, 1618)
	//fmt.Println("highestScore:")
	//fmt.Println(highestScore)
}
