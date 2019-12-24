package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/mayah/go-repository-sample/model"
	"github.com/mayah/go-repository-sample/repository"
	"github.com/mayah/go-repository-sample/repository/database"
)

func f(r repository.Repository) error {
	con, err := r.NewConnection()
	if err != nil {
		return err
	}

	defer con.Close()

	user, err := con.User().Find(model.UserID(1))
	if err != nil {
		return err
	}

	tasks, err := con.Task().List(repository.TaskFilter{
		UserID: model.UserID(1),
	})
	if err != nil {
		return err
	}

	err = con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.Task().UpdateContent(model.TaskID(1), "new content")
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println(user)
	fmt.Println(tasks)

	return nil
}

func main() {
	var db *gorm.DB //
	r := database.NewRepository(db)

	f(r)
}
