package repository

import (
	"context"
	"medical_farm/back/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *models.Order, items []models.OrderItem) error
	GetByID(ctx context.Context, id int) (*models.Order, error)
	GetOrdersByCustomer(ctx context.Context, customerID int) ([]models.Order, error)
	UpdateStatus(ctx context.Context, orderID int, status string) error
}

type orderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) OrderRepository {
	return &orderRepository{db: db}
}

// CreateOrder создаёт заказ и списывает остатки в одной транзакции
func (r *orderRepository) CreateOrder(ctx context.Context, order *models.Order, items []models.OrderItem) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Вставка заказа
	orderQuery := `INSERT INTO orders (
		order_number, customer_id, shop_id, employee_id, prescription_id,
		total_amount, discount_amount, final_amount, status, payment_method,
		payment_status, notes, created_at, updated_at, completed_at
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)
	RETURNING id`
	err = tx.QueryRow(ctx, orderQuery,
		order.OrderNumber, order.CustomerID, order.ShopID, order.EmployeeID, order.PrescriptionID,
		order.TotalAmount, order.DiscountAmount, order.FinalAmount, order.Status,
		order.PaymentMethod, order.PaymentStatus, order.Notes,
		order.CreatedAt, order.UpdatedAt, order.CompletedAt,
	).Scan(&order.ID)
	if err != nil {
		return err
	}

	// Вставка позиций
	itemQuery := `INSERT INTO order_items (
		order_id, inventory_id, drug_id, quantity,
		unit_price, discount_percent, total_price, created_at
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	for _, item := range items {
		_, err := tx.Exec(ctx, itemQuery,
			order.ID, item.InventoryID, item.DrugID, item.Quantity,
			item.UnitPrice, item.DiscountPercent, item.TotalPrice, item.CreatedAt,
		)
		if err != nil {
			return err
		}
	}

	// Списание остатков
	updateInvQuery := `UPDATE shop_inventory 
	                   SET quantity = quantity - $1, last_updated = NOW()
	                   WHERE id = $2 AND quantity >= $1`
	for _, item := range items {
		cmd, err := tx.Exec(ctx, updateInvQuery, item.Quantity, item.InventoryID)
		if err != nil {
			return err
		}
		if cmd.RowsAffected() == 0 {
			return ErrInsufficientStock
		}
	}

	return tx.Commit(ctx)
}

// GetByID возвращает заказ по ID
func (r *orderRepository) GetByID(ctx context.Context, id int) (*models.Order, error) {
	query := `SELECT id, order_number, customer_id, shop_id, employee_id,
	                 prescription_id, total_amount, discount_amount, final_amount,
	                 status, payment_method, payment_status, notes,
	                 created_at, updated_at, completed_at
	          FROM orders WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var order models.Order
	err := row.Scan(
		&order.ID, &order.OrderNumber, &order.CustomerID, &order.ShopID,
		&order.EmployeeID, &order.PrescriptionID, &order.TotalAmount,
		&order.DiscountAmount, &order.FinalAmount, &order.Status,
		&order.PaymentMethod, &order.PaymentStatus, &order.Notes,
		&order.CreatedAt, &order.UpdatedAt, &order.CompletedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}

// GetOrdersByCustomer возвращает все заказы клиента
func (r *orderRepository) GetOrdersByCustomer(ctx context.Context, customerID int) ([]models.Order, error) {
	query := `SELECT id, order_number, customer_id, shop_id, employee_id,
	                 prescription_id, total_amount, discount_amount, final_amount,
	                 status, payment_method, payment_status, notes,
	                 created_at, updated_at, completed_at
	          FROM orders WHERE customer_id = $1
	          ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		err := rows.Scan(
			&o.ID, &o.OrderNumber, &o.CustomerID, &o.ShopID,
			&o.EmployeeID, &o.PrescriptionID, &o.TotalAmount,
			&o.DiscountAmount, &o.FinalAmount, &o.Status,
			&o.PaymentMethod, &o.PaymentStatus, &o.Notes,
			&o.CreatedAt, &o.UpdatedAt, &o.CompletedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

// UpdateStatus изменяет статус заказа
func (r *orderRepository) UpdateStatus(ctx context.Context, orderID int, status string) error {
	query := `UPDATE orders SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(ctx, query, status, orderID)
	return err
}
