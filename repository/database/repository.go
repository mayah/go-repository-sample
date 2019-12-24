package database

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/mayah/go-repository-sample/repository"
)

// NewRepository generates a new repository using DB.
func NewRepository(db *gorm.DB) repository.Repository {
	return &dbRepository{
		db: db,
	}
}

type dbRepository struct {
	db *gorm.DB
}

type dbConnection struct {
	db *gorm.DB
}

type dbTransaction struct {
	db *gorm.DB
}

func (r *dbRepository) NewConnection() (repository.Connection, error) {
	return &dbConnection{
		db: r.db,
	}, nil
}

func (r *dbRepository) MustConnection() repository.Connection {
	con, err := r.NewConnection()
	if err != nil {
		panic(err)
	}

	return con
}

func (con *dbConnection) Close() error {
	// We don't need to close *gorm.DB. No need to do anything.
	return nil
}

func (con *dbConnection) RunTransaction(f func(repository.Transaction) error) error {
	tx := con.db.Begin()

	err := f(&dbTransaction{db: tx})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (con *dbConnection) User() repository.UserQuery {
	return &dbUserRepository{db: con.db}
}
func (con *dbConnection) Task() repository.TaskQuery {
	return &dbTaskRepository{db: con.db}
}

func (tx *dbTransaction) User() repository.UserCommand {
	return &dbUserRepository{db: tx.db}
}
func (tx *dbTransaction) Task() repository.TaskCommand {
	return &dbTaskRepository{db: tx.db}
}

func escapeForLike(searchStr string) string {
	searchStr = strings.Replace(searchStr, "\\", "\\\\", -1)
	searchStr = strings.Replace(searchStr, "%", "\\%", -1)
	searchStr = strings.Replace(searchStr, "_", "\\_", -1)
	return searchStr
}
