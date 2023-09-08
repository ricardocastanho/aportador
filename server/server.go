package server

import (
	"aportador/principles"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, map[string]interface{}{"data": map[string]string{"message": "Hello, World!"}})
}

func GetStocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stocks := r.URL.Query()["stock"]

	if len(stocks) == 0 {
		sendJSONResponse(w, map[string]interface{}{"error": map[string]string{"message": "No stocks provided"}})
		return
	}

	var dividendYield, dividendYears = 10.0, 3.0
	var err error

	dy := r.URL.Query().Get("dividendYield")
	if dy != "" {
		dividendYield, err = strconv.ParseFloat(dy, 64)
		if err != nil {
			sendJSONResponse(w, map[string]interface{}{"error": map[string]string{"message": "Error while parsing dividend yield input"}})
			return
		}
	}

	dys := r.URL.Query().Get("dividendYears")
	if dys != "" {
		dividendYears, err = strconv.ParseFloat(dys, 64)
		if err != nil {
			sendJSONResponse(w, map[string]interface{}{"error": map[string]string{"message": "Error while parsing dividend years input"}})
			return
		}
	}

	results, err := principles.GetStocks(stocks, dividendYield, dividendYears)
	if err != nil {
		fmt.Println("Error getting stocks data:", err)
		sendJSONResponse(w, map[string]interface{}{"error": map[string]string{"message": "Error getting stocks data"}})
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
