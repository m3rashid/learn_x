package routers

import (
	"github.com/gorilla/mux"
	"github.com/m3rashid/learn_x/go/go-postgres/middlewares"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/stock/{id}", middlewares.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/stock/all", middlewares.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/stock/new", middlewares.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/stock/update/{id}", middlewares.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/stock/delete/{id}", middlewares.DeleteStock).Methods("DELETE", "OPTIONS")
	return router
}
