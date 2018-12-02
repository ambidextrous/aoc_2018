package main

import (
	"fmt"
	"io/ioutil"
	"strings"
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
	for i, element := range string_array {
		fmt.Printf("i: %d = %s\n", i, element)
	}
	var no_empties []string
	for i, item := range string_array {
		if len(item) > 0 {
			no_empties = append(no_empties, item)
			fmt.Printf("i: %d = %s\n", i, item)
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

func main() {
	filename := "input2a"
	input := read_file(filename)
	fmt.Println(input)
	string_array := split_string_of_strings(input)
	fmt.Println(string_array)
	letter_count_map := gen_letter_freq_count_map()
	fmt.Println(letter_count_map)
	letter_counts := generate_letter_counts_for_strings(string_array)
	fmt.Println(letter_counts)
	instances_of_2 := count_instances_with_n_repetitions(letter_counts, 2)
	fmt.Println(instances_of_2)
	instances_of_3 := count_instances_with_n_repetitions(letter_counts, 3)
	fmt.Println(instances_of_3)
	fmt.Printf("checksum = %d\n", instances_of_2*instances_of_3)
}
