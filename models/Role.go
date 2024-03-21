package models

type Role struct {
    ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
    Role        string `json:"role"`
    Permessions string `json:"permessions" gorm:"type:varchar(255)"`
    // User User
}
