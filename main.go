package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	. "rdani2005.com/GolangAdventageCode/first/handlers"
	. "rdani2005.com/GolangAdventageCode/second/handlers"
	. "rdani2005.com/GolangAdventageCode/third/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// Day 1 challenge
	router.HandleFunc("/api/v1/dayone/readResultsByNumbersOnly", HandleReadResultsByNumbersOnly).Methods("GET")
	router.HandleFunc("/api/v1/dayone/readResultsByWordsAndNumbers", HandleReadResultsByWordsAndNumbers).Methods("GET")
	// Day 2 Challenge
	router.HandleFunc("/api/v1/daytwo/readValidGames", HandleValidGames).Methods("GET")
	router.HandleFunc("/api/v1/daytwo/readTotalPower", HandleTotalPower).Methods("GET")
	// Day 3 Challenge
	router.HandleFunc("/api/v1/daythree/readByNumber", HandleReadSumNumbers).Methods("GET")
	router.HandleFunc("/api/v1/daythree/readByGear", HandleReadAllRatiosSum).Methods("GET")

	fmt.Println("Server started on port ", 8080)
	log.Fatal(http.ListenAndServe(":3000", router))
}
