package main

import (
	"slices"
	"strconv"
	"strings"
)

func GetPointsFromScratches(lines []string) int {

	total := 0
	copies := make([]int, len(lines))
	for i,line := range lines{
		copies[i]++

		meta := strings.Split(line,":")
		games := strings.Split(meta[1], "|")
		winningNumbers := getNumbers(games[0])
		scratchNumbers := getNumbers(games[1])
		correctNumbers := 0
		for _,number := range winningNumbers{
			if slices.Contains(scratchNumbers, number){
				correctNumbers++
				copies[i+correctNumbers] += copies[i]
			}
		}
		total += copies[i]	
		

	}
	return total
}


func getNumbers(numbers string) []int{
	var actualNumbers []int
	charSlice := strings.Split(numbers, "")
	for i := 1; i < len(numbers) - 1; i += 3{
		number, err := strconv.Atoi(strings.Trim(charSlice[i] + charSlice[i+1], " "))
		if err != nil{panic(err)}
		actualNumbers = append(actualNumbers, number)
	}
	return actualNumbers

}

func multiplyPoints(point int) int{
	if point == 0{
		return 1
	}
	return point * 2
}
