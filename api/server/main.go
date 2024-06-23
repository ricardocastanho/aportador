package server

import (
	"aportador/principles"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
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

	results, err := principles.GetStocks(stocks)
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

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
