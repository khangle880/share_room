package query

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/khangle880/share_room/graph/model"
)

type CategoriesRepo struct {
	DB *pg.DB
}

func (r *CategoriesRepo) GetCategories(filter *model.BudgetFilter, limit *int, offset *int) ([]*model.Category, error) {
	var categories []*model.Category
	query := r.DB.Model(&categories).Order("id")

	if filter.Name != nil && *filter.Name != "" {
		query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
	}
	if filter.Description != nil && *filter.Description != "" {
		query.Where("description ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Description))
	}
	if limit != nil {
		query.Limit(*limit)
	}
	if offset != nil {
		query.Offset(*offset)
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoriesRepo) GetCategoryByID(id uuid.UUID) (*model.Category, error) {
	var category model.Category
	err := r.DB.Model(&category).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoriesRepo) CreateCategory(category *model.Category) (*model.Category, error) {
	_, err := r.DB.Model(category).Returning("*").Insert()

	return category, err
}

func (r *CategoriesRepo) DeleteCategory(category *model.Category) error {
	_, err := r.DB.Model(category).Where("id = ?", category.ID).Delete()
	return err;
}
