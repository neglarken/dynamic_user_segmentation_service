package entity

type SlugsUsers struct {
	UserId int `db:"user_id" json:"user_id"`
	SlugId int `db:"slug_id" json:"slug_id"`
}
