package server

import (
	"aportador/principles"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomResponse struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response := CustomResponse{Data: data, Error: nil}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sendErrorResponse(w http.ResponseWriter, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response := CustomResponse{Data: nil, Error: err}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, map[string]string{"message": "Hello, World!"})
}

func GetStocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stocks := r.URL.Query()["stock"]

	if len(stocks) == 0 {
		sendErrorResponse(w, map[string]string{"message": "No stocks provided"})
		return
	}

	var dividendYield, dividendHistory = 10.0, 3.0
	var err error

	dy := r.URL.Query().Get("dividendYield")
	if dy != "" {
		dividendYield, err = strconv.ParseFloat(dy, 64)
		if err != nil {
			sendErrorResponse(w, map[string]string{"message": "Error while parsing dividend yield input"})
			return
		}
	}

	dys := r.URL.Query().Get("dividendHistory")
	if dys != "" {
		dividendHistory, err = strconv.ParseFloat(dys, 64)
		if err != nil {
			sendErrorResponse(w, map[string]string{"message": "Error while parsing dividend years input"})
			return
		}
	}

	results, err := principles.GetStocks(stocks, dividendYield, dividendHistory)
	if err != nil {
		fmt.Println("Error getting stocks data:", err)
		sendErrorResponse(w, map[string]string{"message": "Error getting stocks data"})
		return
	}

	sendJSONResponse(w, results)
}

func CreateServer(port string) error {
	r := mux.NewRouter()
	r.HandleFunc("/", HelloWorldHandler).Methods("GET")
	r.HandleFunc("/stocks", GetStocksHandler).Methods("GET")

	fmt.Printf("Listening on port %s...\n", port)

	return http.ListenAndServe(port, r)
}
