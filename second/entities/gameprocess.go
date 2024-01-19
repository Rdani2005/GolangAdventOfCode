package entities

type GameProcess struct {
	GameSets []GameSet
}

func (gp *GameProcess) ProcessValidGames() int {
	sum := 0
	for _, gameSet := range gp.GameSets {
		if gameSet.ValidateGame() {
			sum += gameSet.ID
		}
	}
	return sum
}

func (gp *GameProcess) GetSetsPowerTotal() int {
	sum := 0
	for _, gameSet := range gp.GameSets {
		sum += gameSet.MultiplyColors()
	}
	return sum
}
