package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	//"strings"
)

var (
	inputFile string = "input.txt"
	inputFile2 string = "input2.txt"
)

func main() {
	flag.Parse()
	fmt.Println("Processing file:", string(inputFile))
	var lines []string
	lines, err := readInputFile(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	solutionOne(lines)

	fmt.Println("************************************************")
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

func solutionOne(lines []string) {
	total := 0
	for _, currentLine := range lines {
		fmt.Println(len(currentLine))

	}
	fmt.Printf("solutionTwo total: %d\n", total)
}

func solutionTwo(lines []string) {
	total := 0
	for _, currentLine := range lines {
		fmt.Println(len(currentLine))

	}
	fmt.Printf("solutionTwo total: %d\n", total)
}
