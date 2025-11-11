package language

import (
	"fmt"

	"github.com/WilliamAkaWill/tic-tac-toe/shared"
)

type English struct{}

func NewEnglish() *English {
	return &English{}
}

func (g *English) GetName() string {
	return "English"
}

func (g *English) GetString(key shared.Ressource, args ...string) string {
	switch key {
	case shared.InputTicTacToeNumber:
		return fmt.Sprintf("%s, Enter a number from 1 to 9:", args[0])
	case shared.ChooseDifficulty:
		return "Choose difficulty: Easy (0), Medium (1), Hard (2)"
	case shared.PlayerCount:
		return "Please enter the number of players (1 or 2):"
	case shared.ChoosePrinter:
		return "Choose printer: Normal (0), Fancy (1)"
	case shared.Summary:
		return "Summary:"
	case shared.LanguageSelection:
		return fmt.Sprintf("Language: %s", args[0])
	case shared.PlayerXSelection:
		return fmt.Sprintf("Player for X: %s", args[0])
	case shared.PlayerOSelection:
		return fmt.Sprintf("Player for O: %s", args[0])
	case shared.PrinterSelection:
		return fmt.Sprintf("Printer: %s", args[0])
	case shared.OutOfBounds:
		return "Error: The input is out of bounds. Please try again."
	case shared.InvalidMove:
		return "Error: Invalid move. Please try again."
	case shared.PlayerWon:
		return fmt.Sprintf("Player %s has won, thank you for playing!", args[0])
	case shared.TieGame:
		return "It's a tie, thank you for playing!"
	}
	return ""
}