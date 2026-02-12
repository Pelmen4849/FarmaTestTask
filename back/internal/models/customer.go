package models

import "time"

type Customer struct {
	ID        int        `db:"id"`
	UserID    *int       `db:"user_id"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Email     string     `db:"email"`
	PhoneID   *int       `db:"phone_id"`
	AddressID *int       `db:"address_id"`
	BirthDate *time.Time `db:"birth_date"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
