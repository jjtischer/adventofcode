package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	inputFile = flag.String("f", "input.txt", "test input")
)

type elf struct {
	food  []int
	total int
}

//one item per line
//number of calories per line
//each elf separates inventory from previous elf by a blank line
//find the elf carrying the most calories and how many is that?

func main() {
	flag.Parse()
	fmt.Println("Processing file:", string(*inputFile))
	var lines []string
	lines, err := readInputFile(*inputFile)
	if err != nil {
		fmt.Println(err)
	}
	elves := []elf{}
	elves, err = parseElves(lines)
	if err != nil {
		fmt.Println(err)
	}

	printElfMax(elves)
	printElfTop3(elves)
}

func readInputFile(filename string) ([]string, error) {
	readFile, err := os.Open(filename)
	var lines []string
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		lines = append(lines, currentLine)
	}
	readFile.Close()
	return lines, nil
}

func sumArray(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func newElf() *elf {
	e := elf{}
	return &e
}

func parseElves(lines []string) ([]elf, error) {
	elves := []elf{}
	e := newElf()
	for i, currentLine := range lines {
		if currentLine == "" {
			elves = append(elves, *e)
			e = newElf()
			continue
		}

		if currentLine != "" {
			n, err := strconv.Atoi(currentLine)
			if err != nil {
				fmt.Println(err)
			}
			e.food = append(e.food, n)
			//could move to find max
			e.total = sumArray(e.food)
		}

		if i == (len(lines)-1){
			elves = append(elves, *e)
		}
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].total < elves[j].total
	})
	return elves, nil
}

func printElfMax(elves []elf) {
	max := elves[len(elves)-1].total
	fmt.Println("****************")
	fmt.Printf("Max number of calories: %d\n", max)
}

func printElfTop3(elves []elf) {
	fmt.Println("****************")
	fmt.Println("Top 3")
	top_three := len(elves) - 3
	total := 0
	for _, v := range elves[top_three:len(elves)] {
		fmt.Println(v.total)
		total += v.total
	}
	fmt.Printf("total top 3: %d\n", total)
}
