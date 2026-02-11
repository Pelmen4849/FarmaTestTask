package models

import "time"

type Shop struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	AddressID    int       `db:"address_id"`
	PhoneID      *int      `db:"phone_id"`
	Email        *string   `db:"email"`
	OpeningHours *string   `db:"opening_hours"`
	IsActive     bool      `db:"is_active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
