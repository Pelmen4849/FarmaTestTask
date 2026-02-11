package models

type DrugCategory struct {
	ID               int     `db:"id"`
	CategoryName     string  `db:"category_name"`
	ParentCategoryID *int    `db:"parent_category_id"`
	Description      *string `db:"description"`
}
