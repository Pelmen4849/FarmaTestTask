package repository

import (
	"context"
	"medical_farm/back/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DrugRepository interface {
	GetAll(ctx context.Context) ([]models.Drug, error)
	GetByID(ctx context.Context, id int) (*models.Drug, error)
	GetAvailableInShop(ctx context.Context, shopID int) ([]models.ShopInventory, error)
	Create(ctx context.Context, drug *models.Drug) error
	Update(ctx context.Context, drug *models.Drug) error
	Delete(ctx context.Context, id int) error
}

type drugRepository struct {
	db *pgxpool.Pool
}

func NewDrugRepository(db *pgxpool.Pool) DrugRepository {
	return &drugRepository{db: db}
}

func (r *drugRepository) GetAll(ctx context.Context) ([]models.Drug, error) {
	query := `SELECT id, name, international_name, manufacturer_id, category_id,
	                 description, dosage_form, dosage, requires_prescription,
	                 storage_conditions, expiry_days, created_at, updated_at
	          FROM drugs`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drugs []models.Drug
	for rows.Next() {
		var d models.Drug
		err := rows.Scan(
			&d.ID, &d.Name, &d.InternationalName, &d.ManufacturerID,
			&d.CategoryID, &d.Description, &d.DosageForm, &d.Dosage,
			&d.RequiresPrescription, &d.StorageConditions, &d.ExpiryDays,
			&d.CreatedAt, &d.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		drugs = append(drugs, d)
	}
	return drugs, nil
}

func (r *drugRepository) GetByID(ctx context.Context, id int) (*models.Drug, error) {
	query := `SELECT id, name, international_name, manufacturer_id, category_id,
	                 description, dosage_form, dosage, requires_prescription,
	                 storage_conditions, expiry_days, created_at, updated_at
	          FROM drugs WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var d models.Drug
	err := row.Scan(
		&d.ID, &d.Name, &d.InternationalName, &d.ManufacturerID,
		&d.CategoryID, &d.Description, &d.DosageForm, &d.Dosage,
		&d.RequiresPrescription, &d.StorageConditions, &d.ExpiryDays,
		&d.CreatedAt, &d.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

func (r *drugRepository) GetAvailableInShop(ctx context.Context, shopID int) ([]models.ShopInventory, error) {
	query := `SELECT 
        si.id, si.shop_id, si.drug_id, si.batch_number,
        si.quantity, si.purchase_price, si.selling_price,
        si.manufacturing_date, si.expiry_date, si.supplier_id,
        si.received_at, si.last_updated,
        d.name AS drug_name,      -- обязательно
        d.dosage_form,           -- может быть NULL
        d.dosage                 -- может быть NULL
    FROM shop_inventory si
    JOIN drugs d ON si.drug_id = d.id
    WHERE si.shop_id = $1 AND si.quantity > 0 AND si.expiry_date > CURRENT_DATE`

	rows, err := r.db.Query(ctx, query, shopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()


	items := make([]models.ShopInventory, 0)

	for rows.Next() {
		var item models.ShopInventory
		err := rows.Scan(
			&item.ID, &item.ShopID, &item.DrugID, &item.BatchNumber,
			&item.Quantity, &item.PurchasePrice, &item.SellingPrice,
			&item.ManufacturingDate, &item.ExpiryDate, &item.SupplierID,
			&item.ReceivedAt, &item.LastUpdated,
			&item.DrugName,
			&item.DosageForm,
			&item.Dosage,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *drugRepository) Create(ctx context.Context, drug *models.Drug) error {
	query := `INSERT INTO drugs (
		name, international_name, manufacturer_id, category_id,
		description, dosage_form, dosage, requires_prescription,
		storage_conditions, expiry_days, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW())
	RETURNING id`
	return r.db.QueryRow(ctx, query,
		drug.Name, drug.InternationalName, drug.ManufacturerID, drug.CategoryID,
		drug.Description, drug.DosageForm, drug.Dosage, drug.RequiresPrescription,
		drug.StorageConditions, drug.ExpiryDays,
	).Scan(&drug.ID)
}

func (r *drugRepository) Update(ctx context.Context, drug *models.Drug) error {
	query := `UPDATE drugs SET
		name = $1, international_name = $2, manufacturer_id = $3, category_id = $4,
		description = $5, dosage_form = $6, dosage = $7, requires_prescription = $8,
		storage_conditions = $9, expiry_days = $10, updated_at = NOW()
		WHERE id = $11`
	_, err := r.db.Exec(ctx, query,
		drug.Name, drug.InternationalName, drug.ManufacturerID, drug.CategoryID,
		drug.Description, drug.DosageForm, drug.Dosage, drug.RequiresPrescription,
		drug.StorageConditions, drug.ExpiryDays, drug.ID,
	)
	return err
}

func (r *drugRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM drugs WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
