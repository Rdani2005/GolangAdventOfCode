package helpers

import (
	"bufio"
	"os"
)

func ReadDataFile() ([]string, error) {
	// Try to open file name "first.txt"
	file, err := os.Open("first.txt")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Create scanner for current file lines
	scanner := bufio.NewScanner(file)
	var lines []string

	// read all lines from file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// verify Errors on lecture
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
