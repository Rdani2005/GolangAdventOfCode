package entities

type GameRecord struct {
	Red   int
	Blue  int
	Green int
}

func (gr *GameRecord) ValidatePartOne() bool {
	return gr.Red <= 12 && gr.Blue <= 14 && gr.Green <= 13
}
