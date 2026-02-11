package models

type Manufacturer struct {
	ID             int     `db:"id"`
	Name           string  `db:"name"`
	Country        *string `db:"country"`
	ContactEmail   *string `db:"contact_email"`
	ContactPhoneID *int    `db:"contact_phone_id"`
	Website        *string `db:"website"`
}
