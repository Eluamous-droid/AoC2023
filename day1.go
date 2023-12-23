package main

import (
	"strconv"
	"strings"
)

func GetSumOfCalibrationNumbers(fileLines []string) int{
	
	sum := 0;
	for _,line := range fileLines{
		digits := findFirstDigit(line) + findLastDigit(line)
		var number int
		number,_ = strconv.Atoi(digits)

		sum += number
	}
	return sum

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
