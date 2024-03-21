package models

type User struct {
    ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
    Name     string `json:"name"`
    Username string `json:"username"`
    Password string `json:"password"`
    RoleID   int    `json:"role_id" form:"role_id" gorm:"not null"`
    Role     Role   `json:"role" gorm:"foreignKey:RoleID"`
}


