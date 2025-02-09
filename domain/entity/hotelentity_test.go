package entity_test

import (
	"billsHotelService/domain/entity"
	"strconv"
	"testing"
)

func TestNewHotel(t *testing.T) {
	hotel := entity.NewHotel(1, "hoge", 10, 10)
	expect := struct {
		ID              int
		Name            string
		PricePerNight   int
		RoomsAvaillable int
	}{1, "hoge", 10, 10}

	if hotel.ID != expect.ID {
		t.Errorf("failed ID test now=%s, expacted=%s", strconv.Itoa(hotel.ID), strconv.Itoa(expect.ID))
	}
	if hotel.Name != expect.Name {
		t.Errorf("failed ID Name now=%s, expacted=%s", hotel.Name, expect.Name)
	}
	if hotel.PricePerNight != expect.PricePerNight {
		t.Errorf("failed ID PricePerNight now=%s, expacted=%s", strconv.Itoa(hotel.PricePerNight), strconv.Itoa(expect.PricePerNight))
	}
	if hotel.RoomsAvailable != expect.RoomsAvaillable {
		t.Errorf("failed ID RoomsAvailable now=%s, expacted=%s", strconv.Itoa(hotel.RoomsAvailable), strconv.Itoa(expect.RoomsAvaillable))
	}
}
