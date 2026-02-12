package repository

import (
	"context"
	"medical_farm/back/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type InventoryRepository interface {
	GetByID(ctx context.Context, id int) (*models.ShopInventory, error)
}

type inventoryRepository struct {
	db *pgxpool.Pool
}

func NewInventoryRepository(db *pgxpool.Pool) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) GetByID(ctx context.Context, id int) (*models.ShopInventory, error) {
	query := `SELECT id, shop_id, drug_id, batch_number, quantity, 
	                 purchase_price, selling_price, manufacturing_date,
	                 expiry_date, supplier_id, received_at, last_updated,
	                 drug_name, dosage_form, dosage
	          FROM shop_inventory WHERE id = $1`

	var item models.ShopInventory
	err := r.db.QueryRow(ctx, query, id).Scan(
		&item.ID, &item.ShopID, &item.DrugID, &item.BatchNumber,
		&item.Quantity, &item.PurchasePrice, &item.SellingPrice,
		&item.ManufacturingDate, &item.ExpiryDate, &item.SupplierID,
		&item.ReceivedAt, &item.LastUpdated,
		&item.DrugName, &item.DosageForm, &item.Dosage,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
