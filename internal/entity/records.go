package entity

import "time"

type Records struct {
	Id        int       `db:"id"`
	UserId    int       `db:"user_id" json:"user_id"`
	SlugTitle string    `db:"slug_title" json:"slug_title"`
	Operation string    `db:"operation" json:"operation"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
