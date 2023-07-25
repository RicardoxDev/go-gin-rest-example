package entity

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}
