package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	inputFile = flag.String("f", "input2.txt", "test input")
)

var numMap = []map[string]string{
	{"name": "twone", "value": "21"},
	{"name": "oneight", "value": "18"},
	{"name": "eightwo", "value": "82"},
	{"name": "eighthree", "value": "83"},
	{"name": "twone", "value": "21"},
	{"name": "oneight", "value": "18"},
	{"name": "threeight", "value": "38"},
	{"name": "fiveight", "value": "58"},
	{"name": "one", "value": "1"},
	{"name": "two", "value": "2"},
	{"name": "three", "value": "3"},
	{"name": "four", "value": "4"},
	{"name": "five", "value": "5"},
	{"name": "six", "value": "6"},
	{"name": "seven", "value": "7"},
	{"name": "eight", "value": "8"},
	{"name": "nine", "value": "9"},
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

func replaceSpelledNumbers(input string) string {
	for _, s := range numMap {

		//fmt.Println("replaceSpelledNumbs: %s, %s", s["name"], s["value"])
		input = strings.ReplaceAll(input, s["name"], s["value"])
	}
	return input
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
	//	fmt.Printf("%d\n", foundInts)
	fmt.Printf("pair: %s = %d+%d\n", pair, foundInts[0], foundInts[len(foundInts)-1])
	//	fmt.Printf("%d\n", converted)
	return converted
}

func process(lines []string) {
	var result []int
	//	fmt.Println(lines)
	for _, currentLine := range lines {
		fmt.Printf("%s\n", currentLine)
		currentLine = replaceSpelledNumbers(currentLine)
		chars := strToCharArray(currentLine)
		result = append(result, findCalibration(chars))
		//result = append(result, 0)
		fmt.Printf("%s\n", currentLine)
		//fmt.Printf("%d\n", result)
		fmt.Println("-------")
	}

	var total int
	for _, n := range result {
		total += n
	}

	fmt.Printf("Result: %d\n", total)
}
