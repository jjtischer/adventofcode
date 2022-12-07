package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

func checkUnique(s string,l int) bool{
	m := make(map[string]int)
	for _, v := range s{
		m[string(v)]+=1
	}
	if len(m) == l{
		return true
	}else {
		return false
	}

}

func solutionOne(lines []string) {
	total := 0
	foundMarker := 0

	for _, currentLine := range lines {
		fmt.Println(currentLine)
		markerLen := 4
		a:=""
		foundMarker = 0
		endline:=0
		for i:=0;i< len(currentLine);i++{
			endline = i+markerLen
			if endline <= len(currentLine){
				a = currentLine[i:endline]
				if checkUnique(a, markerLen) == true{
					foundMarker = endline
					break
				}
			}
		}
	}
	fmt.Println(foundMarker)
	fmt.Printf("solution one total: %d\n", total)
}

func solutionTwo(lines []string) {
	total := 0
	foundMarker := 0

	for _, currentLine := range lines {
		fmt.Println(currentLine)
		markerLen := 14
		a:=""
		foundMarker = 0
		endline:=0
		for i:=0;i< len(currentLine);i++{
			endline = i+markerLen
			if endline <= len(currentLine){
				a = currentLine[i:endline]
				if checkUnique(a, markerLen) == true{
					foundMarker = endline
					break
				}
			}
		}
	}
	fmt.Println(foundMarker)
	fmt.Printf("solution one total: %d\n", total)
}
