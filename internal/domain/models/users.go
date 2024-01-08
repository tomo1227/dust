package models

const TableNameUser = "users"

// User ドメインモデル
type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Age  uint64 `gorm:"not null" json:"age"`
	Name string `gorm:"not null;default:unknown" json:"name"`
}

// TableName Payment's table name
func (*User) TableName() string {
	return TableNameUser
}

type Users []User
