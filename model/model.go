package model

type UserID int
type User struct {
	ID   UserID
	Name string
}

type TaskID int
type Task struct {
	ID      TaskID
	UserID  UserID
	Content string
}
