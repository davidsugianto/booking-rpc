package booking

import "time"

const (
	BookingTable = "booking"
)

type Booking struct {
	ID        string    `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Status    string    `json:"status" gorm:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}
