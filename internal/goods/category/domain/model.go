package domain

type Category struct {
	ID   string `json:"id" gorm:"primaryKey;<-:create"`
	Name string `json:"name" gorm:"not null"`
	Note string `json:"note"`
}
