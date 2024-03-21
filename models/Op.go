package models

import "time"

type Op struct {
    ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
    Information string    `json:"information"`
    Code        string    `json:"code"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"-"`
    UserID int             `json:"user_id" form:"user_id" gorm:"not null"`
    User User              `json:"user"`

}

