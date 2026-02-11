package models

import "time"

type Drug struct {
	ID                   int       `db:"id"`
	Name                 string    `db:"name"`
	InternationalName    *string   `db:"international_name"`
	ManufacturerID       int       `db:"manufacturer_id"`
	CategoryID           *int      `db:"category_id"`
	Description          *string   `db:"description"`
	DosageForm           *string   `db:"dosage_form"`
	Dosage               *string   `db:"dosage"`
	RequiresPrescription bool      `db:"requires_prescription"`
	StorageConditions    *string   `db:"storage_conditions"`
	ExpiryDays           *int      `db:"expiry_days"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}
