package main

import (
	"strconv"
	"strings"
)



func HowManyGamesPossible(inputData []string) int{
	sum := 0;

	totalObjects := make(map[string]int)

	for _,line := range inputData{
		metadata := strings.Split(line, ":")

		totalObjects["red"] = 0
		totalObjects["blue"] = 0
		totalObjects["green"] = 0
		individualGames := strings.Split(metadata[1], ";")
		for _,game := range individualGames{
			println(game)
			gameValues := getCubesFromGame(game)

			if gameValues["red"] > totalObjects["red"] {
				totalObjects["red"] = gameValues["red"]
			}
			if gameValues["blue"] > totalObjects["blue"] {
				totalObjects["blue"] = gameValues["blue"]
			}
			if gameValues["green"] > totalObjects["green"] {
				totalObjects["green"] = gameValues["green"]
			}

		}
		sum += totalObjects["red"] * totalObjects["green"] * totalObjects["blue"]


	}

	return sum

}

func getCubesFromGame(game string) map[string]int{
	gameMap := make(map[string]int)
	
	shownPieces := strings.Split(game,",")
	for _,cubeType := range shownPieces{
		println("cubetype" + cubeType)
		trimmedCube := strings.Trim(cubeType, " ")
		colorway := strings.Split(trimmedCube, " ")
		trimmedNumber := strings.Trim(colorway[0]," ")
		trimmedColor := strings.Trim(colorway[1], " ")
		println("[" + trimmedColor + "]")
		value,_ := strconv.Atoi(trimmedNumber)
		gameMap[trimmedColor] = value
		

	}
	return gameMap
}
