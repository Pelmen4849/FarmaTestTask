package models

import "time"

type Prescription struct {
	ID                 int       `db:"id"`
	CustomerID         int       `db:"customer_id"`
	DoctorName         *string   `db:"doctor_name"`
	DoctorLicense      *string   `db:"doctor_license"`
	DrugID             int       `db:"drug_id"`
	DosageInstructions *string   `db:"dosage_instructions"`
	QuantityPrescribed int       `db:"quantity_prescribed"`
	DatePrescribed     time.Time `db:"date_prescribed"`
	DateValidUntil     time.Time `db:"date_valid_until"`
	IsUsed             bool      `db:"is_used"`
	CreatedAt          time.Time `db:"created_at"`
}
