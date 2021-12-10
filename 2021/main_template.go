package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	inputFile = flag.String("f", "input.txt","test input")
)

const (
	slidingWindowSize = 3
)

func main() {
	flag.Parse()
	filename := string(*inputFile)
	fmt.Println("processing file: ", filename)

	values, err := readFile(filename)
	if( err != nil){
		log.Fatalf("readlines: %s", err)
	}

	ret1 := problem1(values)
	fmt.Println("solution1:", ret1)

	ret2 := problem2(values)
	fmt.Println("solution2:", ret2)

}

func problem1(data []int) int {

	return 0
}

func problem2(data []int) int {

	return 0
}

func readFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, l)
	}
	return lines, scanner.Err()
}
