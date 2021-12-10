package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	//"io"
)

func main() {

	depths, err := readFile("input.txt")
	if err != nil {
		log.Fatalf("readlines: %s", err)
	}
	num := countDepthIncreases(depths)
	fmt.Println("countDepthIncreases:", num)

	num = slidingWindowIncreases(depths)
	fmt.Println("slidingWindowIncreases:", num)
}

func countDepthIncreases(depths []int) int {
	lastLine := 0
	num := 0
	for i, line := range depths {
		if i == 0 {
			lastLine = line
			continue
		}
		if line > lastLine {
			num++
		}
		lastLine = line
	}
	return num
}

func sum(s []int) int{
	total := 0
	for _, v := range s{
		total += v
	}
	return total
}

func slidingWindowIncreases(depths []int) int{
	windows := make(map[int]int)
	num := 0
	count :=0
	prevWindowTotal := 0

	for i, _ := range depths{
		if( (i+3) < len(depths)){
			s := depths[i:i+3]
			totalWindow := sum(s)
			windows[count] = totalWindow
			if(i == 0){
				fmt.Printf("window:%s total:%d (NA) \n", s, totalWindow)
			}
			if(totalWindow > prevWindowTotal){
				num ++
				fmt.Printf("window:%s total:%d increased ++ \n", s, totalWindow)
			} else {
				fmt.Printf("window:%s total:%d decreased \n", s, totalWindow)
			}

			prevWindowTotal = totalWindow
		}
	}
	fmt.Println(windows)
	return num
}

func readFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, l)
	}
	return lines, scanner.Err()
}
