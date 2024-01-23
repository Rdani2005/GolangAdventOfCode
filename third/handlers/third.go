package handlers

import (
	"encoding/json"
	"net/http"

	. "rdani2005.com/GolangAdventageCode/third/helpers"
)

func HandleReadAllRatiosSum(w http.ResponseWriter, r *http.Request) {
	analyzer, err := ProcessFile()

	if err != nil {
		http.Error(w, "Error reading the data", http.StatusInternalServerError)
		return
	}

	result := analyzer.CalculateSumOfAllGearRatios()

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

func HandleReadSumNumbers(w http.ResponseWriter, r *http.Request) {
	analyzer, err := ProcessFile()

	if err != nil {
		http.Error(w, "Error reading the data", http.StatusInternalServerError)
		return
	}

	result := analyzer.CalculateSumNumbers()

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
