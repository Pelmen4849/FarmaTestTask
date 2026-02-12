package models

import "time"

type Address struct {
	ID         int       `db:"id"`
	Country    string    `db:"country"`
	City       string    `db:"city"`
	Street     string    `db:"street"`
	Building   string    `db:"building"`
	Apartment  *string   `db:"apartment"`
	PostalCode *string   `db:"postal_code"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
