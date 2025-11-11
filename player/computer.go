package player

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/WilliamAkaWill/tic-tac-toe/errors"
	"github.com/WilliamAkaWill/tic-tac-toe/helper"
	"github.com/WilliamAkaWill/tic-tac-toe/minimax"
	"github.com/WilliamAkaWill/tic-tac-toe/shared"
	"github.com/WilliamAkaWill/tic-tac-toe/validate"
)

type difficultFunc func(board [][]string) (int, int)

type (
	Computer struct {
		name            string
		languageService shared.LanguageService
		difficulty      difficultFunc
	}
)

func NewComputer(langService shared.LanguageService) *Computer {
	prompt := langService.GetString(shared.ChooseDifficulty)
	fmt.Println(prompt)
	var level int
	_, err := fmt.Scanf("%d\n", &level)
	if err != nil {
		fmt.Printf("%s: %v\n", errors.ErrInvalidInput, err)
		os.Exit(1)
	}

	difficulty, err := shared.GetDifficultyFromInt(level)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	difficultyFunc := getDifficultyFunction(difficulty)

	return &Computer{
		name:            "Computer",
		languageService: langService,
		difficulty:      difficultyFunc,
	}
}

func (c *Computer) GetName() string {
	return c.name
}

func (c *Computer) GetPlayerMark() shared.Player {
	return shared.PlayerO
}

func (c *Computer) GetMove(board [][]string) (int, error) {
	if c.difficulty == nil {
		return 0, errors.ErrUnknown
	}
	row, col := c.difficulty(board)
	move := shared.ConvertToPosition(row, col)
	return move, nil
}

func getDifficultyFunction(difficulty shared.Difficulty) difficultFunc {
	switch difficulty {
	case shared.Easy:
		return easyDifficulty
	case shared.Medium:
		return mediumDifficulty
	case shared.Hard:
		return minimax.FindBestMove
	}
	return nil
}

func easyDifficulty(board [][]string) (int, int) {
	availableMoves := helper.GetAvailableMoves(board)
	if len(availableMoves) == 0 {
		return 0, 0 // No moves available
	}
	move := availableMoves[rand.Intn(len(availableMoves))]
	row, col := shared.ConvertToCoordinates(move)
	return row, col
}

func mediumDifficulty(board [][]string) (int, int) {
	// 1. Prüfe, ob der Computer (O) gewinnen kann
	if row, col, found := findWinningMove(board, shared.Player2WON); found {
		return row, col
	}

	// 2. Prüfe, ob der Gegner (X) gewinnen kann und blockiere
	if row, col, found := findWinningMove(board, shared.Player1WON); found {
		return row, col
	}

	// 3. Wähle einen zufälligen verfügbaren Zug
	return easyDifficulty(board)
}

// findWinningMove sucht nach einem Zug, der den gegebenen Spieler gewinnen lässt
// Gibt (row, col, true) zurück wenn ein Gewinnzug gefunden wurde, sonst (0, 0, false)
func findWinningMove(board [][]string, player shared.State) (int, int, bool) {
	playerMark := string(shared.PlayerX)
	if player == shared.Player2WON {
		playerMark = string(shared.PlayerO)
	}

	// Probiere jeden verfügbaren Zug aus
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row][col] == string(shared.Empty) {
				// Simuliere den Zug
				board[row][col] = playerMark

				// Prüfe, ob dieser Zug gewinnt
				if validate.CheckBoard(board) == player {
					// Mache den Zug rückgängig
					board[row][col] = string(shared.Empty)
					return row, col, true
				}

				// Mache den Zug rückgängig
				board[row][col] = string(shared.Empty)
			}
		}
	}
	return 0, 0, false
}

