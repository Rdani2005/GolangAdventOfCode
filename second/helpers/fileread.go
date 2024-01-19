package helpers

import (
	"bufio"
	"fmt"
	"os"

	. "rdani2005.com/GolangAdventageCode/second/entities"
)

// ReadFile lee un archivo y devuelve un proceso de juego.
func ReadFile() GameProcess {
	gameSets := make([]GameSet, 0)

	fileName := "second.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		gameID := GetGameID(line)
		setsColors := GetSets(line)

		records := make([]GameRecord, 0)
		for _, setColor := range setsColors {
			red := setColor["red"]
			green := setColor["green"]
			blue := setColor["blue"]

			records = append(records, GameRecord{Red: red, Green: green, Blue: blue})
		}

		gameSet := GameSet{ID: gameID, Records: records}
		gameSets = append(gameSets, gameSet)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return GameProcess{GameSets: gameSets}
}
