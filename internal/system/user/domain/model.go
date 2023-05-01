package domain

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}
