package models

import "time"

type Employee struct {
	ID        int        `db:"id"`
	UserID    *int       `db:"user_id"`
	RoleID    int        `db:"role_id"`
	ShopID    *int       `db:"shop_id"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Email     string     `db:"email"`
	PhoneID   *int       `db:"phone_id"`
	AddressID *int       `db:"address_id"`
	BirthDate *time.Time `db:"birth_date"`
	HireDate  time.Time  `db:"hire_date"`
	Salary    *float64   `db:"salary"`
	IsActive  bool       `db:"is_active"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
