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

func main() {
	file_location := "input1a"
	input := read_file(file_location)
	convert_string_to_int_list(input)
	ints_array := convert_string_to_int_list(input)
	total := 0
	for _, num := range ints_array {
		total += num
	}
	fmt.Println(total)
}
