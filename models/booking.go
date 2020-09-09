package models

type Booking struct {
	ID      int    `json:"id"`
	User    string `json:"user"`
	Members int    `json:"members"`
}
