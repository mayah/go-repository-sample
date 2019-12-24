package database

import (
	"github.com/jinzhu/gorm"

	"github.com/mayah/go-repository-sample/model"
	"github.com/mayah/go-repository-sample/repository"
)

type dbUserRepository struct {
	db *gorm.DB
}

func (r *dbUserRepository) Find(userID model.UserID) (*model.User, error) {
	db := r.db

	user := &model.User{
		ID: userID,
	}

	if err := db.First(user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *dbUserRepository) List(filter repository.UserFilter) ([]*model.User, error) {
	db := r.db

	if filter.NameLike != "" {
		db = db.Where("name LIKE ?", "%"+escapeForLike(filter.NameLike)+"%")
	}

	var users []*model.User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *dbUserRepository) UpdateName(userID model.UserID, name string) error {
	db := r.db

	target := &model.User{
		ID: userID,
	}

	user := map[string]interface{}{
		"name": name,
	}

	err := db.Model(target).Update(user).Error
	if err != nil {
		return err
	}

	return nil
}
