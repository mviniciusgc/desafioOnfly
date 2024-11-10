package models

import "time"

type TravelModel struct {
	RequesterName string    `json:"requesterName"`
	Destination   string    `json:"destination"`
	DepartureDate time.Time `json:"departureDate"`
	ReturnDate    time.Time `json:"returnDate"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
}
