package entity

type Hotel struct {
	ID             int
	Name           string
	PricePerNight  int
	RoomsAvailable int
	//Address        Address
}

//type Address struct {
//	City    string
//	Details string
//}

func NewHotel(id int, name string, pricePerNight int, roomsAvailable int /*, address Address*/) *Hotel {
	return &Hotel{
		ID: id, Name: name,
		PricePerNight:  pricePerNight,
		RoomsAvailable: roomsAvailable,
		//Address:        address,
	}
}

//func NewAddress(city string, details string) *Address {
//	return &Address{
//		City:    city,
//		Details: details,
//	}
//}
