package model

type UserGenre struct {
	UserId  int `db:"user_id"`
	GenreId int `db:"genre_id"`

	// foreign key
	User  *User  `db:"us"`
	Genre *Genre `db:"gr"`
}
