package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func Make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		matrix[i] = rows[startRow : startRow+m]
	}
	return matrix
}


func printMatrix(m [][]string, h int, w int){
	total :=0
	for i:=0;i<h;i++{
		for j:=0;j<w;j++{
			if m[i][j] == ""{
				fmt.Printf(".")
			}else{
				fmt.Printf("%s", m[i][j])
			}

			if m[i][j] =="T"{
				total+=1
			}
		}
		fmt.Println("\n")
	}
	//i dont feel like making a case for T overwriting S
	total-=1
	fmt.Printf("total %d\n", total)
}


func solutionOne(lines []string) {
	total := 0
	maxDir := make(map[string]int)

	//m[1][2] = 1
	//fmt.Printf("%#v\n", m)

	//find max dimensions of matrix
	for _, currentLine := range lines {
		fmt.Println(len(currentLine))
		s := strings.Split(currentLine," ")
		cmd := s[0]
		amt,_ := strconv.Atoi(s[1])
		if maxDir[cmd]< amt{
			maxDir[cmd] = amt
		}
	}
	width := maxDir["L"]+ maxDir["R"]
	height := maxDir["D"]+ maxDir["U"]
	fmt.Printf("width %d height %d\n",width,height)
	fmt.Println(maxDir)
	m := Make2D[string](height,width)


	//assume starting at the bottom left
	startH := height-1
	startW := 0
	headH := startH
	headW := startW
	tailH := startH
	tailW := startW

	m[startH][startW]="s"
	fmt.Printf("%#v\n", m)
	//printMatrix(m, height, width)

	for _, currentLine := range lines {
		s := strings.Split(currentLine," ")
		cmd := s[0]
		amt,_ := strconv.Atoi(s[1])
		if maxDir[cmd]< amt{
			maxDir[cmd] = amt
		}

		switch cmd {
			case "U":
				headH -= amt
			case "D":
				headH += amt
			case "L":
				headW -= amt
			case "R":
				headW += amt
		}
		fmt.Printf("headH %d headW %d \n", headH,headW)
		m[headH][headW]="H"
		//left and right
		if tailH == headH && tailW != headW && amt > 1{
			if tailW < headW{
				for i:=tailW;i < headW;i++{
					tailW = i
					m[tailH][tailW]="T"
				}
			} else {
				for i:=tailW;i > headW;i--{
					tailW = i
					m[tailH][tailW]="T"
				}
			}
		}

		//up and down
		if tailW == headW && tailH != headH && amt >1{
			if tailH < headH{
				for i:=tailH;i < headH;i++{
					tailH = i
					m[tailH][tailW]="T"
				}
			} else {
				for i:=tailH;i > headH;i--{
					tailH = i
					m[tailH][tailW]="T"
				}
			}
		}

		if math.Abs(float64(tailW) - float64(headW)) >= 1 && math.Abs(float64(tailH) - float64(headH))>=1{

			if tailW<=headW{
				tailW +=1
			} else{
				tailW -=1
			}
			if tailH<=headH{
				tailH +=1
			} else{
				tailH -=1
			}

			//move T into its position
			m[tailH][tailW]="T"
			//loop through each position until it arrives close to H
			if tailH == headH {
				if tailW < headW{
					for i:=tailW;i < headW;i++{
						tailW = i
						m[tailH][tailW]="T"
					}
				} else {
					for i:=tailW;i > headW;i--{
						tailW = i
						m[tailH][tailW]="T"
					}
				}

			} else if tailW == headW {
				if tailH < headH{
					for i:=tailH;i < headH;i++{
						tailH = i
						m[tailH][tailW]="T"
					}
				} else {
					for i:=tailH;i > headH;i--{
						tailH = i
						m[tailH][tailW]="T"
					}
				}
			}

		}
		//
		//if tailW != headW && tailH != headW{
		//	fmt.Println(math.Abs(float64(tailW) - float64(headW)))
		//	if math.Abs(float64(tailW) - float64(headW))>1 && math.Abs(float64(tailH) - float64(headH))>1{
		//
		//		if tailW < headW && tailH< headH{
		//			tailW +=1
		//			tailH +=1
		//		} else if tailW > headW && tailH< headH {
		//			tailW -=1
		//			tailH +=1
		//		} else if tailW > headW && tailH> headH {
		//			tailW -=1
		//			tailH -=1
		//		} else if tailW < headW && tailH > headH {
		//			tailW +=1
		//			tailH -=1
		//		}
		//		fmt.Printf("%d %d",tailH, tailW)
		//		m[tailH][tailW]="T"
		//
		//	}
		//}

	//	tailH, tailW = calcTail(headH, headW, tailH, tailW)


	}

	printMatrix(m, height, width)
	fmt.Printf("solutionTwo total: %d\n", total)
}

func solutionTwo(lines []string) {
	total := 0
	for _, currentLine := range lines {
		fmt.Println(len(currentLine))

	}
	fmt.Printf("solutionTwo total: %d\n", total)
}
