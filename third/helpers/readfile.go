package helpers

import (
	"bufio"
	"fmt"
	"os"

	. "rdani2005.com/GolangAdventageCode/third/entities"
)

func ProcessFile() (*Analyzer, error) {
	filename := "third.txt"
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error Opening the file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	schema := ""
	for _, line := range lines {
		schema += line + "\n"
	}

	analyzer := NewAnalyzer(schema)
	return analyzer, nil
}
