package models

import "time"

type ShopInventory struct {
	ID                int        `db:"id" json:"id"`
	ShopID            int        `db:"shop_id" json:"shop_id"`
	DrugID            int        `db:"drug_id" json:"drug_id"`
	BatchNumber       *string    `db:"batch_number" json:"batch_number,omitempty"`
	Quantity          int        `db:"quantity" json:"quantity"`
	PurchasePrice     float64    `db:"purchase_price" json:"purchase_price"`
	SellingPrice      float64    `db:"selling_price" json:"selling_price"`
	ManufacturingDate *time.Time `db:"manufacturing_date" json:"manufacturing_date,omitempty"`
	ExpiryDate        time.Time  `db:"expiry_date" json:"expiry_date"`
	SupplierID        *int       `db:"supplier_id" json:"supplier_id,omitempty"`
	ReceivedAt        time.Time  `db:"received_at" json:"received_at"`
	LastUpdated       time.Time  `db:"last_updated" json:"last_updated"`
	DrugName          string     `db:"drug_name" json:"drug_name"`
	DosageForm        *string    `db:"dosage_form" json:"dosage_form,omitempty"`
	Dosage            *string    `db:"dosage" json:"dosage,omitempty"`
}
