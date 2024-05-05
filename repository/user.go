package repository

import (
	"database/sql"
	"log"

	"github.com/Adekabang/social-cat/model"
	"github.com/Adekabang/social-cat/utils"
	"github.com/google/uuid"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{Db: db}
}

// DeleteUser implements UserRepositoryInterface
func (m *UserRepository) DeleteUser(id string) bool {
	_, err := m.Db.Exec("DELETE FROM user WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// GetAllUsers implements UserRepositoryInterface
func (m *UserRepository) GetAllUsers() []model.GetUser {
	query, err := m.Db.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer query.Close()
	var users []model.GetUser
	if query != nil {
		for query.Next() {
			var (
				id            string
				created_at    string
				email         string
				name          string
				password_hash string
			)
			err := query.Scan(&id, &created_at, &email, &name, &password_hash)
			if err != nil {
				log.Println(err)
			}
			user := model.GetUser{Id: id, Email: email, Name: name}
			users = append(users, user)
		}
	}
	return users
}

// GetOneUser implements UserRepositoryInterface
func (m *UserRepository) GetOneUser(id string) model.GetUser {
	query, err := m.Db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return model.GetUser{}
	}
	defer query.Close()
	var user model.GetUser
	if query != nil {
		for query.Next() {
			var (
				id            string
				created_at    string
				email         string
				name          string
				password_hash string
			)
			err := query.Scan(&id, &created_at, &email, &name, &password_hash)
			if err != nil {
				log.Println(err)
			}
			user = model.GetUser{Id: id, Email: email, Name: name}
		}
	}
	return user
}

// InsertUser implements UserRepositoryInterface
func (m *UserRepository) InsertUser(post model.PostUser) bool {

	// jwtSecret, err := utils.GetJWTSecret()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return false
	// }

	// create jwt
	// tokenString, err := utils.GetToken(post.Email, jwtSecret)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	// c.JSON(500, gin.H{"error": "token creation oopsie"})
	// 	return false
	// }

	hashedPassword, err := utils.HashPassword(post.Password)

	if err != nil {
		log.Println(err.Error())
		return false
	}

	uuidUser := uuid.New()

	stmt, err := m.Db.Prepare("INSERT INTO users(id, email, name, password_hash) VALUES ($1,$2,$3,$4)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(uuidUser, post.Email, post.Name, hashedPassword)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// UpdateUser implements UserRepositoryInterface
func (m *UserRepository) UpdateUser(id string, post model.PostUser) model.GetUser {
	_, err := m.Db.Exec("UPDATE users SET email = $1, name = $2 WHERE id = $3", post.Email, post.Name, id)
	if err != nil {
		log.Println(err)
		return model.GetUser{}
	}
	return m.GetOneUser(id)
}
