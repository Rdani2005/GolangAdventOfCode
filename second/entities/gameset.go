package entities

import "math"

type GameSet struct {
	ID      int
	Records []GameRecord
}

func (gs *GameSet) ValidateGame() bool {
	for _, record := range gs.Records {
		if !record.ValidatePartOne() {
			return false
		}
	}
	return true
}

func (gs *GameSet) ValidateMinRequiredCubes() map[string]int {
	values := map[string]int{"Green": 0, "Red": 0, "Blue": 0}

	for _, record := range gs.Records {
		values["Green"] = int(math.Max(float64(values["Green"]), float64(record.Green)))
		values["Red"] = int(math.Max(float64(values["Red"]), float64(record.Red)))
		values["Blue"] = int(math.Max(float64(values["Blue"]), float64(record.Blue)))
	}

	return values
}

func (gs *GameSet) MultiplyColors() int {
	values := gs.ValidateMinRequiredCubes()
	total := values["Green"] * values["Red"] * values["Blue"]
	return total
}
