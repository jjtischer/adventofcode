package main

import (
	"bufio"
	//"errors"
	"flag"
	"fmt"
//	"math"
	"os"
	"strconv"
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
		str := strings.Split(currentLine, ",")
		elf1 := strings.Split(str[0],"-")
		e1start,_:= strconv.Atoi(elf1[0])
		e1end,_:= strconv.Atoi(elf1[1])

		elf2 := strings.Split(str[1],"-")
		e2start,_:= strconv.Atoi(elf2[0])
		e2end,_:= strconv.Atoi(elf2[1])

		if e1start <=e2start && e1end >= e2end{
			total+=1

		} else if e2start <=e1start && e2end >= e1end{
			total+=1
		}
	}
	fmt.Printf("solution one total: %d\n", total)
}

func solutionTwo(lines []string) {
	total:=0

	for _, currentLine := range lines {
		bucketA := make(map[string]int)
		str := strings.Split(currentLine, ",")
		elf1 := strings.Split(str[0],"-")
		e1start,_:= strconv.Atoi(elf1[0])
		e1end,_:= strconv.Atoi(elf1[1])
		for i:=e1start; i <= e1end; i++{
			bucketA[strconv.Itoa(i)]+=1
		}

		elf2 := strings.Split(str[1],"-")
		e2start,_:= strconv.Atoi(elf2[0])
		e2end,_:= strconv.Atoi(elf2[1])
		for i:=e2start; i <= e2end; i++{
			bucketA[strconv.Itoa(i)]+=1
		}

		for i:= range bucketA{
			if bucketA[i] >1{
				total+=1
				break
			}
		}
	}
	fmt.Printf("solution two total: %d\n", total)
}
