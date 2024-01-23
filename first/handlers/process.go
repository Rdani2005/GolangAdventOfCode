package handlers

import (
	"encoding/json"
	"net/http"

	. "rdani2005.com/GolangAdventageCode/first/helpers"
	. "rdani2005.com/GolangAdventageCode/first/mapper"
)

func HandleReadResultsByNumbersOnly(w http.ResponseWriter, r *http.Request) {
	texts, err := ReadDataFile()
	if err != nil {
		http.Error(w, "Error reading the data", http.StatusInternalServerError)
		return
	}
	textContent := NewTextContent(texts)
	textContent.ProcessTexts()
	result := textContent.GetResults()

	response := map[string]int{
		"result": result,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error converting json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func HandleReadResultsByWordsAndNumbers(w http.ResponseWriter, r *http.Request) {
	texts, err := ReadDataFile()
	if err != nil {
		http.Error(w, "Error reading the data", http.StatusInternalServerError)
		return
	}
	textContent := NewTextContent(texts)
	textContent.ProcessTextsWithWords()
	result := textContent.GetResults()

	response := map[string]int{
		"result": result,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error converting json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
