package entity

import "time"

type Travel struct {
	ID            int       `json:"id"`
	RequesterName string    `json:"requester_name"`
	Destination   string    `json:"destination"`
	DepartureDate time.Time `json:"departure_date"`
	ReturnDate    time.Time `json:"return_date"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
