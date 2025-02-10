package maaain

import (
	"billsHotelService/domain/entity"
	"billsHotelService/infrastructure/database"
	"billsHotelService/server/repository"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.NewMySQL()
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
	// CREATE TABLE IF NOT EXISTS hotels (id INT PRIMARY KEY, name VARCHAR(255), price_pernight INT, rooms_available INT);
	// TODO: usecase層へ移動
	// SQL実行のリポジトリを生成する
	hotelRepo := repository.NewMySQLHotelRepository(db)

	insertHotel := entity.NewHotel(2, "secondHotel", 20000, 5)
	err = hotelRepo.HotelSave(*insertHotel)
	if err != nil {
		panic(err)
	}

	secandHotel, err := hotelRepo.HotelGetById(2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(secandHotel)
	}

	err = hotelRepo.HotelDelete(2)
	if err != nil {
		panic(err)
	}

}
