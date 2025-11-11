package player

import (
	"fmt"
	"github.com/WilliamAkaWill/tic-tac-toe/errors"
	"github.com/WilliamAkaWill/tic-tac-toe/shared"
)

type (
	Human struct {
		name            string
		languageService shared.LanguageService
		playerMark      shared.Player
	}
)

func NewHuman(name string, langService shared.LanguageService, playerMark shared.Player) *Human {
	return &Human{
		name:            name,
		languageService: langService,
		playerMark:      playerMark,
	}
}

func (h *Human) GetPlayerMark() shared.Player {
	return h.playerMark
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) GetMove(_ [][]string) (int, error) {
	prompt := h.languageService.GetString(shared.InputTicTacToeNumber)
	fmt.Println(prompt)
	var move int
	_, err := fmt.Scanf("%d\n", &move)
	if err != nil {
		return 0, errors.ErrInvalidInput
	}
	return move, nil
}