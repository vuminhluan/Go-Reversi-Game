package main

import "fmt"

var DIRECTIONS = [8][2]int{{0,-1}, {-1,-1}, {-1,0}, {-1,1}, {0,1}, {1,1}, {1,0}, {1,-1}}

var board = [8][8]string{}

func initBoard()  {
	board[3][3] = "B"
	board[3][4] = "W"
	board[4][3] = "W"
	board[4][4] = "B"
}

func printBoard()  {
	for _, row := range board  {
		for _, col := range row {
			if col == "" {
				fmt.Print("[ . ]")
			} else {
				fmt.Printf("[ %s ]", col)
			}
		}
		fmt.Println()
	}
}

func getPlayerOwnerCeils(turn string) [][]int  {
	ceils := [][]int{}
	for i, row := range board  {
		for j, col := range row {
			if col == turn {
				ceils = append(ceils, []int{i,j})
			}
		}
	}
	return ceils
}

func revertPlayer(player string) string  {
	if player == "B" {
		return "W"
	} else if player == "W" {
		return "B"
	}
	return ""
}

func getAvailableCeils(turn string) [][]int  {
	var availableCeils = [][]int{}
	existedCeils := getPlayerOwnerCeils(turn)
	for _, coordinate := range existedCeils {
		for _, direction := range DIRECTIONS {

		}
	}
	return availableCeils
}

func init() {
	initBoard()
	//printBoard()
}

func main() {
	getAvailableCeils("B")
	//fmt.Print(getAvailableCeils("C"))
}
