package query

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/khangle880/share_room/graph/model"
)

type UsersRepo struct {
	DB *pg.DB
}

func (r *UsersRepo) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	err := r.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}

func (r *UsersRepo) GetUserByID(id uuid.UUID) (*model.User, error) {
	return r.GetUserByField("id", id.String())
}

func (r *UsersRepo) GetUserByEmail(email string) (*model.User, error) {
	return r.GetUserByField("email", email)
}

func (r *UsersRepo) GetUserByUsername(username string) (*model.User, error) {
	return r.GetUserByField("username", username)
}

func (r *UsersRepo) GetUsers() ([]*model.User, error) {
	var users []*model.User

	err := r.DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (r *UsersRepo) CreateUser(tx *pg.Tx,user *model.User) (*model.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()

	return user, err
}

func (r *UsersRepo) UpdateUser(tx *pg.Tx,user *model.User) (*model.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()

	return user, err
}

func (r *UsersRepo) DeleteUser(user *model.User) error {
	_, err := r.DB.Model(user).Where("id = ?", user.ID).Delete()
	return err
}
