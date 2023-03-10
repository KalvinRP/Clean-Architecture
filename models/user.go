package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	Phone     string    `json:"phone" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: enum('Laki-laki', 'Perempuan', 'Lainnya')"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
