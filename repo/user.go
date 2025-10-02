package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"Id"db:"id"`
	FirstName   string `json:"first_name"db:"first_name"`
	LastName    string `json:"last_name"db:"last_name"`
	Email       string `json:"email"db:"email"`
	Password    string `json:"password"db:"password"`
	IsShopOwner bool   `json:"is_shop_owner"db:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, password string) (*User, error)
}

type userRepo struct {
	// users []User
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r userRepo) Create(user User) (*User, error) {
	// 	if user.ID != 0 {
	// 	return &user,nil
	// }
	// user.ID = len(r.users) + 1
	// r.users = append(r.users, user)
	// return &user,nil

	query := `
		INSERT INTO users (first_name, last_name, email, password, is_shop_owner)
		VALUES (:first_name, :last_name, :email, :password, :is_shop_owner)
		RETURNING id
	`
	var userID int
	rows,err:=  r.db.NamedQuery(query,user )
	if err != nil{
		return nil, err
	}

	if rows.Next(){
		rows.Scan(&userID)
	}
	user.ID = userID
	return &user,err
}

func (r userRepo) Find(email, password string) (*User, error) {
	// for _, u := range r.users {
	// 	if u.Email == email && u.Password == password {
	// 		return &u, nil
	// 	}
	// }
	// return nil, nil

	var user User
	query := `
		SELECT id, first_name, last_name, email, password, is_shop_owner, created_at, updated_at
		FROM users
		WHERE email = $1
		LIMIT 1
	`
	err := r.db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil,nil
		}
		// return nil, err if not found or another DB error
		return nil, err
	}
	return &user, nil

}
