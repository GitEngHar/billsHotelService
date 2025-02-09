package main

import (
	"billsHotelService/domain/entity"
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
	// TODO: usecase層へ移動
	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	// TODO: usecase層へ移動
	// SQL実行のリポジトリを生成する
	hotelRepo := infrastructure.NewMySQLHotelRepository(db)
	hotel, err := hotelRepo.HotelGetById(1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(hotel)
	}

	insertHotel := entity.NewHotel(2, "secondHotel", 20000, 5)
	err = hotelRepo.HotelSave(*insertHotel)
	if err != nil {
		panic(err)
	}

	secondHotel, err := hotelRepo.HotelGetById(2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(secondHotel)
	}
}
