package models

type BookingDTO struct {
	ID      int    `json:"id"`
	User    string `json:"user"`
	Members int    `json:"members"`
}

type Booking struct {
	ID      int
	User    string
	Members int
}

//BookingFromDTO creates a new booking from the given dto
func BookingFromDTO(dto BookingDTO) Booking {
	var booking Booking
	booking.ID = dto.ID
	booking.User = dto.User
	booking.Members = dto.Members
	return booking
}

//ToDTO creates a new dto for the Booking
func (booking Booking) ToDTO() BookingDTO {
	var dto BookingDTO
	dto.ID = booking.ID
	dto.User = booking.User
	dto.Members = booking.Members
	return dto
}
