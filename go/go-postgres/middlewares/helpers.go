package middlewares

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/m3rashid/learn_x/go/go-postgres/models"
)

func CreateConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	return db
}

func getStock(id int64) (models.Stock, error) {
	db := CreateConnection()
	defer db.Close()
	sqlSt := "SELECT * FROM stocks WHERE stockid = $1"
	var stock models.Stock
	row := db.QueryRow(sqlSt, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to get stocks: %v", err)
	}

	fmt.Printf("found stock with id %d\n", stock.StockID)
	return stock, nil
}

func getAllStocks() ([]models.Stock, error) {
	db := CreateConnection()
	defer db.Close()
	sqlSt := "SELECT * FROM stocks"
	var stocks []models.Stock
	rows, err := db.Query(sqlSt)
	if err != nil {
		log.Fatalf("Unable to execute query: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			log.Fatalf("Unable to scan row: %v", err)
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}

func insertStock(stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()
	sql := "INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING stockid"
	var id int64
	err := db.QueryRow(sql, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted stock with id %d\n", id)
	return id
}

func updateStock(stock models.Stock) int64 {
	db := CreateConnection()
	defer db.Close()
	sqlSt := `UPDATE stocks SET name = $2, price = $3, company = $4 WHERE stockid = $1`
	res, err := db.Exec(sqlSt, stock.StockID, stock.Name, stock.Price, stock.Company)
	if err != nil {
		log.Fatalf("Unable to execute query: %v", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Unable while checking rows affected: %v", err)
	}
	fmt.Printf("updated %d rows\n", rows)
	return rows
}

func deleteStock(id int64) int64 {
	db := CreateConnection()
	defer db.Close()
	sqlSt := `DELETE from stocks WHERE stockid = $1`
	res, err := db.Exec(sqlSt, id)
	if err != nil {
		log.Fatalf("Unable to execute query: %v", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Unable while checking rows affected: %v", err)
	}
	fmt.Printf("updated %d rows\n", rows)
	return rows
}
