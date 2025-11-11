package shared

import "github.com/WilliamAkaWill/tic-tac-toe/errors"

type Language int

const (
	German Language = iota
	English
)

type Ressource int

const (
	InputTicTacToeNumber Ressource = iota
	ChooseDifficulty
	PlayerCount
	Summary 
	LanguageSelection
	PlayerXSelection
	PlayerOSelection
	PrinterSelection
	OutOfBounds
	InvalidMove
	PlayerWon
	TieGame
	ChoosePrinter
)

type (
	LanguageService interface {
		GetName() string
		GetString(key Ressource, args ...string) string
	}
)

func GetLanguageFromInt(langInt int) (Language, error) {
	switch langInt {
	case 0:
		return German, nil
	case 1:
		return English, nil
	default:
		return German, errors.ErrUnknown
	}
}