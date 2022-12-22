package repositories

import (
	"dewetour/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAcc() ([]models.User, error)
	GetAcc(ID int) (models.User, error)
	MakeAcc(user models.User) (models.User, error)
	EditAcc(user models.User, ID int) (models.User, error)
	DeleteAcc(user models.User, ID int) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAcc() ([]models.User, error) {
	var users []models.User
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

	return users, err
}

func (r *repository) GetAcc(ID int) (models.User, error) {
	var user models.User
	err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}

func (r *repository) MakeAcc(user models.User) (models.User, error) {
	err := r.db.Exec("INSERT INTO users(name, email, password, phone, address, gender, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?)", user.Name, user.Email, user.Password, user.Phone, user.Address, user.Gender, time.Now(), time.Now()).Error

	return user, err
}

func (r *repository) EditAcc(user models.User, ID int) (models.User, error) {
	err := r.db.Raw("UPDATE users SET name=?, email=?, password=? WHERE id=?", user.Name, user.Email, user.Password, ID).Scan(&user).Error

	return user, err
}

func (r *repository) DeleteAcc(user models.User, ID int) (models.User, error) {
	err := r.db.Raw("DELETE FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}
