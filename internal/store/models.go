package store

type UserItem struct {
	Email       string `db:"email"`
	AccessToken string `db:"access_token"`
}
