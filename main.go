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
	currentPlayer := 1

	clearConsole()
	fmt.Println("Which size should the board be?\nA normal game has a size of three")
	line, _, _ := reader.ReadLine()
	boardsize, _ := strconv.ParseInt(string(line), 10, 64)
	if boardsize < 1 {
		fmt.Println("Oops! Seems like thats not possible")
		return
	}

	board := createBoard(int(boardsize)) // Initialize a new gameboard
	showBoard(board)

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

		// Check if the game has a winner
		winner := hasWinner(board)
		if winner != 0 {
			showBoard(board)
			fmt.Println("Player", winner, "won!")
			return
		}

		if isDraw(board) {
			showBoard(board)
			fmt.Println("It's a draw!")
			return
		}

		// Prepare opponent move
		updatePlayer(&currentPlayer)
		showBoard(board)
	}
}

func createBoard(boardsize int) (board [][]string) {
	board = [][]string{} // Create new 3x3 Array which represents the board

	// Set all positions on the board to a default value of '-'
	for x := 0; x < boardsize; x++ {
		board = append(board, []string{})
		for y := 0; y < boardsize; y++ {
			board[x] = append(board[x], " ")
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

func getInput(reader *bufio.Reader, board [][]string, errorMsg string) (selected int64, err bool) {
	line, _, _ := reader.ReadLine()
	selected, _ = strconv.ParseInt(string(line), 10, 64)
	if selected < 1 || selected > int64(len(board)) {
		showBoard(board)
		fmt.Println(errorMsg)
		err = true
	}
	return
}

func showBoard(board [][]string) {
	clearConsole() // Clear console before showing the board for a better overview

	prettyBoard := getPrettyBoard(board)

	for i := range prettyBoard {
		fmt.Printf("%s\n", strings.Join(prettyBoard[i], ""))
	}
}

func clearConsole() {
	var empty string
	for i := 0; i < 50; i++ {
		empty += "\n"
	}
	fmt.Print(empty) // Insert 50 empty lines into the console
}

func getPrettyBoard(board [][]string) (prettyBoard [][]string) {
	prettyBoard = append(prettyBoard, getPrettyRow(board[0]))

	for x := 1; x < len(board); x++ {
		prettyBoard = append(prettyBoard, getPrettyDelimiter(len(board[x])))
		prettyBoard = append(prettyBoard, getPrettyRow(board[x]))
	}
	return
}

func getPrettyRow(row []string) (prettyRow []string) {
	prettyRow = append(prettyRow, row[0])
	for i := 1; i < len(row); i++ {
		prettyRow = append(prettyRow, "║", row[i])
	}
	return
}

func getPrettyDelimiter(columns int) (prettyDelimiter []string) {
	prettyDelimiter = append(prettyDelimiter, "═")
	for i := 1; i < columns; i++ {
		prettyDelimiter = append(prettyDelimiter, "╬", "═")
	}
	return
}

func updatePlayer(currentPlayer *int) {
	switch *currentPlayer {
	case 1:
		*currentPlayer = 2
	default:
		*currentPlayer = 1
	}
}

func hasWinner(board [][]string) (winner int) {
	potWinner := make([]int, 0)
	potWinner = append(potWinner, hasWinnerHorizontal(board))
	potWinner = append(potWinner, hasWinnerVertical(board))
	potWinner = append(potWinner, hasWinnerDiagonal(board))
	for _, winner := range potWinner {
		if winner != 0 {
			return winner
		}
	}
	return 0
}

func hasWinnerHorizontal(board [][]string) (winner int) {
Loop:
	for _, row := range board {
		winner = resolveMarker(row[0]) // We asume that whoever has a marker in the first spot of the row won
		// If there is a marker which is different to the first one, we know that nobody won the row
		for _, marker := range row {
			if winner != resolveMarker(marker) {
				continue Loop
			}
		}
		return winner
	}
	return 0
}

func hasWinnerVertical(board [][]string) (winner int) {
Loop:
	for column := range board[0] {
		winner = resolveMarker(board[0][column]) // We asume that whoever has a marker in the first spot of the column won
		// If there is a marker which is different to the first one, we know that nobody won the column
		for _, marker := range board {
			if winner != resolveMarker(marker[column]) {
				continue Loop
			}
		}
		return winner
	}
	return 0
}

func hasWinnerDiagonal(board [][]string) (winner int) {
	// Check from top left to bottom right
	winner = resolveMarker(board[0][0])
	for i := range board {
		if winner != resolveMarker(board[i][i]) {
			winner = 0
			break
		}
	}

	if winner != 0 {
		return
	}

	// Check from top right to bottom left
	winner = resolveMarker(board[0][len(board[0])-1])
	for i := range board {
		if winner != resolveMarker(board[len(board[0])-1-i][i]) {
			winner = 0
			break
		}
	}
	return
}

func resolveMarker(marker string) (player int) {
	switch marker {
	case "X":
		return 1
	case "O":
		return 2
	default:
		return 0
	}
}

func isDraw(board [][]string) (draw bool) {
	draw = true

	for _, x := range board {
		for _, y := range x {
			if y == " " {
				draw = false
				return
			}
		}
	}
	return
}
