package shared

import "testing"

// Tests für ConvertToCoordinates Funktion

// TestConvertToCoordinates_Position1 testet Position 1 (oben links)
func TestConvertToCoordinates_Position1(t *testing.T) {
	row, col := ConvertToCoordinates(1)
	expectedRow, expectedCol := 0, 0
	if row != expectedRow || col != expectedCol {
		t.Errorf("Position 1: expected (%d,%d), got (%d,%d)", expectedRow, expectedCol, row, col)
	}
}

// TestConvertToCoordinates_Position2 testet Position 2 (oben mitte)
func TestConvertToCoordinates_Position2(t *testing.T) {
	row, col := ConvertToCoordinates(2)
	expectedRow, expectedCol := 0, 1
	if row != expectedRow || col != expectedCol {
		t.Errorf("Position 2: expected (%d,%d), got (%d,%d)", expectedRow, expectedCol, row, col)
	}
}

// TestConvertToCoordinates_Position3 testet Position 3 (oben rechts)
func TestConvertToCoordinates_Position3(t *testing.T) {
	row, col := ConvertToCoordinates(3)
	expectedRow, expectedCol := 0, 2
	if row != expectedRow || col != expectedCol {
		t.Errorf("Position 3: expected (%d,%d), got (%d,%d)", expectedRow, expectedCol, row, col)
	}
}

// TestConvertToCoordinates_Position4 testet Position 4 (mitte links)
func TestConvertToCoordinates_Position4(t *testing.T) {
	row, col := ConvertToCoordinates(4)
	expectedRow, expectedCol := 1, 0
	if row != expectedRow || col != expectedCol {
		t.Errorf("Position 4: expected (%d,%d), got (%d,%d)", expectedRow, expectedCol, row, col)
	}
}

// TestConvertToCoordinates_Position5 testet Position 5 (mitte mitte/Zentrum)
func TestConvertToCoordinates_Position5(t *testing.T) {
	row, col := ConvertToCoordinates(5)
	expectedRow, expectedCol := 1, 1
	if row != expectedRow || col != expectedCol {
		t.Errorf("Position 5: expected (%d,%d), got (%d,%d)", expectedRow, expectedCol, row, col)
	}
}

// TestConvertToCoordinates_Position6 testet Position 6 (mitte rechts)
func TestConvertToCoordinates_Position6(t *testing.T) {
	row, col := ConvertToCoordinates(6)
	expectedRow, expectedCol := 1, 2
	if row != expectedRow || col != expectedCol {
		t.Errorf("Position 6: expected (%d,%d), got (%d,%d)", expectedRow, expectedCol, row, col)
	}
}

// TestConvertToCoordinates_Position7 testet Position 7 (unten links)
func TestConvertToCoordinates_Position7(t *testing.T) {
	row, col := ConvertToCoordinates(7)
	expectedRow, expectedCol := 2, 0
	if row != expectedRow || col != expectedCol {
		t.Errorf("Position 7: expected (%d,%d), got (%d,%d)", expectedRow, expectedCol, row, col)
	}
}

// TestConvertToCoordinates_Position8 testet Position 8 (unten mitte)
func TestConvertToCoordinates_Position8(t *testing.T) {
	row, col := ConvertToCoordinates(8)
	expectedRow, expectedCol := 2, 1
	if row != expectedRow || col != expectedCol {
		t.Errorf("Position 8: expected (%d,%d), got (%d,%d)", expectedRow, expectedCol, row, col)
	}
}

// TestConvertToCoordinates_Position9 testet Position 9 (unten rechts)
func TestConvertToCoordinates_Position9(t *testing.T) {
	row, col := ConvertToCoordinates(9)
	expectedRow, expectedCol := 2, 2
	if row != expectedRow || col != expectedCol {
		t.Errorf("Position 9: expected (%d,%d), got (%d,%d)", expectedRow, expectedCol, row, col)
	}
}

// TestConvertToCoordinates_AllPositions testet alle Positionen 1-9 systematisch
func TestConvertToCoordinates_AllPositions(t *testing.T) {
	testCases := []struct {
		position    int
		expectedRow int
		expectedCol int
	}{
		{1, 0, 0}, {2, 0, 1}, {3, 0, 2},
		{4, 1, 0}, {5, 1, 1}, {6, 1, 2},
		{7, 2, 0}, {8, 2, 1}, {9, 2, 2},
	}

	for _, tc := range testCases {
		row, col := ConvertToCoordinates(tc.position)
		if row != tc.expectedRow || col != tc.expectedCol {
			t.Errorf("Position %d: expected (%d,%d), got (%d,%d)",
				tc.position, tc.expectedRow, tc.expectedCol, row, col)
		}
	}
}

// TestConvertToCoordinates_FirstRow testet alle Positionen der ersten Zeile
func TestConvertToCoordinates_FirstRow(t *testing.T) {
	for pos := 1; pos <= 3; pos++ {
		row, _ := ConvertToCoordinates(pos)
		if row != 0 {
			t.Errorf("Position %d should be in row 0, got row %d", pos, row)
		}
	}
}

// TestConvertToCoordinates_SecondRow testet alle Positionen der zweiten Zeile
func TestConvertToCoordinates_SecondRow(t *testing.T) {
	for pos := 4; pos <= 6; pos++ {
		row, _ := ConvertToCoordinates(pos)
		if row != 1 {
			t.Errorf("Position %d should be in row 1, got row %d", pos, row)
		}
	}
}

// TestConvertToCoordinates_ThirdRow testet alle Positionen der dritten Zeile
func TestConvertToCoordinates_ThirdRow(t *testing.T) {
	for pos := 7; pos <= 9; pos++ {
		row, _ := ConvertToCoordinates(pos)
		if row != 2 {
			t.Errorf("Position %d should be in row 2, got row %d", pos, row)
		}
	}
}

// TestConvertToCoordinates_FirstColumn testet alle Positionen der ersten Spalte
func TestConvertToCoordinates_FirstColumn(t *testing.T) {
	positions := []int{1, 4, 7}
	for _, pos := range positions {
		_, col := ConvertToCoordinates(pos)
		if col != 0 {
			t.Errorf("Position %d should be in col 0, got col %d", pos, col)
		}
	}
}

// TestConvertToCoordinates_SecondColumn testet alle Positionen der zweiten Spalte
func TestConvertToCoordinates_SecondColumn(t *testing.T) {
	positions := []int{2, 5, 8}
	for _, pos := range positions {
		_, col := ConvertToCoordinates(pos)
		if col != 1 {
			t.Errorf("Position %d should be in col 1, got col %d", pos, col)
		}
	}
}

// TestConvertToCoordinates_ThirdColumn testet alle Positionen der dritten Spalte
func TestConvertToCoordinates_ThirdColumn(t *testing.T) {
	positions := []int{3, 6, 9}
	for _, pos := range positions {
		_, col := ConvertToCoordinates(pos)
		if col != 2 {
			t.Errorf("Position %d should be in col 2, got col %d", pos, col)
		}
	}
}

// TestConvertToCoordinates_Corners testet alle Eckpositionen
func TestConvertToCoordinates_Corners(t *testing.T) {
	corners := map[int][2]int{
		1: {0, 0}, // oben links
		3: {0, 2}, // oben rechts
		7: {2, 0}, // unten links
		9: {2, 2}, // unten rechts
	}

	for pos, expected := range corners {
		row, col := ConvertToCoordinates(pos)
		if row != expected[0] || col != expected[1] {
			t.Errorf("Corner position %d: expected (%d,%d), got (%d,%d)",
				pos, expected[0], expected[1], row, col)
		}
	}
}

// TestConvertToCoordinates_Edges testet alle Kantenpositionen
func TestConvertToCoordinates_Edges(t *testing.T) {
	edges := map[int][2]int{
		2: {0, 1}, // oben mitte
		4: {1, 0}, // mitte links
		6: {1, 2}, // mitte rechts
		8: {2, 1}, // unten mitte
	}

	for pos, expected := range edges {
		row, col := ConvertToCoordinates(pos)
		if row != expected[0] || col != expected[1] {
			t.Errorf("Edge position %d: expected (%d,%d), got (%d,%d)",
				pos, expected[0], expected[1], row, col)
		}
	}
}

// Tests für ConvertToPosition Funktion

// TestConvertToPosition_TopLeft testet die obere linke Ecke (0,0)
func TestConvertToPosition_TopLeft(t *testing.T) {
	pos := ConvertToPosition(0, 0)
	expected := 1
	if pos != expected {
		t.Errorf("(0,0): expected position %d, got %d", expected, pos)
	}
}

// TestConvertToPosition_TopMiddle testet die obere Mitte (0,1)
func TestConvertToPosition_TopMiddle(t *testing.T) {
	pos := ConvertToPosition(0, 1)
	expected := 2
	if pos != expected {
		t.Errorf("(0,1): expected position %d, got %d", expected, pos)
	}
}

// TestConvertToPosition_TopRight testet die obere rechte Ecke (0,2)
func TestConvertToPosition_TopRight(t *testing.T) {
	pos := ConvertToPosition(0, 2)
	expected := 3
	if pos != expected {
		t.Errorf("(0,2): expected position %d, got %d", expected, pos)
	}
}

// TestConvertToPosition_MiddleLeft testet die mittlere linke Position (1,0)
func TestConvertToPosition_MiddleLeft(t *testing.T) {
	pos := ConvertToPosition(1, 0)
	expected := 4
	if pos != expected {
		t.Errorf("(1,0): expected position %d, got %d", expected, pos)
	}
}

// TestConvertToPosition_Center testet die Zentrumsposition (1,1)
func TestConvertToPosition_Center(t *testing.T) {
	pos := ConvertToPosition(1, 1)
	expected := 5
	if pos != expected {
		t.Errorf("(1,1): expected position %d, got %d", expected, pos)
	}
}

// TestConvertToPosition_MiddleRight testet die mittlere rechte Position (1,2)
func TestConvertToPosition_MiddleRight(t *testing.T) {
	pos := ConvertToPosition(1, 2)
	expected := 6
	if pos != expected {
		t.Errorf("(1,2): expected position %d, got %d", expected, pos)
	}
}

// TestConvertToPosition_BottomLeft testet die untere linke Ecke (2,0)
func TestConvertToPosition_BottomLeft(t *testing.T) {
	pos := ConvertToPosition(2, 0)
	expected := 7
	if pos != expected {
		t.Errorf("(2,0): expected position %d, got %d", expected, pos)
	}
}

// TestConvertToPosition_BottomMiddle testet die untere Mitte (2,1)
func TestConvertToPosition_BottomMiddle(t *testing.T) {
	pos := ConvertToPosition(2, 1)
	expected := 8
	if pos != expected {
		t.Errorf("(2,1): expected position %d, got %d", expected, pos)
	}
}

// TestConvertToPosition_BottomRight testet die untere rechte Ecke (2,2)
func TestConvertToPosition_BottomRight(t *testing.T) {
	pos := ConvertToPosition(2, 2)
	expected := 9
	if pos != expected {
		t.Errorf("(2,2): expected position %d, got %d", expected, pos)
	}
}

// TestConvertToPosition_AllCoordinates testet alle Koordinaten systematisch
func TestConvertToPosition_AllCoordinates(t *testing.T) {
	testCases := []struct {
		row              int
		col              int
		expectedPosition int
	}{
		{0, 0, 1}, {0, 1, 2}, {0, 2, 3},
		{1, 0, 4}, {1, 1, 5}, {1, 2, 6},
		{2, 0, 7}, {2, 1, 8}, {2, 2, 9},
	}

	for _, tc := range testCases {
		pos := ConvertToPosition(tc.row, tc.col)
		if pos != tc.expectedPosition {
			t.Errorf("(%d,%d): expected position %d, got %d",
				tc.row, tc.col, tc.expectedPosition, pos)
		}
	}
}

// TestConvertToPosition_FirstRow testet alle Koordinaten der ersten Zeile
func TestConvertToPosition_FirstRow(t *testing.T) {
	for col := 0; col < 3; col++ {
		pos := ConvertToPosition(0, col)
		expected := col + 1
		if pos != expected {
			t.Errorf("(0,%d): expected position %d, got %d", col, expected, pos)
		}
	}
}

// TestConvertToPosition_SecondRow testet alle Koordinaten der zweiten Zeile
func TestConvertToPosition_SecondRow(t *testing.T) {
	for col := 0; col < 3; col++ {
		pos := ConvertToPosition(1, col)
		expected := 4 + col
		if pos != expected {
			t.Errorf("(1,%d): expected position %d, got %d", col, expected, pos)
		}
	}
}

// TestConvertToPosition_ThirdRow testet alle Koordinaten der dritten Zeile
func TestConvertToPosition_ThirdRow(t *testing.T) {
	for col := 0; col < 3; col++ {
		pos := ConvertToPosition(2, col)
		expected := 7 + col
		if pos != expected {
			t.Errorf("(2,%d): expected position %d, got %d", col, expected, pos)
		}
	}
}

// TestConvertToPosition_FirstColumn testet alle Koordinaten der ersten Spalte
func TestConvertToPosition_FirstColumn(t *testing.T) {
	expectedPositions := []int{1, 4, 7}
	for row := 0; row < 3; row++ {
		pos := ConvertToPosition(row, 0)
		if pos != expectedPositions[row] {
			t.Errorf("(%d,0): expected position %d, got %d", row, expectedPositions[row], pos)
		}
	}
}

// TestConvertToPosition_SecondColumn testet alle Koordinaten der zweiten Spalte
func TestConvertToPosition_SecondColumn(t *testing.T) {
	expectedPositions := []int{2, 5, 8}
	for row := 0; row < 3; row++ {
		pos := ConvertToPosition(row, 1)
		if pos != expectedPositions[row] {
			t.Errorf("(%d,1): expected position %d, got %d", row, expectedPositions[row], pos)
		}
	}
}

// TestConvertToPosition_ThirdColumn testet alle Koordinaten der dritten Spalte
func TestConvertToPosition_ThirdColumn(t *testing.T) {
	expectedPositions := []int{3, 6, 9}
	for row := 0; row < 3; row++ {
		pos := ConvertToPosition(row, 2)
		if pos != expectedPositions[row] {
			t.Errorf("(%d,2): expected position %d, got %d", row, expectedPositions[row], pos)
		}
	}
}

// TestConvertToPosition_Corners testet alle Eckkoordinaten
func TestConvertToPosition_Corners(t *testing.T) {
	corners := map[[2]int]int{
		{0, 0}: 1, // oben links
		{0, 2}: 3, // oben rechts
		{2, 0}: 7, // unten links
		{2, 2}: 9, // unten rechts
	}

	for coord, expected := range corners {
		pos := ConvertToPosition(coord[0], coord[1])
		if pos != expected {
			t.Errorf("Corner (%d,%d): expected position %d, got %d",
				coord[0], coord[1], expected, pos)
		}
	}
}

// TestConvertToPosition_Edges testet alle Kantenkoordinaten
func TestConvertToPosition_Edges(t *testing.T) {
	edges := map[[2]int]int{
		{0, 1}: 2, // oben mitte
		{1, 0}: 4, // mitte links
		{1, 2}: 6, // mitte rechts
		{2, 1}: 8, // unten mitte
	}

	for coord, expected := range edges {
		pos := ConvertToPosition(coord[0], coord[1])
		if pos != expected {
			t.Errorf("Edge (%d,%d): expected position %d, got %d",
				coord[0], coord[1], expected, pos)
		}
	}
}

// Roundtrip-Tests: Testen, ob die Konvertierung hin und zurück funktioniert

// TestRoundtrip_PositionToCoordinatesToPosition testet Position -> Koordinaten -> Position
func TestRoundtrip_PositionToCoordinatesToPosition(t *testing.T) {
	for originalPos := 1; originalPos <= 9; originalPos++ {
		row, col := ConvertToCoordinates(originalPos)
		resultPos := ConvertToPosition(row, col)
		if resultPos != originalPos {
			t.Errorf("Roundtrip failed: %d -> (%d,%d) -> %d",
				originalPos, row, col, resultPos)
		}
	}
}

// TestRoundtrip_CoordinatesToPositionToCoordinates testet Koordinaten -> Position -> Koordinaten
func TestRoundtrip_CoordinatesToPositionToCoordinates(t *testing.T) {
	for originalRow := 0; originalRow < 3; originalRow++ {
		for originalCol := 0; originalCol < 3; originalCol++ {
			pos := ConvertToPosition(originalRow, originalCol)
			resultRow, resultCol := ConvertToCoordinates(pos)
			if resultRow != originalRow || resultCol != originalCol {
				t.Errorf("Roundtrip failed: (%d,%d) -> %d -> (%d,%d)",
					originalRow, originalCol, pos, resultRow, resultCol)
			}
		}
	}
}

// TestRoundtrip_AllPositions testet vollständigen Roundtrip für alle Positionen
func TestRoundtrip_AllPositions(t *testing.T) {
	for pos := 1; pos <= 9; pos++ {
		row, col := ConvertToCoordinates(pos)
		backToPos := ConvertToPosition(row, col)

		if backToPos != pos {
			t.Errorf("Position %d roundtrip failed: got %d", pos, backToPos)
		}

		// Prüfe auch Rückweg
		backRow, backCol := ConvertToCoordinates(backToPos)
		if backRow != row || backCol != col {
			t.Errorf("Coordinates roundtrip failed: (%d,%d) != (%d,%d)",
				row, col, backRow, backCol)
		}
	}
}

// Grenzwert- und Edge-Case-Tests

// TestConvertToPosition_ValidRange testet, ob alle Positionen im Bereich 1-9 liegen
func TestConvertToPosition_ValidRange(t *testing.T) {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			pos := ConvertToPosition(row, col)
			if pos < 1 || pos > 9 {
				t.Errorf("Position out of range for (%d,%d): %d", row, col, pos)
			}
		}
	}
}

// TestConvertToCoordinates_ValidRange testet, ob alle Koordinaten im Bereich 0-2 liegen
func TestConvertToCoordinates_ValidRange(t *testing.T) {
	for pos := 1; pos <= 9; pos++ {
		row, col := ConvertToCoordinates(pos)
		if row < 0 || row > 2 {
			t.Errorf("Row out of range for position %d: %d", pos, row)
		}
		if col < 0 || col > 2 {
			t.Errorf("Col out of range for position %d: %d", pos, col)
		}
	}
}

// TestConvertToPosition_UniquePositions testet, ob alle Koordinaten eindeutige Positionen ergeben
func TestConvertToPosition_UniquePositions(t *testing.T) {
	positions := make(map[int]bool)
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			pos := ConvertToPosition(row, col)
			if positions[pos] {
				t.Errorf("Duplicate position %d for coordinates (%d,%d)", pos, row, col)
			}
			positions[pos] = true
		}
	}

	if len(positions) != 9 {
		t.Errorf("Expected 9 unique positions, got %d", len(positions))
	}
}

// TestConvertToCoordinates_UniqueCoordinates testet, ob alle Positionen eindeutige Koordinaten ergeben
func TestConvertToCoordinates_UniqueCoordinates(t *testing.T) {
	coordinates := make(map[[2]int]bool)
	for pos := 1; pos <= 9; pos++ {
		row, col := ConvertToCoordinates(pos)
		coord := [2]int{row, col}
		if coordinates[coord] {
			t.Errorf("Duplicate coordinates (%d,%d) for position %d", row, col, pos)
		}
		coordinates[coord] = true
	}

	if len(coordinates) != 9 {
		t.Errorf("Expected 9 unique coordinates, got %d", len(coordinates))
	}
}
