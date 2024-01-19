package entities

import (
	"fmt"
	"strconv"
	"strings"
)

var DIRECTIONS = [][]int{
	{-1, -1}, // TOP LEFT
	{-1, 0},  // TOP
	{-1, 1},  // TOP RIGHT
	{0, -1},  // CENTER LEFT
	{0, 1},   // CENTER RIGHT
	{1, -1},  // BOTTOM LEFT
	{1, 0},   // BOTTOM
	{1, 1},   // BOTTOM RIGHT
}

type Analyzer struct {
	schema    [][]string
	processed map[string]bool
}

func NewAnalyzer(inputSchema string) *Analyzer {
	rows := strings.Split(inputSchema, "\n")
	schema := make([][]string, len(rows))
	processed := make(map[string]bool)

	for i, row := range rows {
		schema[i] = strings.Split(row, "")
	}

	return &Analyzer{schema: schema, processed: processed}
}

func (a *Analyzer) validSymbol(char string) bool {
	_, err := strconv.Atoi(char)
	return err != nil && char != "."
}

func (a *Analyzer) validNumber(row, col int) bool {
	for _, dir := range DIRECTIONS {
		dx, dy := dir[0], dir[1]
		newRow, newCol := row+dx, col+dy
		if newRow >= 0 && newRow < len(a.schema) &&
			newCol >= 0 && newCol < len(a.schema[newRow]) &&
			a.validSymbol(a.schema[newRow][newCol]) {
			return true
		}
	}
	return false
}

func (a *Analyzer) getGearRatio(row, col int) int {
	adjacentNumbers := make(map[int]bool)

	for _, dir := range DIRECTIONS {
		dx, dy := dir[0], dir[1]
		newRow, newCol := row+dx, col+dy
		if newRow >= 0 && newRow < len(a.schema) &&
			newCol >= 0 && newCol < len(a.schema[newRow]) &&
			isDigit(a.schema[newRow][newCol]) {
			adjacentNumbers[a.getFullPartNumber(newRow, newCol)] = true
		}
	}

	if len(adjacentNumbers) == 2 {
		result := 1
		for num := range adjacentNumbers {
			result *= num
		}
		return result
	}
	return 0
}

func (a *Analyzer) CalculateSumOfAllGearRatios() int {
	sum := 0
	for i, row := range a.schema {
		for j, char := range row {
			if char == "*" {
				sum += a.getGearRatio(i, j)
			}
		}
	}
	return sum
}

func (a *Analyzer) CalculateSumNumbers() int {
	result := 0
	for i, row := range a.schema {
		j := 0
		for j < len(row) {
			if isDigit(row[j]) && !a.processed[fmt.Sprintf("%d,%d", i, j)] {
				number, initJ := row[j], j
				for j+1 < len(row) && isDigit(row[j+1]) {
					number += row[j+1]
					j++
				}
				for colIndex := initJ; colIndex <= j; colIndex++ {
					a.processed[fmt.Sprintf("%d,%d", i, colIndex)] = true
					if a.validNumber(i, colIndex) {
						result += toInt(number)
						break
					}
				}
			}
			j++
		}
	}
	return result
}

func (a *Analyzer) getFullPartNumber(row, col int) int {
	numberStr := a.schema[row][col]

	leftCol := col - 1
	for leftCol >= 0 && isDigit(a.schema[row][leftCol]) {
		numberStr = a.schema[row][leftCol] + numberStr
		leftCol--
	}

	rightCol := col + 1
	for rightCol < len(a.schema[row]) && isDigit(a.schema[row][rightCol]) {
		numberStr += a.schema[row][rightCol]
		rightCol++
	}

	return toInt(numberStr)
}

func isDigit(char string) bool {
	_, err := strconv.Atoi(char)
	return err == nil
}

func toInt(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}
