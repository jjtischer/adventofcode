package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	//"strings"
)

var (
	inputFile = flag.String("f", "input.txt", "test input")
)

//The score for a single round is the score for the shape you selected
//(1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round
//(0 if you lost, 3 if the round was a draw, and 6 if you won
func main() {
	flag.Parse()
	fmt.Println("Processing file:", string(*inputFile))
	var lines []string
	lines, err := readInputFile(*inputFile)
	if err != nil {
		fmt.Println(err)
	}
	solutionOne(lines)
	solutionTwo(lines)

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

func calcPoints(a string) int{

	m := make(map[string]int)
	//rock
	m["A X"] = 3
	m["A Y"] = 6
	m["A Z"] = 0
	//paper
	m["B X"] = 0
	m["B Y"] = 3
	m["B Z"] = 6
	//sissors
	m["C X"] = 6
	m["C Y"] = 0
	m["C Z"] = 3

	return m[a]
}

func solutionOne(lines []string) {
	m := make(map[string]int)
	m["X"]=1  //"rock"
	m["Y"]=2  //"paper"
	m["Z"]=3  //"sissors"

	total :=0
	for _, currentLine := range lines {
		s := strings.Split(currentLine," ")
		total += m[s[1]]
		total += calcPoints(currentLine)
		//fmt.Printf("currentLine %s - %d points %d\n", currentLine, m[s[1]], calcPoints(currentLine))
	}

	fmt.Printf("Solution one:%d\n",total)
}

//translates the strategy in col 2 to get the outcome
func getShape(a string) string{

	m := make(map[string]string)
	//rock
	m["A X"] = "Z"
	m["A Y"] = "X"
	m["A Z"] = "Y"
	//paper
	m["B X"] = "X"
	m["B Y"] = "Y"
	m["B Z"] = "Z"
	//sissors
	m["C X"] = "Y"
	m["C Y"] = "Z"
	m["C Z"] = "X"

	return m[a]
}

func solutionTwo(lines []string) {
	shapesValue := make(map[string]int)
	shapesValue["X"]=1 //"rock"
	shapesValue["Y"]=2 //"paper"
	shapesValue["Z"]=3 //"sissors"

	score := make(map[string]int)
	score["X"] = 0
	score["Y"] = 3
	score["Z"] = 6

	total :=0
	for _, currentLine := range lines {
		s := strings.Split(currentLine," ")
		//opponent := s[0]
		result := s[1]
		me := getShape(currentLine)
		total += shapesValue[me] + score[result]

		//fmt.Printf("%s opponent %s vs me %s score %d\n", currentLine, opponent, result, (shapesValue[me] + winnerScore(result)))
	}
	fmt.Printf("Solution two:%d\n",total)
}
