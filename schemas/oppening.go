package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Oppening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type OppeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createAt"`
	UpdateAt  time.Time `json:"updateAt"`
	DeletedAt time.Time `json:"deleteAt,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}
