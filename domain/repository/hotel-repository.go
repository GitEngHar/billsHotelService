package repository

import (
	"billsHotelService/domain/entity"
	"database/sql"
)

type MySQLHotelRepository struct {
	db *sql.DB
}

func NewMySQLHotelRepository(db *sql.DB) *MySQLHotelRepository {
	return &MySQLHotelRepository{db: db}
}

func (r *MySQLHotelRepository) HotelGetById(id int) (*entity.Hotel, error) {
	// ホテルを実体化
	hotel := &entity.Hotel{}
	// SQLを実行し結果をhotelに代入
	err := r.db.QueryRow("SELECT id, name, price_pernight, rooms_available FROM hotels WHERE id = ?", id).Scan(&hotel.ID, &hotel.Name, &hotel.PricePerNight, &hotel.RoomsAvailable)
	// エラーの場合は nilとエラーの内容を返す
	if err != nil {
		return nil, err
	}
	return hotel, nil
}

//func (r *MySQLHotelRepository) HotelInsert(hotel entity.Hotel) {}
