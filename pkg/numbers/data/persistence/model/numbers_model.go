package model

import "time"

//Reservation ..
type Reservation struct {
	ID        int64
	Client    string
	Number    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

//TableName name of table
func (r Reservation) TableName() string {
	return "reservation"
}
