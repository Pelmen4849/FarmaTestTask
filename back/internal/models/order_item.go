package models

import "time"

type OrderItem struct {
	ID              int       `db:"id"`
	OrderID         int       `db:"order_id"`
	InventoryID     int       `db:"inventory_id"`
	DrugID          int       `db:"drug_id"`
	Quantity        int       `db:"quantity"`
	UnitPrice       float64   `db:"unit_price"`
	DiscountPercent float64   `db:"discount_percent"`
	TotalPrice      float64   `db:"total_price"`
	CreatedAt       time.Time `db:"created_at"`
}
