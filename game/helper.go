package game

import "github.com/WilliamAkaWill/tic-tac-toe/shared"

func initBoard() [][]string {
	return [][]string{
		{string(shared.Empty), string(shared.Empty), string(shared.Empty)},
		{string(shared.Empty), string(shared.Empty), string(shared.Empty)},
		{string(shared.Empty), string(shared.Empty), string(shared.Empty)},
	}
}