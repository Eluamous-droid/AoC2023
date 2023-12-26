package main

import (
	"bufio"
	"os"
)

func main(){
	fileLines := getFileInput(os.Args[1])
	output := HowManyGamesPossible(fileLines);
	println(output)
}

func writeOutputToFile(output []string){
	f, err := os.Create("output.txt")
	check(err)
	defer f.Close()

	for _,line := range output{
	f.WriteString(line + "\n")
	}
	f.Sync()
}

func getFileInput(input string) []string{

	file, err := os.Open(input)

	check(err)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan(){
		fileLines = append(fileLines, fileScanner.Text())
	}
	file.Close()

	return fileLines
}


func check(e error){
	if e != nil{
		panic(e)
	}
}
