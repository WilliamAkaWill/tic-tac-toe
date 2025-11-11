package print

// FancyPrinter uses ASCII art to display the Tic-Tac-Toe board in a stylish way.
// Special thanks to go-figure for the ASCII art rendering: https://github.com/common-nighthawk/go-figure

import (
	"fmt"

	"github.com/WilliamAkaWill/tic-tac-toe/shared"
	"github.com/common-nighthawk/go-figure"
)

type (
	FancyPrinter struct{}
)

func NewFancyPrinter() *FancyPrinter {
	return &FancyPrinter{}
}

func (p *FancyPrinter) GetName() string {
	return "Fancy Printer"
}

func (p *FancyPrinter) PrintBoard(board [][]string) {
	println("\n")
	firstRow := fmt.Sprintf(" %s | %s | %s ", getBoardChar(board, 0, 0), getBoardChar(board, 0, 1), getBoardChar(board, 0, 2))
	secondRow := fmt.Sprintf(" %s | %s | %s ", getBoardChar(board, 1, 0), getBoardChar(board, 1, 1), getBoardChar(board, 1, 2))
	thirdRow := fmt.Sprintf(" %s | %s | %s ", getBoardChar(board, 2, 0), getBoardChar(board, 2, 1), getBoardChar(board, 2, 2))
	firstFigure := figure.NewFigure(firstRow, "", false)
	secondFigure := figure.NewFigure(secondRow, "", false)
	thirdFigure := figure.NewFigure(thirdRow, "", false)

	for _, printRow := range firstFigure.Slicify() {
		fmt.Printf("%s\n", printRow)
	}
	fmt.Println("--------------------------------------------")

	for _, printRow := range secondFigure.Slicify() {
		fmt.Printf("%s\n", printRow)
	}
	fmt.Println("--------------------------------------------")

	for _, printRow := range thirdFigure.Slicify() {
		fmt.Printf("%s\n", printRow)
	}
}

func getBoardChar(board [][]string, i, j int) string {
	cell := board[i][j]
	if cell == string(shared.Empty) {
		cell = fmt.Sprintf("%d", shared.ConvertToPosition(i, j))
	}
	return cell
}
