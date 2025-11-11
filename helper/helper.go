package helper

import "github.com/WilliamAkaWill/tic-tac-toe/shared"

func GetAvailableMoves(board [][]string) []int {
	var moves []int
	for i := 0; i < 9; i++ {
		row, col := shared.ConvertToCoordinates(i + 1)
		if board[row][col] == string(shared.Empty) {
			moves = append(moves, i+1) // +1 für 1-basierte Eingabe
		}
	}
	return moves
}