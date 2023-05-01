package domain

type Repository interface {
	GetAll() ([]*User, error)
	GetByID(userID string) (*User, error)
	Save(user *User) error
	Replace(user *User) error
	Delete(userID string) error
}
