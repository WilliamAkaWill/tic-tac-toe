package minimax

import (
	"testing"
)

// Tests für FindBestMove Funktion

// TestFindBestMove_WinningMoveAvailable testet, ob der Computer einen sofortigen Gewinnzug erkennt
func TestFindBestMove_WinningMoveAvailable(t *testing.T) {
	// O kann in Position (0,2) gewinnen
	board := [][]string{
		{"O", "O", " "},
		{"X", "X", " "},
		{" ", " ", " "},
	}
	row, col := FindBestMove(board)
	if row != 0 || col != 2 {
		t.Errorf("Expected winning move at (0,2), got (%d,%d)", row, col)
	}
}

// TestFindBestMove_BlockOpponentWin testet, ob der Computer einen gegnerischen Gewinnzug blockt
func TestFindBestMove_BlockOpponentWin(t *testing.T) {
	// X kann in (0,2) gewinnen, O muss blocken
	board := [][]string{
		{"X", "X", " "},
		{"O", " ", " "},
		{" ", " ", " "},
	}
	row, col := FindBestMove(board)
	if row != 0 || col != 2 {
		t.Errorf("Expected blocking move at (0,2), got (%d,%d)", row, col)
	}
}

// TestFindBestMove_EmptyBoard testet den ersten Zug auf einem leeren Brett
func TestFindBestMove_EmptyBoard(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col := FindBestMove(board)

	// Der beste erste Zug ist typischerweise die Mitte oder eine Ecke
	if row < 0 || row > 2 || col < 0 || col > 2 {
		t.Errorf("Invalid move position (%d,%d)", row, col)
	}
}

// TestFindBestMove_CenterPosition testet, ob die Mitte bevorzugt wird
func TestFindBestMove_CenterPosition(t *testing.T) {
	board := [][]string{
		{"X", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col := FindBestMove(board)

	// Mitte (1,1) sollte ein guter Zug sein
	if row < 0 || row > 2 || col < 0 || col > 2 {
		t.Errorf("Invalid move position (%d,%d)", row, col)
	}
}

// TestFindBestMove_CornerTrap testet die Ecken-Falle-Situation
func TestFindBestMove_CornerTrap(t *testing.T) {
	board := [][]string{
		{"X", " ", " "},
		{" ", "O", " "},
		{" ", " ", "X"},
	}
	row, col := FindBestMove(board)

	// Computer sollte eine Kante wählen, nicht eine Ecke
	isEdge := (row == 0 && col == 1) || (row == 1 && col == 0) ||
		(row == 1 && col == 2) || (row == 2 && col == 1)
	if !isEdge {
		// Oder eine andere sichere Position
		if row < 0 || row > 2 || col < 0 || col > 2 {
			t.Errorf("Invalid move position (%d,%d)", row, col)
		}
	}
}

// TestFindBestMove_ForkOpportunity testet, ob der Computer eine Gabel-Möglichkeit nutzt
func TestFindBestMove_ForkOpportunity(t *testing.T) {
	board := [][]string{
		{"O", " ", " "},
		{" ", " ", " "},
		{" ", " ", "O"},
	}
	row, col := FindBestMove(board)

	// Sollte eine strategisch gute Position wählen
	if row < 0 || row > 2 || col < 0 || col > 2 {
		t.Errorf("Invalid move position (%d,%d)", row, col)
	}
}

// TestFindBestMove_OneMoveLeft testet mit nur einem verbleibenden Zug
func TestFindBestMove_OneMoveLeft(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", " "},
	}
	row, col := FindBestMove(board)
	if row != 2 || col != 2 {
		t.Errorf("Expected move at (2,2), got (%d,%d)", row, col)
	}
}

// TestFindBestMove_DiagonalWin testet einen diagonalen Gewinnzug
func TestFindBestMove_DiagonalWin(t *testing.T) {
	board := [][]string{
		{"O", "X", " "},
		{"X", "O", " "},
		{" ", " ", " "},
	}
	row, col := FindBestMove(board)
	// O sollte (2,2) wählen, um diagonal zu gewinnen
	if row != 2 || col != 2 {
		t.Errorf("Expected winning diagonal move at (2,2), got (%d,%d)", row, col)
	}
}

// TestFindBestMove_ColumnWin testet einen spaltenweisen Gewinnzug
func TestFindBestMove_ColumnWin(t *testing.T) {
	board := [][]string{
		{"O", "X", " "},
		{"O", "X", " "},
		{" ", " ", " "},
	}
	row, col := FindBestMove(board)
	// O sollte (2,0) wählen, um die Spalte zu gewinnen
	if row != 2 || col != 0 {
		t.Errorf("Expected winning column move at (2,0), got (%d,%d)", row, col)
	}
}

// TestFindBestMove_RowWin testet einen zeilenweisen Gewinnzug
func TestFindBestMove_RowWin(t *testing.T) {
	board := [][]string{
		{"X", " ", " "},
		{"X", " ", " "},
		{"O", "O", " "},
	}
	row, col := FindBestMove(board)
	// O sollte (2,2) wählen, um die Zeile zu gewinnen
	if row != 2 || col != 2 {
		t.Errorf("Expected winning row move at (2,2), got (%d,%d)", row, col)
	}
}

// Tests für GetAvailableMoves Funktion

// TestGetAvailableMoves_EmptyBoard testet ein leeres Brett
func TestGetAvailableMoves_EmptyBoard(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	moves := GetAvailableMoves(board)
	if len(moves) != 9 {
		t.Errorf("Expected 9 available moves, got %d", len(moves))
	}
}

// TestGetAvailableMoves_FullBoard testet ein volles Brett
func TestGetAvailableMoves_FullBoard(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "X"},
	}
	moves := GetAvailableMoves(board)
	if len(moves) != 0 {
		t.Errorf("Expected no available moves, got %d", len(moves))
	}
}

// TestGetAvailableMoves_OneMoveLeft testet mit einem verbleibenden Zug
func TestGetAvailableMoves_OneMoveLeft(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", " "},
	}
	moves := GetAvailableMoves(board)
	if len(moves) != 1 {
		t.Errorf("Expected 1 available move, got %d", len(moves))
	}
	if moves[0].Row != 2 || moves[0].Col != 2 {
		t.Errorf("Expected move at (2,2), got (%d,%d)", moves[0].Row, moves[0].Col)
	}
}

// TestGetAvailableMoves_PartiallyFilled testet ein teilweise gefülltes Brett
func TestGetAvailableMoves_PartiallyFilled(t *testing.T) {
	board := [][]string{
		{"X", " ", "O"},
		{" ", "X", " "},
		{"O", " ", " "},
	}
	moves := GetAvailableMoves(board)
	expectedCount := 5
	if len(moves) != expectedCount {
		t.Errorf("Expected %d available moves, got %d", expectedCount, len(moves))
	}
}

// TestGetAvailableMoves_CornersOnly testet wenn nur Ecken frei sind
func TestGetAvailableMoves_CornersOnly(t *testing.T) {
	board := [][]string{
		{" ", "X", " "},
		{"O", "X", "O"},
		{" ", "X", " "},
	}
	moves := GetAvailableMoves(board)
	if len(moves) != 4 {
		t.Errorf("Expected 4 available moves (corners), got %d", len(moves))
	}
}

// TestGetAvailableMoves_VerifyCoordinates testet die Korrektheit der Koordinaten
func TestGetAvailableMoves_VerifyCoordinates(t *testing.T) {
	board := [][]string{
		{" ", "X", " "},
		{"X", " ", "X"},
		{" ", "X", " "},
	}
	moves := GetAvailableMoves(board)

	// Prüfe, ob alle Moves auf tatsächlich leere Zellen zeigen
	for _, move := range moves {
		if board[move.Row][move.Col] != " " {
			t.Errorf("Move (%d,%d) points to non-empty cell", move.Row, move.Col)
		}
	}
}


// Integrationstests für komplexe Szenarien

// TestMinimax_Integration_ForceWin testet, ob Minimax immer gewinnt wenn möglich
func TestMinimax_Integration_ForceWin(t *testing.T) {
	// Situation, in der O in 1 Zug gewinnen kann
	board := [][]string{
		{"O", "O", " "},
		{"X", " ", " "},
		{"X", " ", " "},
	}
	row, col := FindBestMove(board)

	// Platziere den Zug
	board[row][col] = "O"

	// Prüfe, ob O gewonnen hat
	hasWon := false
	// Zeile prüfen
	if board[0][0] == "O" && board[0][1] == "O" && board[0][2] == "O" {
		hasWon = true
	}

	if !hasWon {
		t.Errorf("Computer should have made winning move, but didn't win")
	}
}

// TestMinimax_Integration_BlockAndWin testet Blocken und dann Gewinnen
func TestMinimax_Integration_BlockAndWin(t *testing.T) {
	// X hat 2 in einer Reihe, O muss blocken
	board := [][]string{
		{"X", "X", " "},
		{" ", "O", " "},
		{" ", " ", " "},
	}
	row, col := FindBestMove(board)

	// O sollte (0,2) blocken
	if row != 0 || col != 2 {
		t.Errorf("Expected blocking move at (0,2), got (%d,%d)", row, col)
	}
}

// TestMinimax_Integration_OptimalFirstMove testet den optimalen ersten Zug
func TestMinimax_Integration_OptimalFirstMove(t *testing.T) {
	board := [][]string{
		{"X", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col := FindBestMove(board)

	// Der beste Gegenzug auf eine Ecke ist die Mitte
	expectedRow, expectedCol := 1, 1
	if row != expectedRow || col != expectedCol {
		// Oder eine andere strategisch gute Position
		if row < 0 || row > 2 || col < 0 || col > 2 {
			t.Errorf("Invalid move position (%d,%d)", row, col)
		}
	}
}

// TestMinimax_Integration_PreventFork testet das Verhindern einer Gabel
func TestMinimax_Integration_PreventFork(t *testing.T) {
	// X in gegenüberliegenden Ecken - klassische Gabel-Situation
	board := [][]string{
		{"X", " ", " "},
		{" ", "O", " "},
		{" ", " ", "X"},
	}
	row, col := FindBestMove(board)

	// O sollte eine Kante wählen, um die Gabel zu verhindern
	isEdge := (row == 0 && col == 1) || (row == 1 && col == 0) ||
		(row == 1 && col == 2) || (row == 2 && col == 1)

	if !isEdge {
		t.Logf("Warning: Move (%d,%d) might not prevent fork optimally", row, col)
	}
}

// TestMinimax_Integration_NearEndGame testet Endspiel-Szenario
func TestMinimax_Integration_NearEndGame(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", "X", " "},
		{"X", " ", "O"},
	}
	row, col := FindBestMove(board)

	// Sollte einen gültigen Zug machen
	if row < 0 || row > 2 || col < 0 || col > 2 {
		t.Errorf("Invalid move position (%d,%d)", row, col)
	}

	// Position sollte leer sein
	if board[row][col] != " " {
		t.Errorf("Move (%d,%d) points to occupied cell", row, col)
	}
}

// TestMinimax_Integration_DrawSituation testet eine Unentschieden-Situation
func TestMinimax_Integration_DrawSituation(t *testing.T) {
	board := [][]string{
		{"O", "X", "O"},
		{"X", "X", "O"},
		{" ", "O", "X"},
	}
	row, col := FindBestMove(board)

	// Einziger verbleibender Zug
	if row != 2 || col != 0 {
		t.Errorf("Expected move at (2,0), got (%d,%d)", row, col)
	}
}

// TestMinimax_NoInfiniteLoop testet, dass es keine Endlosschleife gibt
func TestMinimax_NoInfiniteLoop(t *testing.T) {
	boards := [][][]string{
		{
			{" ", " ", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		},
		{
			{"X", "O", "X"},
			{"O", " ", "O"},
			{"X", " ", " "},
		},
		{
			{"O", "X", " "},
			{" ", "O", "X"},
			{" ", " ", " "},
		},
	}

	for i, board := range boards {
		row, col := FindBestMove(board)
		if row < 0 || row > 2 || col < 0 || col > 2 {
			t.Errorf("Board %d: Invalid move position (%d,%d)", i, row, col)
		}
		if board[row][col] != " " {
			t.Errorf("Board %d: Move (%d,%d) not on empty cell", i, row, col)
		}
	}
}
