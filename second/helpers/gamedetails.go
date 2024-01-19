package helpers

import (
	"regexp"
	"strconv"
	"strings"
)

func GetGameID(entry string) int {
	pattern := regexp.MustCompile(`Game (\d+):`)
	matches := pattern.FindStringSubmatch(entry)

	if len(matches) > 1 {
		gameID, _ := strconv.Atoi(matches[1])
		return gameID
	}
	return 0
}

// GetSets extrae la información de conjuntos de colores desde una entrada.
func GetSets(entry string) []map[string]int {
	greenPattern := regexp.MustCompile(`(\d+)\s+green`)
	redPattern := regexp.MustCompile(`(\d+)\s+red`)
	bluePattern := regexp.MustCompile(`(\d+)\s+blue`)

	results := strings.Split(entry, "; ")
	sets := make([]map[string]int, 0)

	for _, result := range results {
		set := make(map[string]int)

		if greenMatch := greenPattern.FindStringSubmatch(result); len(greenMatch) > 1 {
			set["green"], _ = strconv.Atoi(greenMatch[1])
		}
		if redMatch := redPattern.FindStringSubmatch(result); len(redMatch) > 1 {
			set["red"], _ = strconv.Atoi(redMatch[1])
		}
		if blueMatch := bluePattern.FindStringSubmatch(result); len(blueMatch) > 1 {
			set["blue"], _ = strconv.Atoi(blueMatch[1])
		}

		sets = append(sets, set)
	}

	return sets
}
