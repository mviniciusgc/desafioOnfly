package models

import "time"

type TravelModel struct {
	ID            *int64    `json:"id,omitempty"`
	RequesterName string    `json:"requesterName"`
	Destination   string    `json:"destination"`
	DepartureDate time.Time `json:"departureDate"`
	ReturnDate    time.Time `json:"returnDate"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
}
