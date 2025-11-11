package helper

import (
	"reflect"
	"testing"
)

// TestGetAvailableMoves_EmptyBoard testet ein komplett leeres Spielfeld
// Erwartet: Alle Positionen 1-9 sind verfügbar
func TestGetAvailableMoves_EmptyBoard(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_FullBoard testet ein komplett besetztes Spielfeld
// Erwartet: Keine verfügbaren Züge (leeres Slice)
func TestGetAvailableMoves_FullBoard(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "X"},
	}
	expected := []int{}
	result := GetAvailableMoves(board)

	if result == nil {
		result = []int{}
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected empty slice, got %v", result)
	}
}

// TestGetAvailableMoves_OneMoveLeft testet ein fast volles Spielfeld
// Erwartet: Nur eine Position verfügbar
func TestGetAvailableMoves_OneMoveLeft(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", " "},
	}
	expected := []int{9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_FirstPositionOnly testet wenn nur Position 1 frei ist
func TestGetAvailableMoves_FirstPositionOnly(t *testing.T) {
	board := [][]string{
		{" ", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "X"},
	}
	expected := []int{1}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_MiddlePositionOnly testet wenn nur Position 5 (Mitte) frei ist
func TestGetAvailableMoves_MiddlePositionOnly(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", " ", "O"},
		{"X", "O", "X"},
	}
	expected := []int{5}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_CornersOnly testet wenn nur die vier Ecken frei sind
func TestGetAvailableMoves_CornersOnly(t *testing.T) {
	board := [][]string{
		{" ", "X", " "},
		{"O", "X", "O"},
		{" ", "X", " "},
	}
	expected := []int{1, 3, 7, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_EdgesOnly testet wenn nur die Kantenpositionen frei sind
func TestGetAvailableMoves_EdgesOnly(t *testing.T) {
	board := [][]string{
		{"X", " ", "X"},
		{" ", "O", " "},
		{"X", " ", "X"},
	}
	expected := []int{2, 4, 6, 8}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_AlternatingPattern testet ein Schachbrettmuster
func TestGetAvailableMoves_AlternatingPattern(t *testing.T) {
	board := [][]string{
		{" ", "X", " "},
		{"X", " ", "X"},
		{" ", "X", " "},
	}
	expected := []int{1, 3, 5, 7, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_FirstRowFree testet wenn nur die erste Zeile frei ist
func TestGetAvailableMoves_FirstRowFree(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{"X", "O", "X"},
		{"O", "X", "O"},
	}
	expected := []int{1, 2, 3}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_SecondRowFree testet wenn nur die zweite Zeile frei ist
func TestGetAvailableMoves_SecondRowFree(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{" ", " ", " "},
		{"O", "X", "O"},
	}
	expected := []int{4, 5, 6}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_ThirdRowFree testet wenn nur die dritte Zeile frei ist
func TestGetAvailableMoves_ThirdRowFree(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{" ", " ", " "},
	}
	expected := []int{7, 8, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_FirstColumnFree testet wenn nur die erste Spalte frei ist
func TestGetAvailableMoves_FirstColumnFree(t *testing.T) {
	board := [][]string{
		{" ", "O", "X"},
		{" ", "X", "O"},
		{" ", "O", "X"},
	}
	expected := []int{1, 4, 7}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_SecondColumnFree testet wenn nur die zweite Spalte frei ist
func TestGetAvailableMoves_SecondColumnFree(t *testing.T) {
	board := [][]string{
		{"X", " ", "O"},
		{"O", " ", "X"},
		{"X", " ", "O"},
	}
	expected := []int{2, 5, 8}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_ThirdColumnFree testet wenn nur die dritte Spalte frei ist
func TestGetAvailableMoves_ThirdColumnFree(t *testing.T) {
	board := [][]string{
		{"X", "O", " "},
		{"O", "X", " "},
		{"X", "O", " "},
	}
	expected := []int{3, 6, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_DiagonalFree testet wenn nur die Hauptdiagonale frei ist
func TestGetAvailableMoves_DiagonalFree(t *testing.T) {
	board := [][]string{
		{" ", "O", "X"},
		{"X", " ", "O"},
		{"O", "X", " "},
	}
	expected := []int{1, 5, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_AntiDiagonalFree testet wenn nur die Nebendiagonale frei ist
func TestGetAvailableMoves_AntiDiagonalFree(t *testing.T) {
	board := [][]string{
		{"X", "O", " "},
		{"O", " ", "X"},
		{" ", "X", "O"},
	}
	expected := []int{3, 5, 7}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_EarlyGame testet eine typische frühe Spielsituation
func TestGetAvailableMoves_EarlyGame(t *testing.T) {
	board := [][]string{
		{"X", " ", " "},
		{" ", "O", " "},
		{" ", " ", " "},
	}
	expected := []int{2, 3, 4, 6, 7, 8, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_MidGame testet eine typische Mittelspielsituation
func TestGetAvailableMoves_MidGame(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{" ", "X", " "},
		{"O", " ", " "},
	}
	expected := []int{4, 6, 8, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_LateGame testet eine späte Spielsituation
func TestGetAvailableMoves_LateGame(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{" ", " ", "X"},
	}
	expected := []int{7, 8}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_OnlyOMarks testet ein Brett nur mit O-Markierungen
func TestGetAvailableMoves_OnlyOMarks(t *testing.T) {
	board := [][]string{
		{" ", "O", " "},
		{"O", " ", "O"},
		{" ", "O", " "},
	}
	expected := []int{1, 3, 5, 7, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_OnlyXMarks testet ein Brett nur mit X-Markierungen
func TestGetAvailableMoves_OnlyXMarks(t *testing.T) {
	board := [][]string{
		{"X", " ", "X"},
		{" ", "X", " "},
		{"X", " ", "X"},
	}
	expected := []int{2, 4, 6, 8}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_SingleMoveInEachRow testet ein Move pro Zeile
func TestGetAvailableMoves_SingleMoveInEachRow(t *testing.T) {
	board := [][]string{
		{" ", "X", "O"},
		{"X", " ", "O"},
		{"X", "O", " "},
	}
	expected := []int{1, 5, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_RandomPattern1 testet ein zufälliges Muster
func TestGetAvailableMoves_RandomPattern1(t *testing.T) {
	board := [][]string{
		{"X", " ", "O"},
		{" ", "X", " "},
		{" ", " ", "O"},
	}
	expected := []int{2, 4, 6, 7, 8}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_RandomPattern2 testet ein weiteres zufälliges Muster
func TestGetAvailableMoves_RandomPattern2(t *testing.T) {
	board := [][]string{
		{" ", "X", " "},
		{"O", " ", "X"},
		{" ", "O", " "},
	}
	expected := []int{1, 3, 5, 7, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_ConsecutiveFreePositions testet aufeinanderfolgende freie Positionen
func TestGetAvailableMoves_ConsecutiveFreePositions(t *testing.T) {
	board := [][]string{
		{"X", "X", "X"},
		{" ", " ", " "},
		{"O", "O", "O"},
	}
	expected := []int{4, 5, 6}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_TwoMovesLeft testet mit genau zwei freien Positionen
func TestGetAvailableMoves_TwoMovesLeft(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", " ", "X"},
		{"X", " ", "O"},
	}
	expected := []int{5, 8}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_ThreeMovesLeft testet mit genau drei freien Positionen
func TestGetAvailableMoves_ThreeMovesLeft(t *testing.T) {
	board := [][]string{
		{"X", "O", " "},
		{"O", "X", " "},
		{" ", "O", "X"},
	}
	expected := []int{3, 6, 7}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_FourMovesLeft testet mit genau vier freien Positionen
func TestGetAvailableMoves_FourMovesLeft(t *testing.T) {
	board := [][]string{
		{" ", "X", "O"},
		{"X", " ", "O"},
		{" ", " ", "X"},
	}
	expected := []int{1, 5, 7, 8}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestGetAvailableMoves_SymmetricPattern testet ein symmetrisches Muster
func TestGetAvailableMoves_SymmetricPattern(t *testing.T) {
	board := [][]string{
		{" ", "X", " "},
		{"X", "O", "X"},
		{" ", "X", " "},
	}
	expected := []int{1, 3, 7, 9}
	result := GetAvailableMoves(board)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
