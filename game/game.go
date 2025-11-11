package game

import (
	"errors"
	"fmt"
	"os"

	"github.com/WilliamAkaWill/tic-tac-toe/colors"
	custom_errors "github.com/WilliamAkaWill/tic-tac-toe/errors"
	"github.com/WilliamAkaWill/tic-tac-toe/player"
	"github.com/WilliamAkaWill/tic-tac-toe/print"
	"github.com/WilliamAkaWill/tic-tac-toe/shared"
	"github.com/WilliamAkaWill/tic-tac-toe/validate"
)

type (
	Game struct {
		languageService shared.LanguageService
	}
)

func NewGame(langService shared.LanguageService) *Game {
	return &Game{
		languageService: langService,
	}
}

func (g *Game) Exec() {
	player1, player2 := g.getPlayers()
	players := []player.Player{player1, player2}
	currentPlayer := player1
	printer := g.getPrinter()
	board := initBoard()

	summaryPrompt := g.languageService.GetString(shared.Summary)
	fmt.Println(summaryPrompt)
	playerXPrompt := g.languageService.GetString(shared.PlayerXSelection, player1.GetName())
	fmt.Println(playerXPrompt)
	playerOPrompt := g.languageService.GetString(shared.PlayerOSelection, player2.GetName())
	fmt.Println(playerOPrompt)
	printerSelectionPrompt := g.languageService.GetString(shared.PrinterSelection, printer.GetName())
	fmt.Println(printerSelectionPrompt)

	i := 0
	for {
		printer.PrintBoard(board)
		currentPlayer = players[i%2]
		move, err := currentPlayer.GetMove(board)
		if err != nil {
			colors.PrintError("%v\n", err)
			os.Exit(1)
		}

		row, col, err := validate.ValidateAndGetCoordinate(board, move)
		if err != nil {
			if errors.Is(err, custom_errors.ErrOutOfBounds) {
				outOfBoundsPrompt := g.languageService.GetString(shared.OutOfBounds)
				colors.PrintError("%s\n", outOfBoundsPrompt)
			}
			if errors.Is(err, custom_errors.ErrInvalidMove) {
				invalidMovePrompt := g.languageService.GetString(shared.InvalidMove)
				colors.PrintError("%s\n", invalidMovePrompt)
			}
			continue
		}

		board[row][col] = string(currentPlayer.GetPlayerMark())

		state := validate.CheckBoard(board)
		if state == shared.Player1WON || state == shared.Player2WON {
			playerWonPrompt := g.languageService.GetString(shared.PlayerWon, currentPlayer.GetName())
			colors.PrintSuccess("%s\n", playerWonPrompt)
			printer.PrintBoard(board)
			break
		}

		if state == shared.Tie {
			tieGamePrompt := g.languageService.GetString(shared.TieGame)
			colors.PrintWarning("%s\n", tieGamePrompt)
			printer.PrintBoard(board)
			break
		}

		i++
	}
}

func (g *Game) getPlayers() (player.Player, player.Player) {
	playerCountPrompt := g.languageService.GetString(shared.PlayerCount)
	fmt.Println(playerCountPrompt)
	var playerCount int
	_, err := fmt.Scanf("%d\n", &playerCount)
	if err != nil {
		fmt.Printf("%s: %v\n", custom_errors.ErrInvalidInput, err)
		os.Exit(1)
	}

	switch playerCount {
	case 1:
		human := player.NewHuman("Player 1", g.languageService, shared.PlayerX)
		computer := player.NewComputer(g.languageService)
		return human, computer
	case 2:
		human1 := player.NewHuman("Player 1", g.languageService, shared.PlayerX)
		human2 := player.NewHuman("Player 2", g.languageService, shared.PlayerO)
		return human1, human2

	default:
		fmt.Println(custom_errors.ErrInvalidInput)
		os.Exit(1)
	}
	return nil, nil
}

func (g *Game) getPrinter() print.Printer {
	printerPrompt := g.languageService.GetString(shared.ChoosePrinter)
	fmt.Println(printerPrompt)
	var printerChoice int
	_, err := fmt.Scanf("%d\n", &printerChoice)
	if err != nil {
		fmt.Printf("%s: %v\n", custom_errors.ErrInvalidInput, err)
		os.Exit(1)
	}

	typ, err := shared.GetPrintFromInt(printerChoice)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	return print.GetPrinter(typ)
}
