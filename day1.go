package main

import (
	"strconv"
	"strings"
)

func GetSumOfCalibrationNumbers(fileLines []string) int{
	
	sum := 0;
	for _,line := range fileLines{
		actualLine := replaceWordWithDigit(line)
		digits := findFirstDigit(actualLine) + findLastDigit(actualLine)
		println(digits)
		var number int
		number,_ = strconv.Atoi(digits)

		sum += number
	}
	return sum

}

func replaceWordWithDigit(line string) string{
	newLine := strings.ReplaceAll(line, "one", "o1e")
	newLine = strings.ReplaceAll(newLine, "two", "t2o")
	newLine = strings.ReplaceAll(newLine, "three", "t3ree")
	newLine = strings.ReplaceAll(newLine, "four", "f4or")
	newLine = strings.ReplaceAll(newLine, "five", "f5ve")
	newLine = strings.ReplaceAll(newLine, "six", "s6x")
	newLine = strings.ReplaceAll(newLine, "seven", "s7ven")
	newLine = strings.ReplaceAll(newLine, "eight", "e8ght")
	newLine = strings.ReplaceAll(newLine, "nine", "n9ne")

	return newLine
}

func findFirstDigit(line string) string{
	
		chars := strings.Split(line, "")
		for _,char := range chars{
			if _,err := strconv.Atoi(char); err == nil {
				return char
			}
		}
	return "fake"
}

func findLastDigit(line string) string{
	
		chars := strings.Split(line, "")
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	
		for _,char := range chars{
			if _,err := strconv.Atoi(char); err == nil {
				return char
			}
		}
	return "fake"
}
