package domain

type Admin struct {
	Id        int    `db:"id"`
	Username  string `db:"username"`
	Fullname  string `db:"fullname"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	DeletedAt string `db:"deleted_at"`
	IsDeleted bool   `db:"is_deleted"`
}

type Terminal struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	DeletedAt string `db:"deleted_at"`
	IsDeleted bool   `db:"is_deleted"`
}
