package service

import (
	"context"
	"fmt"
	"medical_farm/back/internal/models"
	"medical_farm/back/internal/repository"
	"time"

	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(ctx context.Context, customerID int, shopID int, items []OrderItemRequest) (*models.Order, error)
	GetOrderByID(ctx context.Context, id int) (*models.Order, error)
	GetCustomerOrders(ctx context.Context, customerID int) ([]models.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID int, status string) error
}

type orderService struct {
	orderRepo repository.OrderRepository
	drugRepo  repository.DrugRepository

}

type OrderItemRequest struct {
	InventoryID int `json:"inventory_id"`
	Quantity    int `json:"quantity"`
}

func NewOrderService(
	orderRepo repository.OrderRepository,
	drugRepo repository.DrugRepository,
) OrderService {
	return &orderService{
		orderRepo: orderRepo,
		drugRepo:  drugRepo,
	}
}

// CreateOrder создаёт новый заказ
func (s *orderService) CreateOrder(ctx context.Context, customerID int, shopID int, items []OrderItemRequest) (*models.Order, error) {
	// Генерируем уникальный номер заказа
	orderNumber := uuid.New().String()[:8]

	var total float64
	var orderItems []models.OrderItem

	for _, req := range items {
		

		item := models.OrderItem{
			InventoryID: req.InventoryID,
			Quantity:    req.Quantity,
			UnitPrice:   100.00, // заглушка
			TotalPrice:  100.00 * float64(req.Quantity),
			CreatedAt:   time.Now(),
		}
		orderItems = append(orderItems, item)
		total += item.TotalPrice
	}

	order := &models.Order{
		OrderNumber:    orderNumber,
		CustomerID:     customerID,
		ShopID:         shopID,
		TotalAmount:    total,
		DiscountAmount: 0,
		FinalAmount:    total,
		Status:         "pending",
		PaymentStatus:  "unpaid",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := s.orderRepo.CreateOrder(ctx, order, orderItems)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}
	return order, nil
}

// GetOrderByID возвращает заказ по ID
func (s *orderService) GetOrderByID(ctx context.Context, id int) (*models.Order, error) {
	return s.orderRepo.GetByID(ctx, id)
}

// GetCustomerOrders возвращает все заказы клиента
func (s *orderService) GetCustomerOrders(ctx context.Context, customerID int) ([]models.Order, error) {
	return s.orderRepo.GetOrdersByCustomer(ctx, customerID)
}

// UpdateOrderStatus обновляет статус заказа
func (s *orderService) UpdateOrderStatus(ctx context.Context, orderID int, status string) error {
	return s.orderRepo.UpdateStatus(ctx, orderID, status)
}
