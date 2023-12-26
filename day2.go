package main

import (
	"strconv"
	"strings"
)



func HowManyGamesPossible(inputData []string) int{
	sum := 0;

	totalObjects := make(map[string]int)
	totalObjects["red"] = 12
	totalObjects["blue"] = 14
	totalObjects["green"] = 13

	for _,line := range inputData{
		metadata := strings.Split(line, ":")
		gameNumber := strings.Split(metadata[0], " ")

		impossibleGame := true 
		individualGames := strings.Split(metadata[1], ";")
		for _,game := range individualGames{
			println(game)
			gameValues := getCubesFromGame(game)

			if gameValues["red"] > totalObjects["red"] {
				impossibleGame = false 
				break
			}
			if gameValues["blue"] > totalObjects["blue"] {
				impossibleGame = false
				break
			}
			if gameValues["green"] > totalObjects["green"] {
				impossibleGame = false
				break
			}

		}

			if impossibleGame{
				gameValue,_ :=strconv.Atoi(gameNumber[1])
				sum += gameValue 
			}

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
