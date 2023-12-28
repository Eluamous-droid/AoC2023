package main

import (
	"slices"
	"strconv"
	"strings"
)

func GetPointsFromScratches(lines []string) int {

	sum := 0
	for _,line := range lines{

		meta := strings.Split(line,":")
		games := strings.Split(meta[1], "|")
		winningNumbers := getNumbers(games[0])
		scratchNumbers := getNumbers(games[1])
		points := 0
		for _,number := range winningNumbers{
			println(number)
			if slices.Contains(scratchNumbers, number){
				points = multiplyPoints(points)
			}
		}
		
		println(points)
		sum += points

	}
	return sum
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
