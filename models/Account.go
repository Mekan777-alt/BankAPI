package models

type Account struct {
	ID         int    `json:"ID" gorm:"primaryKey"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Bill       []Bill `json:"bills"`
}
