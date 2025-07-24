package Models

type User struct {
	Model

	Name     string `json:"name"     gorm:"size:255"`
	Email    string `json:"email"    gorm:"unique"`
	Password string `json:"password" gorm:"size:255"`
	Image    string `json:"image" gorm:"size:255"`
	Phone    string `json:"phone" gorm:"size:11"` // محدودیت مین و ماکس باید داشته باشد
}
