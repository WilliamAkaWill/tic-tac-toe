package validate

import (
	"github.com/WilliamAkaWill/tic-tac-toe/errors"
	"github.com/WilliamAkaWill/tic-tac-toe/shared"
)

var winningCombinations = [][]coordinate{
	// Rows
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	// Columns
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
	// Diagonals
	{{0, 0}, {1, 1}, {2, 2}},
	{{0, 2}, {1, 1}, {2, 0}},
}

func CheckBoard(board [][]string) shared.State {
	for _, combo := range winningCombinations {
		mark1 := board[combo[0].row][combo[0].col]
		mark2 := board[combo[1].row][combo[1].col]
		mark3 := board[combo[2].row][combo[2].col]
		if mark1 != string(shared.Empty) && mark1 == mark2 && mark2 == mark3 {
			if mark1 == string(shared.PlayerX) {
				return shared.Player1WON
			}
			return shared.Player2WON
		}
	}

	// Check for tie or in-progress
	for _, row := range board {
		for _, cell := range row {
			if cell == string(shared.Empty) {
				return shared.InProgress
			}
		}
	}
	return shared.Tie
}

func ValidateAndGetCoordinate(board [][]string, input int) (row int, col int, err error) {
	if input < 1 || input > 9 {
		return 0, 0, errors.ErrOutOfBounds
	}

	row, col = shared.ConvertToCoordinates(input)
	if board[row][col] != string(shared.Empty) {
		return 0, 0, errors.ErrInvalidMove
	}

	return row, col, nil
}
