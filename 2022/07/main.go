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
type item struct {
	name string
	size int
}

func newItem() *item {
	e := item{}
	return &e
}

func newDir() *item {
	e := item{}
	return &e
}

func solutionOne(lines []string) {
	total :=0
	//elves := []elf{}
	//e := newElf()
	currentPath := ""
	fileStruct := make(map[string][]item)
	fileSize := make(map[string]int)
	for _, currentLine := range lines {
		//fmt.Printf("*************************************** Current Line: %s \n",currentLine)
		lastIdx := strings.LastIndex(currentLine, " ")
		cmd := currentLine[0:lastIdx]
		value := strings.Trim(currentLine[lastIdx:len(currentLine)]," ")
		switch cmd {
		case "$ cd":
			if value == "/"{
				currentPath = "/"
			} else if value == ".."{
				tmpIdx := strings.LastIndex(currentPath,"/")
				tmpPath := currentPath[0:tmpIdx]
				if tmpPath == ""{
					tmpPath = "/"
				}
				//fmt.Println("cd .. WTF UP DIR "+ tmpPath)
				currentPath = tmpPath
			} else {
				if currentPath == "/"{
					currentPath += value
				}else {
					currentPath += "/"+value
				}

			}
			fmt.Println(currentLine + " \t\t\t "+  currentPath)
			//fmt.Println("currentPath: "+ currentPath)
		case "$":
			if value == "ls"{
				fmt.Printf("%s \t %s %d\n",currentLine, currentPath, fileSize[currentPath])
				for i, _ := range fileStruct[currentPath]{
					fmt.Printf("\t%s \t\t\t %d\n",fileStruct[currentPath][i].name,fileStruct[currentPath][i].size)
				}

				fmt.Println("___________________________________________")
			}

		case "dir":
			//fmt.Println("dir  "+ fmt.Sprint(fileStruct[value]))
			f := newItem()
			f.size = 0
			f.name = value
			d := ""
			if currentPath == "/"{
				d = currentPath+value
			}else {
				d = currentPath+"/"+value
			}

			//fmt.Println(">>>>************* ADDING DIR: ", d)
			fileStruct[currentPath] = append(fileStruct[d], *f)
			fmt.Println(currentLine + " \t\t\t "+  d)
		default:

			fmt.Println("ADDING FILE: >>>> ", currentLine," to ", currentPath)
			p := strings.Split(currentLine," ")
			stupidInt,_ :=  strconv.Atoi(p[0])

			f := newItem()
			f.size = stupidInt
			f.name = p[1]
			//keep track of file size in a directory; does not keep track of children
			fileSize[currentPath] += stupidInt
			fileStruct[currentPath] = append(fileStruct[currentPath], *f)
			fmt.Printf("\t%s %s %d\n", currentPath, f.name , fileSize[currentPath])
			fmt.Println("___________________________________________")
			//fmt.Printf(" added: %s ---- size: %s ---- path:%s total: %d\n",p[1], p[0], currentPath, fileSize[currentPath])
		}

	}



	//fmt.Println("********************************************** FILESIZE")
	//fmt.Println(fileSize)
	//fmt.Println("********************************************** RESULTS")
	////fmt.Println(fileStruct)
	//for k, _ := range fileStruct{
	//
	//	fmt.Println(fileStruct[k])
	//	fmt.Println(fileSize[k])
	//	//fmt.Printf("key[%s] value[%s] %d\n", k, k[v].name, k[v].size)
	//}
	//fmt.Printf("solution one total: %d\n", total)

	fmt.Println("********************************************** TOTALS WTF \"**********************************************")

	totalByDir := make(map[string]int)
	for k, v := range fileSize {
		fmt.Printf("Searching for %s %d %d\n", k, v, totalByDir[k])
		for k2, v2 := range fileSize {

			if strings.Contains(k2, k){

				totalByDir[k] += v2
				//fmt.Printf("Searching %s in %s %d %d\n", k, k2, v, totalByDir[k])
				fmt.Printf("Adding %d == %d\n",v2, totalByDir[k])
			}
		}
		//fmt.Printf("key[%s] value[%d] %d\n", k, v, totalByDir[k])
	}
	fmt.Println("MAP totalByDir")
	fmt.Println(totalByDir)
	fmt.Println(total)

	fmt.Println("********************************************** TOTALS <= 100000 **********************************************")
	total =0
	for k, v := range totalByDir{
		if v <= 100000{
			total+=v
			fmt.Printf("ADDED key[%s] value[%d]  total %d \n", k, v, total)
		} else {
			fmt.Printf("SKIP key[%s] value[%d] \n", k, v)
		}

	}
	fmt.Println("total")
	//1495837 is to low
	fmt.Println(total)
}


func solutionTwo(lines []string) {
	total := 0
	for _, currentLine := range lines {
		fmt.Println(len(currentLine))

	}
	fmt.Printf("solutionTwo total: %d\n", total)
}
