package print

type (
	Printer interface {
		GetName() string
		PrintBoard(board [][]string)
	}
)