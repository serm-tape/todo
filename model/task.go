package model

import (
	"github.com/serm-tape/todo/core"
	"github.com/jinzhu/gorm"

	"database/sql/driver"
)

type Task struct{
	gorm.Model
	Subject string `gorm:"not null"`
	Content string `gorm:"size:500"`
	Status TaskStatus `gorm:"not null" sql:"type:ENUM('PENDING', 'DONE');default:'PENDING'"`
}

func (t Task)TableName() string{
	return "tasks"
}

func GetAllTask() []Task{
	var tasks []Task
	core.App.Database.Find(&tasks)
	return tasks
}

func GetTaskById(id uint) Task{
	var task Task
	core.App.Database.Where("id = ?", id).Find(&task)
	return task
}

func CreateTask(t *Task) error{
	query := core.App.Database.Create(&t)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func EditTask(t *Task) error{
	query := core.App.Database.
		Model(&Task{}).
		Where(Task{Model:gorm.Model{ID:t.ID}}).
		Updates(&t)

	if query.Error != nil {
		return query.Error
	}

	core.App.Database.Where("id = ?", t.ID).Find(&t)

	return nil
}

func DeleteTask(id uint) error{
	query := core.App.Database.
		Where("id = ?", id).
		Delete(&Task{})

	if query.Error != nil {
		return query.Error
	}
	return nil
}

//enum mapping
type TaskStatus string
const (
	pending TaskStatus = "PENDING"
	done TaskStatus = "DONE"
)

func (s *TaskStatus) Scan(value interface{}) error{
	*s = TaskStatus(string(value.([]byte)))
	return nil
}
func (s TaskStatus) Value() (driver.Value, error){
	return string(s), nil
}