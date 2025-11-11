package shared

// player definiert die Spielertypen
type Player string

const (
	PlayerX Player = "X" // Maximierender Spieler (Mensch/Player 1)
	PlayerO Player = "O" // Minimierender Spieler (Computer/Player 2)
	Empty   Player = " " // Leere Zelle
)

type State int

const (
	InProgress State = iota
	Tie
	Player1WON
	Player2WON
)
