package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	inputFile         = flag.String("f", "input.txt", "test input")
	inputFile2 string = "input2.txt"
)

func main() {
	flag.Parse()
	fmt.Println("Processing file:", string(*inputFile))
	var lines []string
	lines, err := readInputFile(*inputFile)
	if err != nil {
		fmt.Println(err)
	}
	solutionOne(lines)
	fmt.Println("**************")
	var lines2 []string
	lines2, err = readInputFile(inputFile2)
	if err != nil {
		fmt.Println(err)
	}
	solutionTwo(lines2)
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

func findIdentifier(a string, b string) (string, error) {

	var str = strings.Split(a, "")
	for i := 0; i < len(str); i++ {
		if strings.Contains(b, str[i]) {
			return str[i], nil
		}
	}
	return "", errors.New("empty name")
}

func findCommonChars(a string, b string) string {

	var s = ""
	var str = strings.Split(a, "")
	for i := 0; i < len(str); i++ {
		if strings.Contains(b, str[i]) {
			s += str[i]
		}
	}
	return s
}

func calcPriority(a string) int {
	var smallA byte = 'a'   //97
	var capitolA byte = 'A' //65
	bigAValue := int(capitolA)
	found := []byte(a)
	asciiNum := int(found[0])

	if asciiNum >= 97 {
		return int(asciiNum) - int(smallA) + 1
	} else {
		return int(asciiNum) - int(bigAValue) + 27
	}

}

func solutionOne(lines []string) {
	total := 0
	for _, currentLine := range lines {
		compartment1 := currentLine[0:(len(currentLine) / 2)]
		compartment2 := currentLine[(len(currentLine) / 2):len(currentLine)]
		foundChar, err := findIdentifier(compartment1, compartment2)
		if err != nil {
			errors.New("character not found")
		}
		total += calcPriority(foundChar)
	}
	fmt.Printf("solution one total: %d\n", total)
}

func solutionTwo(lines []string) {
	fmt.Println(len(lines))

	var groups []string
	numGroups := int(math.Ceil(float64(len(lines)) / 3))
	fmt.Printf("groups %d numGroups %d\n", groups, numGroups)
	offset := 0
	total := 0
	for i := 0; i < numGroups; i++ {
		offset = i * 3

		fmt.Printf("********** group %d offset %d **********\n", i, offset)
		lineGroup := lines[offset:(offset + 3)]
		priority := 0
		//assumes input is always groups of 3
		commonChars := findCommonChars(lineGroup[0], lineGroup[1])
		commonChars2 := findCommonChars(lineGroup[1], lineGroup[2])
		foundChar, _ := findIdentifier(commonChars, commonChars2)

		fmt.Println(commonChars + " " + commonChars2 + " " + foundChar)
		priority = calcPriority(foundChar)

		fmt.Println(lineGroup[0])
		fmt.Println(lineGroup[1])
		fmt.Println(lineGroup[2])
		fmt.Println(priority)
		total += priority

	}
	fmt.Println(total)
}
