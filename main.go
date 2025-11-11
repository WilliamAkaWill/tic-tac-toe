package main

import (
	"fmt"
	"os"

	"github.com/WilliamAkaWill/tic-tac-toe/errors"
	"github.com/WilliamAkaWill/tic-tac-toe/game"
	"github.com/WilliamAkaWill/tic-tac-toe/language"
	"github.com/WilliamAkaWill/tic-tac-toe/shared"
)

func main() {
	fmt.Println("Please select a language / Bitte wählen Sie ein Sprache (0 German, 1 English)")
	var langInt int
	_, err := fmt.Scanf("%d\n", &langInt)
	if err != nil {
		fmt.Printf("%s: %v\n", errors.ErrInvalidInput, err)
		os.Exit(1)
	}

	lang, err := shared.GetLanguageFromInt(langInt)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	langService := language.GetLanguageService(lang)
	game := game.NewGame(langService)
	game.Exec()
}
