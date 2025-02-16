package repository

import (
	"billsHotelService/domain/entity"
	"billsHotelService/domain/repository"
	"database/sql"
)

type MySQLHotelRepository struct {
	db *sql.DB
}

// repository.HotelRepository
func NewMySQLHotelRepository(db *sql.DB) repository.HotelRepository {
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

func (r *MySQLHotelRepository) HotelSave(hotel entity.Hotel) error {
	_, err := r.db.Exec(
		"INSERT INTO hotels (id,name,price_perNight,rooms_available) VALUES(?,?,?,?)",
		hotel.ID, hotel.Name, hotel.PricePerNight, hotel.RoomsAvailable,
	)
	return err
}

func (r *MySQLHotelRepository) HotelUpdateById(hotel entity.Hotel) error {
	_, err := r.db.Exec(
		"UPDATE hotels SET name=?, price_perNight=?, rooms_available=? WHERE id=?",
		hotel.Name, hotel.PricePerNight, hotel.RoomsAvailable, hotel.ID,
	)
	return err
}

func (r *MySQLHotelRepository) HotelDelete(id int) error {
	_, err := r.db.Exec("DELETE FROM hotels WHERE id=?", id)
	return err
}
