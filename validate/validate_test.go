package validate

import (
	"testing"

	"github.com/WilliamAkaWill/tic-tac-toe/errors"
	"github.com/WilliamAkaWill/tic-tac-toe/shared"
)

// Tests für CheckBoard Funktion

func TestCheckBoard_Player1WinsRow1(t *testing.T) {
	board := [][]string{
		{"X", "X", "X"},
		{" ", "O", " "},
		{"O", " ", " "},
	}
	result := CheckBoard(board)
	if result != shared.Player1WON {
		t.Errorf("Expected Player1WON, got %v", result)
	}
}

func TestCheckBoard_Player1WinsRow2(t *testing.T) {
	board := [][]string{
		{"O", " ", " "},
		{"X", "X", "X"},
		{" ", "O", " "},
	}
	result := CheckBoard(board)
	if result != shared.Player1WON {
		t.Errorf("Expected Player1WON, got %v", result)
	}
}

func TestCheckBoard_Player1WinsRow3(t *testing.T) {
	board := [][]string{
		{" ", "O", " "},
		{"O", " ", " "},
		{"X", "X", "X"},
	}
	result := CheckBoard(board)
	if result != shared.Player1WON {
		t.Errorf("Expected Player1WON, got %v", result)
	}
}

func TestCheckBoard_Player1WinsColumn1(t *testing.T) {
	board := [][]string{
		{"X", "O", " "},
		{"X", " ", "O"},
		{"X", " ", " "},
	}
	result := CheckBoard(board)
	if result != shared.Player1WON {
		t.Errorf("Expected Player1WON, got %v", result)
	}
}

func TestCheckBoard_Player1WinsColumn2(t *testing.T) {
	board := [][]string{
		{"O", "X", " "},
		{" ", "X", "O"},
		{" ", "X", " "},
	}
	result := CheckBoard(board)
	if result != shared.Player1WON {
		t.Errorf("Expected Player1WON, got %v", result)
	}
}

func TestCheckBoard_Player1WinsColumn3(t *testing.T) {
	board := [][]string{
		{"O", " ", "X"},
		{" ", "O", "X"},
		{" ", " ", "X"},
	}
	result := CheckBoard(board)
	if result != shared.Player1WON {
		t.Errorf("Expected Player1WON, got %v", result)
	}
}

func TestCheckBoard_Player1WinsDiagonal1(t *testing.T) {
	board := [][]string{
		{"X", "O", " "},
		{" ", "X", "O"},
		{" ", " ", "X"},
	}
	result := CheckBoard(board)
	if result != shared.Player1WON {
		t.Errorf("Expected Player1WON, got %v", result)
	}
}

func TestCheckBoard_Player1WinsDiagonal2(t *testing.T) {
	board := [][]string{
		{"O", " ", "X"},
		{" ", "X", "O"},
		{"X", " ", " "},
	}
	result := CheckBoard(board)
	if result != shared.Player1WON {
		t.Errorf("Expected Player1WON, got %v", result)
	}
}

func TestCheckBoard_Player2WinsRow1(t *testing.T) {
	board := [][]string{
		{"O", "O", "O"},
		{"X", "X", " "},
		{" ", "X", " "},
	}
	result := CheckBoard(board)
	if result != shared.Player2WON {
		t.Errorf("Expected Player2WON, got %v", result)
	}
}

func TestCheckBoard_Player2WinsRow2(t *testing.T) {
	board := [][]string{
		{"X", "X", " "},
		{"O", "O", "O"},
		{" ", "X", " "},
	}
	result := CheckBoard(board)
	if result != shared.Player2WON {
		t.Errorf("Expected Player2WON, got %v", result)
	}
}

func TestCheckBoard_Player2WinsRow3(t *testing.T) {
	board := [][]string{
		{"X", "X", " "},
		{" ", "X", " "},
		{"O", "O", "O"},
	}
	result := CheckBoard(board)
	if result != shared.Player2WON {
		t.Errorf("Expected Player2WON, got %v", result)
	}
}

func TestCheckBoard_Player2WinsColumn1(t *testing.T) {
	board := [][]string{
		{"O", "X", " "},
		{"O", "X", " "},
		{"O", " ", "X"},
	}
	result := CheckBoard(board)
	if result != shared.Player2WON {
		t.Errorf("Expected Player2WON, got %v", result)
	}
}

func TestCheckBoard_Player2WinsColumn2(t *testing.T) {
	board := [][]string{
		{"X", "O", " "},
		{" ", "O", "X"},
		{" ", "O", "X"},
	}
	result := CheckBoard(board)
	if result != shared.Player2WON {
		t.Errorf("Expected Player2WON, got %v", result)
	}
}

func TestCheckBoard_Player2WinsColumn3(t *testing.T) {
	board := [][]string{
		{"X", " ", "O"},
		{" ", "X", "O"},
		{"X", " ", "O"},
	}
	result := CheckBoard(board)
	if result != shared.Player2WON {
		t.Errorf("Expected Player2WON, got %v", result)
	}
}

func TestCheckBoard_Player2WinsDiagonal1(t *testing.T) {
	board := [][]string{
		{"O", "X", " "},
		{"X", "O", " "},
		{" ", "X", "O"},
	}
	result := CheckBoard(board)
	if result != shared.Player2WON {
		t.Errorf("Expected Player2WON, got %v", result)
	}
}

func TestCheckBoard_Player2WinsDiagonal2(t *testing.T) {
	board := [][]string{
		{"X", " ", "O"},
		{" ", "O", "X"},
		{"O", "X", " "},
	}
	result := CheckBoard(board)
	if result != shared.Player2WON {
		t.Errorf("Expected Player2WON, got %v", result)
	}
}

func TestCheckBoard_Tie(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"X", "O", "O"},
		{"O", "X", "X"},
	}
	result := CheckBoard(board)
	if result != shared.Tie {
		t.Errorf("Expected Tie, got %v", result)
	}
}

func TestCheckBoard_TieAlternativePattern(t *testing.T) {
	board := [][]string{
		{"O", "X", "O"},
		{"X", "X", "O"},
		{"X", "O", "X"},
	}
	result := CheckBoard(board)
	if result != shared.Tie {
		t.Errorf("Expected Tie, got %v", result)
	}
}

func TestCheckBoard_InProgressEmptyBoard(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	result := CheckBoard(board)
	if result != shared.InProgress {
		t.Errorf("Expected InProgress, got %v", result)
	}
}

func TestCheckBoard_InProgressOneMark(t *testing.T) {
	board := [][]string{
		{"X", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	result := CheckBoard(board)
	if result != shared.InProgress {
		t.Errorf("Expected InProgress, got %v", result)
	}
}

func TestCheckBoard_InProgressSeveralMarks(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", "X", " "},
		{" ", " ", " "},
	}
	result := CheckBoard(board)
	if result != shared.InProgress {
		t.Errorf("Expected InProgress, got %v", result)
	}
}

func TestCheckBoard_InProgressOneEmptyCell(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"X", "O", "O"},
		{"O", "X", " "},
	}
	result := CheckBoard(board)
	if result != shared.InProgress {
		t.Errorf("Expected InProgress, got %v", result)
	}
}

// Tests für ValidateAndGetCoordinate Funktion

func TestValidateAndGetCoordinate_ValidInput1(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if row != 0 || col != 0 {
		t.Errorf("Expected row=0, col=0, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_ValidInput2(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if row != 0 || col != 1 {
		t.Errorf("Expected row=0, col=1, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_ValidInput3(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if row != 0 || col != 2 {
		t.Errorf("Expected row=0, col=2, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_ValidInput4(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if row != 1 || col != 0 {
		t.Errorf("Expected row=1, col=0, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_ValidInput5(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if row != 1 || col != 1 {
		t.Errorf("Expected row=1, col=1, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_ValidInput6(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 6)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if row != 1 || col != 2 {
		t.Errorf("Expected row=1, col=2, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_ValidInput7(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 7)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if row != 2 || col != 0 {
		t.Errorf("Expected row=2, col=0, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_ValidInput8(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 8)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if row != 2 || col != 1 {
		t.Errorf("Expected row=2, col=1, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_ValidInput9(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 9)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if row != 2 || col != 2 {
		t.Errorf("Expected row=2, col=2, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_OutOfBoundsZero(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 0)
	if err != errors.ErrOutOfBounds {
		t.Errorf("Expected ErrOutOfBounds, got %v", err)
	}
	if row != 0 || col != 0 {
		t.Errorf("Expected row=0, col=0, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_OutOfBoundsNegative(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, -1)
	if err != errors.ErrOutOfBounds {
		t.Errorf("Expected ErrOutOfBounds, got %v", err)
	}
	if row != 0 || col != 0 {
		t.Errorf("Expected row=0, col=0, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_OutOfBounds10(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 10)
	if err != errors.ErrOutOfBounds {
		t.Errorf("Expected ErrOutOfBounds, got %v", err)
	}
	if row != 0 || col != 0 {
		t.Errorf("Expected row=0, col=0, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_OutOfBounds100(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 100)
	if err != errors.ErrOutOfBounds {
		t.Errorf("Expected ErrOutOfBounds, got %v", err)
	}
	if row != 0 || col != 0 {
		t.Errorf("Expected row=0, col=0, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_CellOccupiedByX(t *testing.T) {
	board := [][]string{
		{"X", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 1)
	if err != errors.ErrInvalidMove {
		t.Errorf("Expected ErrInvalidMove, got %v", err)
	}
	if row != 0 || col != 0 {
		t.Errorf("Expected row=0, col=0, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_CellOccupiedByO(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", "O", " "},
		{" ", " ", " "},
	}
	row, col, err := ValidateAndGetCoordinate(board, 5)
	if err != errors.ErrInvalidMove {
		t.Errorf("Expected ErrInvalidMove, got %v", err)
	}
	if row != 0 || col != 0 {
		t.Errorf("Expected row=0, col=0, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_LastCellOccupied(t *testing.T) {
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", "X"},
	}
	row, col, err := ValidateAndGetCoordinate(board, 9)
	if err != errors.ErrInvalidMove {
		t.Errorf("Expected ErrInvalidMove, got %v", err)
	}
	if row != 0 || col != 0 {
		t.Errorf("Expected row=0, col=0, got row=%d, col=%d", row, col)
	}
}

func TestValidateAndGetCoordinate_AllCellsOccupied(t *testing.T) {
	board := [][]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"O", "X", "O"},
	}
	for i := 1; i <= 9; i++ {
		_, _, err := ValidateAndGetCoordinate(board, i)
		if err != errors.ErrInvalidMove {
			t.Errorf("Expected ErrInvalidMove for input %d, got %v", i, err)
		}
	}
}

func TestValidateAndGetCoordinate_MixedBoardValidMoves(t *testing.T) {
	board := [][]string{
		{"X", " ", "O"},
		{" ", "X", " "},
		{"O", " ", " "},
	}
	testCases := []struct {
		input       int
		expectedRow int
		expectedCol int
	}{
		{2, 0, 1},
		{4, 1, 0},
		{6, 1, 2},
		{8, 2, 1},
		{9, 2, 2},
	}

	for _, tc := range testCases {
		row, col, err := ValidateAndGetCoordinate(board, tc.input)
		if err != nil {
			t.Errorf("Expected no error for input %d, got %v", tc.input, err)
		}
		if row != tc.expectedRow || col != tc.expectedCol {
			t.Errorf("For input %d: expected row=%d, col=%d, got row=%d, col=%d",
				tc.input, tc.expectedRow, tc.expectedCol, row, col)
		}
	}
}
