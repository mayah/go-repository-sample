package repository

import (
	"github.com/mayah/go-repository-sample/model"
)

type Repository interface {
	NewConnection() (Connection, error)
	MustConnection() Connection
}

type Connection interface {
	Close() error
	RunTransaction(func(tx Transaction) error) error

	User() UserQuery
	Task() TaskQuery
}

type Transaction interface {
	User() UserCommand
	Task() TaskCommand
}

type UserQuery interface {
	Find(id model.UserID) (*model.User, error)
	List(filter UserFilter) ([]*model.User, error)
}

type UserCommand interface {
	UserQuery

	UpdateName(id model.UserID, name string) error
}

type UserFilter struct {
	NameLike string
}

type TaskQuery interface {
	Find(id model.TaskID) (*model.Task, error)
	List(filter TaskFilter) ([]*model.Task, error)
}

type TaskCommand interface {
	TaskQuery

	UpdateContent(id model.TaskID, content string) error
}

type TaskFilter struct {
	UserID model.UserID
}
