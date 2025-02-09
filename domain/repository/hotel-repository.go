package repository

import "billsHotelService/domain/entity"

type HotelRepository interface {
	HotelGetById(id int) (*entity.Hotel, error)
}
