package main

import (
	"billsHotelService/infrastructure"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:password@tcp(localhost:3306)/hotel_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// テーブル作成
	createTableSQL := `CREATE TABLE IF NOT EXISTS hotels (
        id INT PRIMARY KEY,
        name VARCHAR(255),
        price_pernight INT,
    	rooms_available INT);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	// リポジトリを生成する
	hotelRepo := infrastructure.NewMySQLHotelRepository(db)
	hotel, err := hotelRepo.HotelGetById(1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(hotel)
	}

}
