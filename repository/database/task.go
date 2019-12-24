package database

import (
	"github.com/jinzhu/gorm"

	"github.com/mayah/go-repository-sample/model"
	"github.com/mayah/go-repository-sample/repository"
)

type dbTaskRepository struct {
	db *gorm.DB
}

func (r *dbTaskRepository) Find(taskID model.TaskID) (*model.Task, error) {
	db := r.db

	task := &model.Task{
		ID: taskID,
	}

	if err := db.First(task).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return task, nil
}

func (r *dbTaskRepository) List(filter repository.TaskFilter) ([]*model.Task, error) {
	db := r.db

	if filter.UserID != 0 {
		db = db.Where("user_id = ?", filter.UserID)
	}

	var tasks []*model.Task
	err := db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *dbTaskRepository) UpdateContent(taskID model.TaskID, content string) error {
	db := r.db

	target := &model.Task{
		ID: taskID,
	}

	task := map[string]interface{}{
		"content": content,
	}

	err := db.Model(target).Update(task).Error
	if err != nil {
		return err
	}

	return nil
}
