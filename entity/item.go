package entity

import "time"

type Item struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"created_at" json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`

	/*...*/
}
