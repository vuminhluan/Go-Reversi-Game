package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var DIRECTIONS = [8][2]int{{0,-1}, {-1,-1}, {-1,0}, {-1,1}, {0,1}, {1,1}, {1,0}, {1,-1}}

var BOARD = [8][8]string{}
var TURN string
var AVAILABLE_MOVES = [][]int{}

func initBoard()  {
	BOARD[3][3] = "B"
	BOARD[3][4] = "W"
	BOARD[4][3] = "W"
	BOARD[4][4] = "B"
}

func initTurn(turn string)  {
	TURN = turn
}

func printBoard()  {
	for _, row := range BOARD  {
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
	for i, row := range BOARD  {
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
	return "Error"
}

func isOutOfBoard(x int, y int) bool  {
	if x > len(BOARD) || x < 0 {
		return true
	} else if y > len(BOARD[x]) || y < 0 {
		return true
	}
	return false
}

func setAvailableMoves(turn string)  {
	AVAILABLE_MOVES = nil // empty last AVAILABLE_MOVES
	existedCeils := getPlayerOwnerCeils(TURN)
	for _, coordinate := range existedCeils {
		for _, direction := range DIRECTIONS {
			x := direction[0]
			y := direction[1]
			newX := coordinate[0] + x
			newY := coordinate[1] + y
			count := 0
			for !isOutOfBoard(newX, newY) && BOARD[newX][newY] == revertPlayer(TURN) {
				newX += x
				newY += y
				count++
			}
			if !isOutOfBoard(newX, newY) && BOARD[newX][newY] == "" && count > 0 {
				AVAILABLE_MOVES = append(AVAILABLE_MOVES, []int{newX,newY})
			}
		}
	}
}

func revertChess(moveX int, moveY int)  {
	var revertableChesses = [][]int{}
	revertableChesses = append(revertableChesses, []int{moveX, moveY})
	for _, direction := range DIRECTIONS {
		x := direction[0]
		y := direction[1]
		newX := moveX + x
		newY := moveY + y
		for !isOutOfBoard(newX, newY) && BOARD[newX][newY] == revertPlayer(TURN) {
			revertableChesses = append(revertableChesses, []int{newX, newY})
			newX += x
			newY += y
		}
		// detect if stop point == TURN && len(revertableChess) > 0 => Revert between chess
		if !isOutOfBoard(newX, newY) && BOARD[newX][newY] == TURN && len(revertableChesses) > 0 {
			// revert chess
			for _, chessCoordinates := range revertableChesses {
				BOARD[chessCoordinates[0]][chessCoordinates[1]] = TURN
			}
		}
	}
}

func init() {
	initBoard()
	initTurn("B")
	setAvailableMoves(TURN)
	printBoard()
	fmt.Println()
}

func main() {
	for len(AVAILABLE_MOVES) > 0 {
		fmt.Printf("Available moves for %s: %v \n", TURN, AVAILABLE_MOVES)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter your choose: ")
		scanner.Scan()
		move := scanner.Text()
		moveX, _ := strconv.Atoi(string(move[0]))
		moveY, _ := strconv.Atoi(string(move[1]))
		//fmt.Print(isOutOfBoard(moveX+0, moveY-1))

		revertChess(moveX, moveY)

		TURN = revertPlayer(TURN)
		printBoard()
		setAvailableMoves(TURN)
	}
	fmt.Println()
	fmt.Printf("Winner is: %s", TURN)

}
