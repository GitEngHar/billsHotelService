package repository

import "billsHotelService/domain/entity"

type HotelRepository interface {
	HotelSave(hotel entity.Hotel) error
	HotelGetById(id int) (*entity.Hotel, error)
	HotelUpdateById(hotel entity.Hotel) error
	HotelDelete(id int) error
}
