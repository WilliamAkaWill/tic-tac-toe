package shared

func ConvertToCoordinates(position int) (int, int) {
	row := (position - 1) / 3
	col := (position - 1) % 3
	return row, col
}

func ConvertToPosition(row int, col int) int {
	return row*3 + col + 1
}