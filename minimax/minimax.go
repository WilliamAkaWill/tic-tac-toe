package minimax

import (
	"math"

	"github.com/WilliamAkaWill/tic-tac-toe/helper"
	"github.com/WilliamAkaWill/tic-tac-toe/shared"
	"github.com/WilliamAkaWill/tic-tac-toe/validate"
)

// Innerhalb dieser Datei wird der Minimax-Algorithmus implementiert.
//Für Referenz siehe: https://en.wikipedia.org/wiki/Minimax

// FindBestMove findet den besten Zug für den Computer (Spieler O) mittels Minimax-Algorithmus
//
// Der Minimax-Algorithmus bewertet rekursiv alle möglichen Spielzüge und wählt den
// optimalen Zug aus. Der Algorithmus geht davon aus, dass beide Spieler optimal spielen.
//
// Parameter:
//   - board: Das aktuelle Spielfeld (3x3 String-Array)
//
// Rückgabe:
//   - Move: Der beste Zug mit Zeile, Spalte und Bewertung
//
// Beispiel:
//
//	board := [][]string{
//	    {"X", "O", " "},
//	    {" ", "X", " "},
//	    {" ", " ", " "},
//	}
//	bestMove := FindBestMove(board)
//	// bestMove enthält die optimale Position für den nächsten Zug
func FindBestMove(board [][]string) (int, int) {
	bestMove := move{
		Row:   -1,
		Col:   -1,
		Score: math.MinInt32,
	}

	// Iteriere durch alle Zellen des Spielfelds
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			// Prüfe, ob die Zelle leer ist
			if board[row][col] == string(shared.Empty) {
				// Führe den Zug aus (für den Computer, Spieler O)
				board[row][col] = string(shared.PlayerO)

				// Berechne die Bewertung dieses Zugs rekursiv
				// false bedeutet, dass der nächste Zug vom Gegner (Minimierer) kommt
				moveScore := minimax(board, 0, false)

				// Mache den Zug rückgängig (Backtracking)
				board[row][col] = string(shared.Empty)

				// Aktualisiere den besten Zug, wenn dieser Zug besser ist
				if moveScore > bestMove.Score {
					bestMove.Row = row
					bestMove.Col = col
					bestMove.Score = moveScore
				}
			}
		}
	}

	return bestMove.Row, bestMove.Col
}

// minimax implementiert den Minimax-Algorithmus mit Rekursion
//
// Der Algorithmus funktioniert wie folgt:
// 1. Wenn das Spiel beendet ist (Sieg/Niederlage/Unentschieden), gib die Bewertung zurück
// 2. Wenn der Maximierer am Zug ist (Computer): Finde den Zug mit der höchsten Bewertung
// 3. Wenn der Minimierer am Zug ist (Gegner): Finde den Zug mit der niedrigsten Bewertung
//
// Parameter:
//   - board: Das aktuelle Spielfeld
//   - depth: Die aktuelle Tiefe im Spielbaum (verwendet für Bewertungsanpassung)
//   - isMaximizing: true = Computer (O) ist am Zug, false = Gegner (X) ist am Zug
//
// Rückgabe:
//   - int: Die Bewertung des Spielzustands aus Sicht des Computers
func minimax(board [][]string, depth int, isMaximizing bool) int {
	// Prüfe den aktuellen Spielzustand
	state := validate.CheckBoard(board)

	// Basisfall 1: Spieler 2 (O, Computer) hat gewonnen
	// Bewertung wird um depth reduziert, um schnellere Siege zu bevorzugen
	if state == shared.Player2WON {
		return winScore - depth
	}

	// Basisfall 2: Spieler 1 (X, Gegner) hat gewonnen
	// Bewertung wird um depth erhöht, um spätere Niederlagen zu bevorzugen
	if state == shared.Player1WON {
		return loseScore + depth
	}

	// Basisfall 3: Unentschieden
	if state == shared.Tie {
		return tieScore
	}

	// Rekursiver Fall: Das Spiel läuft noch
	if isMaximizing {
		// Maximierer (Computer) versucht, die Bewertung zu maximieren
		return maximizeScore(board, depth)
	}

	// Minimierer (Gegner) versucht, die Bewertung zu minimieren
	return minimizeScore(board, depth)

}

// maximizeScore findet die maximale Bewertung für den Computer (Spieler O)
//
// Diese Funktion iteriert durch alle möglichen Züge, führt jeden aus,
// bewertet ihn rekursiv und wählt die höchste Bewertung.
//
// Parameter:
//   - board: Das aktuelle Spielfeld
//   - depth: Die aktuelle Tiefe im Spielbaum
//
// Rückgabe:
//   - int: Die beste (höchste) Bewertung, die der Computer erreichen kann
func maximizeScore(board [][]string, depth int) int {
	bestScore := math.MinInt32

	// Probiere alle möglichen Züge aus
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			// Prüfe, ob die Zelle leer ist
			if board[row][col] == string(shared.Empty) {
				// Führe den Zug aus
				board[row][col] = string(shared.PlayerO)

				// Berechne die Bewertung rekursiv (nächster Zug: Minimierer)
				score := minimax(board, depth+1, false)

				// Mache den Zug rückgängig
				board[row][col] = string(shared.Empty)

				// Aktualisiere die beste Bewertung
				if score > bestScore {
					bestScore = score
				}
			}
		}
	}

	return bestScore
}

// minimizeScore findet die minimale Bewertung für den Gegner (Spieler X)
//
// Diese Funktion iteriert durch alle möglichen Züge, führt jeden aus,
// bewertet ihn rekursiv und wählt die niedrigste Bewertung.
// Dies simuliert einen optimal spielenden Gegner.
//
// Parameter:
//   - board: Das aktuelle Spielfeld
//   - depth: Die aktuelle Tiefe im Spielbaum
//
// Rückgabe:
//   - int: Die beste (niedrigste) Bewertung aus Sicht des Gegners
func minimizeScore(board [][]string, depth int) int {
	bestScore := math.MaxInt32

	// Probiere alle möglichen Züge aus
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			// Prüfe, ob die Zelle leer ist
			if board[row][col] == string(shared.Empty) {
				// Führe den Zug aus
				board[row][col] = string(shared.PlayerX)

				// Berechne die Bewertung rekursiv (nächster Zug: Maximierer)
				score := minimax(board, depth+1, true)

				// Mache den Zug rückgängig
				board[row][col] = string(shared.Empty)

				// Aktualisiere die beste Bewertung
				if score < bestScore {
					bestScore = score
				}
			}
		}
	}

	return bestScore
}

// GetAvailableMoves gibt alle verfügbaren (leeren) Positionen auf dem Brett zurück
//
// Hilfsfunktion, die eine Liste aller möglichen Züge erstellt.
// Nützlich für Debugging und alternative Implementierungen.
//
// Parameter:
//   - board: Das aktuelle Spielfeld
//
// Rückgabe:
//   - []move: Slice mit allen verfügbaren Positionen (Score ist 0)
func GetAvailableMoves(board [][]string) []move {
	availableMoves := helper.GetAvailableMoves(board)
	if len(availableMoves) == 0 {
		return nil
	}

	moves := make([]move, len(availableMoves))
	for i, pos := range availableMoves {
		row, col := shared.ConvertToCoordinates(pos)
		moves[i] = move{
			Row:   row,
			Col:   col,
			Score: 0,
		}
	}

	return moves
}
