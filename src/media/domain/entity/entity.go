package entity

import "time"

type ModelMedia struct {
	GUID      string     `db:"guid"`
	Name      string     `db:"name"`
	Type      string     `db:"type"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type StructQuery struct {
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Sort      string `json:"sort"`
	Order     string `json:"order"`
	Keys      string `json:"keys"`
	Keyword   string `json:"keyword"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}
