package shared

import (
	"github.com/WilliamAkaWill/tic-tac-toe/errors"
)


type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)


func GetDifficultyFromInt(value int) (Difficulty, error) {
	switch value {
	case 0:
		return Easy, nil
	case 1:
		return Medium, nil
	case 2:
		return Hard, nil
	default:
		return Easy, errors.ErrUnknown
	}
}