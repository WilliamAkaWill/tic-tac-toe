package print

import "github.com/WilliamAkaWill/tic-tac-toe/shared"

type (
	Printer interface {
		GetName() string
		PrintBoard(board [][]string)
	}
)

func GetPrinter(printerType shared.Print) Printer {
	switch printerType {
	case shared.Normal:
		return NewNormalPrinter()
	case shared.Fancy:
		return NewFancyPrinter()
	}
	return nil
}