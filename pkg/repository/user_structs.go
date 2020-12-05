package repository

type User struct {
	ID         int64  `db:"user_id" json:"id"`
	Email      string `db:"email" json:"email"`
	Pass       string `db:"pass" json:"pass,omitempty"`
	Version    int64  `db:"version" json:"version,omitempty"`
	FirstName  string `db:"first_name" json:"first_name"`
	LastName   string `db:"last_name" json:"last_name"`
	RegisterAt string `db:"register_date" json:"register_date"`
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetEmail() string {
	return u.Email
}
