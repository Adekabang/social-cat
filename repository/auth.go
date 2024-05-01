package repository

import (
	"database/sql"
	"log"

	"github.com/Adekabang/social-cat/model"
	"github.com/Adekabang/social-cat/utils"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type AuthRepository struct {
	Db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepositoryInterface {
	return &AuthRepository{Db: db}
}

func (m *AuthRepository) Register(user model.RegisterUser) string {

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		log.Println(err.Error())
		return "failed hashing"
	}

	uuidUser := uuid.New()

	stmt, err := m.Db.Prepare("INSERT INTO users(id, email, name, password_hash) VALUES ($1,$2,$3,$4)")
	if err != nil {
		log.Println(err)
		return "server failed"
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(uuidUser, user.Email, user.Name, hashedPassword)
	if err2 != nil {
		log.Println(err2)
		log.Println(string(err2.(*pq.Error).Code))
		return string(err2.(*pq.Error).Code)
	}
	return "success"
}
