package Models

type User struct {
	Model

	Name     string `json:"name"     gorm:"size:255"`
	Email    string `json:"email"    gorm:"unique"`
	Password string `json:"password" gorm:"size:255"`
}
