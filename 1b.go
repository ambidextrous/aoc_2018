package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_webpage(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body_bytes, err := ioutil.ReadAll(resp.Body)
	body_string := string(body_bytes)
	if err != nil {
		panic(err)
	}
	return body_string
}

func read_file(path string) string {
	dat, err := ioutil.ReadFile(path)
	check(err)
	contents := string(dat)
	return contents
}

func convert_string_to_int_list(ints_string string) []int {
	string_array := strings.Split(ints_string, "\n")
	ints_array := make([]int, len(string_array)-1)
	var err error
	for i := range ints_array {
		ints_array[i], err = strconv.Atoi(string_array[i])
		check(err)
	}
	return ints_array
}

func get_first_repeating_total(ints_list []int) (int, error) {
	totals := make(map[int]bool)
	totals[0] = true
	current_total := 0
	for {
		for _, num := range ints_list {
			current_total += num
			if _, ok := totals[current_total]; ok {
				return current_total, nil
			} else {
				fmt.Printf("Current number = `%d`.\n", num)
				fmt.Printf("New total key `%d` added to set.\n\n", current_total)
				totals[current_total] = true
			}
		}
	}
}

func main() {
	file_location := "input1a"
	input := read_file(file_location)
	convert_string_to_int_list(input)
	ints_array := convert_string_to_int_list(input)
	first_repeating_total, err := get_first_repeating_total(ints_array)
	check(err)
	fmt.Printf("First repeating total = %d\n\n", first_repeating_total)
}
