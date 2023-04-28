package model

type User struct {
	ID       string `gorm:"primaryKey;<-:create;size:191"`
	Name     string
	Phone    string
	Email    string
	Gender   string
	Password string
}
