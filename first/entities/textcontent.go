package entities

import (
	"regexp"
	"strconv"

	. "rdani2005.com/GolangAdventageCode/first/helpers"
)

type TextContent struct {
	Texts   []string
	Results []int
}

func (tp *TextContent) extractNumbers(text string) int {
	re := regexp.MustCompile("\\d")
	numbers := re.FindAllString(text, -1)
	if len(numbers) > 0 {
		result, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		return result
	}
	return 0
}

func (tp *TextContent) ProcessTexts() {
	for _, text := range tp.Texts {
		tp.Results = append(tp.Results, tp.extractNumbers(text))
	}
}

func (tp *TextContent) GetResults() int {
	sum := 0
	for _, result := range tp.Results {
		sum += result
	}
	return sum
}

func (tp *TextContent) findAllNumbersWithPositions(line string, numberWords map[string]int) map[string]interface{} {
	result := make(map[string]interface{})

	for word := range numberWords {
		re := regexp.MustCompile(word)
		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			if _, ok := result[word]; !ok {
				// Initialize the slice if it doesn't exist
				result[word] = make([]int, 0)
			}
			// Append to the existing slice
			result[word] = append(result[word].([]int), match[0])
		}
	}

	re := regexp.MustCompile("\\d")
	matches := re.FindAllStringIndex(line, -1)

	for _, match := range matches {
		char := string(line[match[0]])
		if _, ok := result[char]; !ok {
			// Initialize the slice if it doesn't exist
			result[char] = make([]int, 0)
		}
		// Append to the existing slice
		result[char] = append(result[char].([]int), match[0])
	}

	// Process the result map
	for key, value := range result {
		if arr, ok := value.([]int); ok && len(arr) == 1 {
			// If it's a slice with one element, store the element itself
			result[key] = arr[0]
		}
	}

	return result
}

func (tp *TextContent) getWordsResult() []int {
	var listOfNumbers []int
	for _, line := range tp.Texts {
		numDict := tp.findAllNumbersWithPositions(line, NumberWordsToDigits)

		convertedDict := make(map[string]interface{})
		for key, value := range numDict {
			numKey, ok := NumberWordsToDigits[key]
			if !ok {
				numKey, _ = strconv.Atoi(key)
			}

			// Ensure convertedDict[numKey] is initialized as a slice of integers
			if _, ok := convertedDict[strconv.Itoa(numKey)]; !ok {
				convertedDict[strconv.Itoa(numKey)] = []int{}
			}

			// Handle the case where value is not a slice of integers
			if arr, ok := value.([]int); ok {
				convertedDict[strconv.Itoa(numKey)] = append(convertedDict[strconv.Itoa(numKey)].([]int), arr...)
			} else if num, ok := value.(int); ok {
				// If it's an int, append it directly
				convertedDict[strconv.Itoa(numKey)] = append(convertedDict[strconv.Itoa(numKey)].([]int), num)
			}
		}

		minIndexNum, _ := strconv.Atoi(MinKey(convertedDict, func(k string, v []int) int { return Min(v) }))
		maxIndexNum, _ := strconv.Atoi(MaxKey(convertedDict, func(k string, v []int) int { return Max(v) }))

		listOfNumbers = append(listOfNumbers, minIndexNum*10+maxIndexNum)
	}
	return listOfNumbers
}

func (tp *TextContent) ProcessTextsWithWords() {
	tp.Results = tp.getWordsResult()
}
