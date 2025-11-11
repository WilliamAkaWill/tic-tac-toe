package shared

import "github.com/WilliamAkaWill/tic-tac-toe/errors"

type Print int

const (
	Normal Print = iota
	Fancy
)

func GetPrintFromInt(printInt int) (Print, error) {
	switch printInt {
	case 0:
		return Normal, nil
	case 1:
		return Fancy, nil
	default:
		return Normal, errors.ErrUnknown
	}
}