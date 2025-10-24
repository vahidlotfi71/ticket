package Models

type Users struct {
	Model
	Name     string `json:"name" gorm:"size:255"`
	Email    string `json:"email" gorm:"size:400 ,unique"`
	Password string `json:"password" gorm:"size:255"`
	Image    string `json:"image" gorm:"size:255"`
	Phone    string `json:"phone" gorm:"size:11"`
}
