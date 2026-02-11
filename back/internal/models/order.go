package models

import "time"

type Order struct {
	ID             int        `db:"id"`
	OrderNumber    string     `db:"order_number"`
	CustomerID     int        `db:"customer_id"`
	ShopID         int        `db:"shop_id"`
	EmployeeID     *int       `db:"employee_id"`
	PrescriptionID *int       `db:"prescription_id"`
	TotalAmount    float64    `db:"total_amount"`
	DiscountAmount float64    `db:"discount_amount"`
	FinalAmount    float64    `db:"final_amount"`
	Status         string     `db:"status"`
	PaymentMethod  *string    `db:"payment_method"`
	PaymentStatus  string     `db:"payment_status"`
	Notes          *string    `db:"notes"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at"`
	CompletedAt    *time.Time `db:"completed_at"`
}
