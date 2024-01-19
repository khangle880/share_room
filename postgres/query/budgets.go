package query

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/khangle880/share_room/graph/model"
)

type BudgetRepo struct {
	DB *pg.DB
}

func (r *BudgetRepo) GetBudgets() ([]*model.Budget, error) {
	var budgets []*model.Budget
	err := r.DB.Model(&budgets).Order("id").Select()
	if err != nil {
		return nil, err
	}
	return budgets, nil
}

func (r *BudgetRepo) GetBudgetByID(id uuid.UUID) (*model.Budget, error) {
	var budget model.Budget
	err := r.DB.Model(&budget).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *BudgetRepo) UpdateBudget(budget *model.Budget) (*model.Budget, error) {
	_, err := r.DB.Model(budget).Where("id = ?", budget.ID).UpdateNotZero()
	return budget, err
}

func (r *BudgetRepo) CreateBudget(budget *model.Budget) (*model.Budget, error) {
	_, err := r.DB.Model(budget).Returning("*").Insert()
	if err != nil {
		fmt.Println(err)
	}

	return budget, err
}

func (r *BudgetRepo) DeleteBudget(budget *model.Budget) error {
	_, err := r.DB.Model(budget).Where("id = ?", budget.ID).Delete()
	return err
}
