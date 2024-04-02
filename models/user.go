package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserManager struct {
	db *gorm.DB
}

func NewUserManager(db *gorm.DB) *UserManager {
	return &UserManager{
		db: db,
	}
}

func (m *UserManager) Create(u User) bool {
	if res := m.db.Create(&u); res.Error != nil || res.RowsAffected != 1 {
		return false
	}
	return true
}

func (m *UserManager) CreateMany(us []User) bool {
	if res := m.db.Create(&us); res.Error != nil || res.RowsAffected != 1 {
		return false
	}
	return true
}

func (m *UserManager) GetById(id uint) (User, error) {
	var user User
	m.db.First(&user, id)
	return user, nil
}

func (m *UserManager) GetMany() ([]User, error) {
	var users []User
	if result := m.db.Find(&users); result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (m *UserManager) Update(id uint, u interface{}) error {
	var user User
	if res := m.db.Model(&user).Updates(u); res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *UserManager) Delete(id uint) error {
	if res := m.db.Delete(&User{}, 10); res.Error != nil {
		return res.Error
	}
	return nil
}
