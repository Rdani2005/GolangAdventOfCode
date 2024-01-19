package handlers

import (
	"encoding/json"
	"net/http"

	. "rdani2005.com/GolangAdventageCode/second/helpers"
)

func HandleValidGames(w http.ResponseWriter, r *http.Request) {
	game := ReadFile()
	result := game.ProcessValidGames()
	response := map[string]int{
		"result": result,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error converting json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func HandleTotalPower(w http.ResponseWriter, r *http.Request) {
	game := ReadFile()
	result := game.GetSetsPowerTotal()
	response := map[string]int{
		"result": result,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error converting json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
