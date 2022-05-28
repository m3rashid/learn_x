package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/m3rashid/learn_x/go/go-postgres/models"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(stock)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(stocks)
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal(err)
	}
	insertId := insertStock(stock)
	res := response{
		ID:      insertId,
		Message: "Stock created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal(err)
	}
	stock.StockID = id
	updatedRows := updateStock(stock)
	msg := fmt.Sprintf("%d rows updated", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	deletedRows := deleteStock(int64(id))
	msg := fmt.Sprintf("%d rows deleted", deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}
