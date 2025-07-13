package Models

type User struct {
	Model
	Name     string `gorm:"size:255"`
	Email    string `gorm:"unique"`
	Password string `gorm:"size:255"`
}
