package player

import "github.com/WilliamAkaWill/tic-tac-toe/shared"

type (
	Player interface {
		GetName() string
		GetMove(board [][]string) (int, error)
		GetPlayerMark() shared.Player
	}
)