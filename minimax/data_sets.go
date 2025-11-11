package minimax

// move repräsentiert einen möglichen Zug mit seiner Position und Bewertung
type (
	move struct {
		Row   int // Zeile (0-2)
		Col   int // Spalte (0-2)
		Score int // Bewertung des Zugs (-10 bis +10)
	}
)

// Bewertungskonstanten für den Minimax-Algorithmus
const (
	winScore  = 10  // Punktzahl für einen Gewinn
	loseScore = -10 // Punktzahl für eine Niederlage
	tieScore  = 0   // Punktzahl für ein Unentschieden
)