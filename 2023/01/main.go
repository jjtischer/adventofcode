package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	inputFile = flag.String("f", "input.txt", "test input")
)

var numMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	flag.Parse()
	fmt.Println("Processing file:", string(*inputFile))
	var lines []string
	lines, err := readInputFile(*inputFile)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(lines)
	process(lines)
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

func strToCharArray(input string) []rune {
	return []rune(input)
}

func findCalibration(chars []rune) int {
	var foundInts []int
	var pair string

	for _, c := range chars {
		if val, err := strconv.Atoi(string(c)); err == nil {
			foundInts = append(foundInts, val)
		}
	}
	pair = strconv.Itoa(foundInts[0]) + strconv.Itoa(foundInts[len(foundInts)-1])
	converted, err := strconv.Atoi(pair)
	if err != nil {
		fmt.Errorf("findCalibraion error: %v", err)
	}
	fmt.Printf("%d\n", foundInts)
	fmt.Printf("pair: %s = %d+%d\n", pair, foundInts[0], foundInts[len(foundInts)-1])
	//	fmt.Printf("%d\n", converted)
	return converted
}

func process(lines []string) {
	var result []int
	fmt.Println(lines)
	for _, currentLine := range lines {
		chars := strToCharArray(currentLine)
		result = append(result, findCalibration(chars))
		fmt.Printf("%d\n", result)
		fmt.Println("-------")
	}

	var total int
	for _, n := range result {
		total += n
	}

	fmt.Printf("Result: %d\n", total)
}
