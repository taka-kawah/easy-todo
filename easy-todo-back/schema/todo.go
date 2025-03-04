package schema

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	Id        int64 `gorm:"primaryKey"`
	Value     string
	UserId    uint
	IsDone    bool
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
	DeleteAt  *gorm.DeletedAt `gorm:"index"`
}

type ToDoDriver struct {
	db *gorm.DB
}

func NewToDoDriver(db *gorm.DB) *ToDoDriver {
	db.AutoMigrate(&Todo{})
	return &ToDoDriver{db: db}
}

func (d *ToDoDriver) CreateToDo(value string, userId uint) error {
	id, err := generateNewID()
	if err != nil {
		return err
	}
	todo := Todo{Id: id, Value: value, UserId: userId, IsDone: false}
	if err := d.db.Create(&todo).Error; err != nil {
		return fmt.Errorf("failed to Connect New Todo Record: %w", err)
	}
	return nil
}

func (d *ToDoDriver) ReadToDos() ([]Todo, error) {
	var todos []Todo
	res := d.db.Find(&todos)
	if err := res.Error; err != nil {
		return nil, fmt.Errorf("failed to Read Todo Records %w", err)
	}
	return todos, nil
}

func (d *ToDoDriver) ReadSingleTodoById(id int64) (Todo, error) {
	var todo Todo
	res := d.db.First(&todo, id)
	if err := res.Error; err != nil {
		return Todo{}, fmt.Errorf("failed to Read Single Todo Record")
	}
	return todo, nil
}

func (d *ToDoDriver) UpdateTodoValById(id int64, newVal string) error {
	var todo Todo
	res := d.db.Model(Todo{}).Where("id = ?", id).First(&todo).Update("Value", newVal)
	if err := res.Error; err != nil {
		return fmt.Errorf("failed to Update Todo Record: %w", err)
	}
	return nil
}

func (d *ToDoDriver) ToggleTodoById(id int64) error {
	var todo Todo
	res := d.db.Model(Todo{}).Where("id = ?", id).First(&todo).Update("is_done", gorm.Expr("NOT is_done"))
	if err := res.Error; err != nil {
		return fmt.Errorf("failed to Check Todo Record %w", err)
	}
	return nil
}

func (d *ToDoDriver) DeleteTodoById(id int64) error {
	res := d.db.Delete(&Todo{}, id)
	if err := res.Error; err != nil {
		return fmt.Errorf("failed to Deleting Todo Record %w", err)
	}
	return nil
}

func generateNewID() (int64, error) {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return 0, fmt.Errorf("failed to Generate ID: %w", err)
	}

	return int64(binary.BigEndian.Uint64(b[:])) & 0x7FFFFFFFFFFFFFFF, nil
}
