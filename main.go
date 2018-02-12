package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Create new input stream from console
	board := createBoard()
	clearConsole()
	showBoard(board)

	currentPlayer := 1

	for {
		// Get row
		fmt.Println("Player", currentPlayer, "it's your move! In which row will to play?")
		row, err := getInput(reader, board, "That's not a valid row")
		if err {
			continue
		}

		// Get column
		showBoard(board)
		fmt.Println("Now tell me the column")
		column, err := getInput(reader, board, "That's not a valid column")
		if err {
			continue
		}

		// Check if target is already occupied by X or O
		target := &board[row-1][column-1]
		if *target != " " {
			showBoard(board)
			fmt.Println("The location is already occupied")
			continue
		}

		// Set X/O
		marker := getMarker(&currentPlayer)
		board[row-1][column-1] = marker

		// Prepare oponent move
		updatePlayer(&currentPlayer)
		showBoard(board)
	}
}

func createBoard() (board [3][3]string) {
	board = [3][3]string{} // Create new 3x3 Array which represents the board

	// Set all positions on the board to a default value of '-'
	for x, ySlice := range board {
		for y := range ySlice {
			board[x][y] = " "
		}
	}
	return
}

func getMarker(currentPlayer *int) (marker string) {
	switch *currentPlayer {
	case 1:
		marker = "X"
	default:
		marker = "O"
	}
	return
}

func getInput(reader *bufio.Reader, board [3][3]string, errorMsg string) (selected int64, err bool) {
	line, _, _ := reader.ReadLine()
	selected, _ = strconv.ParseInt(string(line), 10, 64)
	if selected < 1 || selected > 3 {
		showBoard(board)
		fmt.Println(errorMsg)
		err = true
	}
	return
}

func showBoard(board [3][3]string) {
	clearConsole() // Clear console before showing the board for a better overview

	prettyBoard := [][]string{{board[0][0], "║", board[0][1], "║", board[0][2]},
		{"═", "╬", "═", "╬", "═"},
		{board[1][0], "║", board[1][1], "║", board[1][2]},
		{"═", "╬", "═", "╬", "═"},
		{board[2][0], "║", board[2][1], "║", board[2][2]}}

	for i := range prettyBoard {
		fmt.Printf("%s\n", strings.Join(prettyBoard[i], ""))
	}
}

func updatePlayer(currentPlayer *int) {
	switch *currentPlayer {
	case 1:
		*currentPlayer = 2
	default:
		*currentPlayer = 1
	}
}

func clearConsole() {
	var empty string
	for i := 0; i < 50; i++ {
		empty += "\n"
	}
	fmt.Print(empty) // Insert 50 empty lines into the console
}
