package entity

import "time"

type TravelEntity struct {
	ID            int
	RequesterName string
	Destination   string
	DepartureDate time.Time
	ReturnDate    time.Time
	Status        string
	CreatedAt     time.Time
}
