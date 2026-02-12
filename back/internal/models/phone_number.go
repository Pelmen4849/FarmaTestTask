package models

type PhoneNumber struct {
	ID          int    `db:"id"`
	CountryCode string `db:"country_code"`
	Number      string `db:"number"`
}
