package language

import (
	"fmt"

	"github.com/WilliamAkaWill/tic-tac-toe/shared"
)

type German struct{}

func NewGerman() *German {
	return &German{}
}

func (g *German) GetName() string {
	return "Deutsch"
}

func (g *German) GetString(key shared.Ressource, args ...string) string {
	switch key {
	case shared.InputTicTacToeNumber:
		return "Geben Sie eine Zahl von 1 bis 9 ein:"
	case shared.ChooseDifficulty:
		return "Wählen Sie die Schwierigkeit: Leicht (0), Mittel (1), Schwer (2)"
	case shared.PlayerCount:
		return "Bitte geben Sie die Anzahl der Spieler ein (1 oder 2):"
	case shared.ChoosePrinter:
		return "Wählen Sie den Drucker: Normal (0), Fancy (1)"
	case shared.Summary:
		return "Zusammenfassung:"
	case shared.LanguageSelection:
		return fmt.Sprintf("Sprache: %s", args[0])
	case shared.PlayerXSelection:
		return fmt.Sprintf("Spieler für X: %s", args[0])
	case shared.PlayerOSelection:
		return fmt.Sprintf("Spieler für O: %s", args[0])
	case shared.PrinterSelection:
		return fmt.Sprintf("Drucker: %s", args[0])
	case shared.OutOfBounds:
		return "Fehler: Die Eingabe liegt außerhalb des gültigen Bereichs. Bitte versuchen Sie es erneut."
	case shared.InvalidMove:
		return "Fehler: Ungültiger Zug. Bitte versuchen Sie es erneut."
	case shared.PlayerWon:
		return fmt.Sprintf("Spieler %s hat gewonnen, danke für das Spielen!", args[0])
	case shared.TieGame:
		return "Unentschieden, danke für das Spielen!"
	}

	return ""
}