package print

import (
	"fmt"

	"github.com/WilliamAkaWill/tic-tac-toe/shared"
)

type (
	NormalPrinter struct{}
)

func NewNormalPrinter() *NormalPrinter {
	return &NormalPrinter{}
}

func (p *NormalPrinter) GetName() string {
	return "Normal Printer"
}

func (p *NormalPrinter) PrintBoard(board [][]string) {
	println("\n")
	for i, row := range board {
		print(" ")
		for j, cell := range row {
			if cell == string(shared.Empty) {
				cell = fmt.Sprintf("%d", shared.ConvertToPosition(i, j))
			}
			print(cell)
			if j < 2 {
				print(" | ")
			}
		}
		println()

		// Horizontale Trennlinie zwischen Zeilen
		if i < 2 {
			println("-----------")
		}
	}
	println()
}
