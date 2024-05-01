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

func (m *AuthRepository) Register(user model.RegisterUser) model.ResponseMessage {
	var response model.ResponseMessage
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		log.Println(err.Error())
		response = model.ResponseMessage{Status: "failed", Msg: "failed hashing"}
		return response
	}

	uuidUser := uuid.New()

	stmt, err := m.Db.Prepare("INSERT INTO users(id, email, name, password_hash) VALUES ($1,$2,$3,$4)")
	if err != nil {
		log.Println(err)
		response = model.ResponseMessage{Status: "failed", Msg: "server failed"}
		return response
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(uuidUser, user.Email, user.Name, hashedPassword)
	if err2 != nil {
		log.Println(err2)
		log.Println(string(err2.(*pq.Error).Code))
		response = model.ResponseMessage{Status: "failed", Msg: string(err2.(*pq.Error).Code)}
		return response
	}

	token, err := utils.GenerateToken(uuidUser.String())

	if err != nil {
		log.Println(err)
		response = model.ResponseMessage{Status: "failed", Msg: "error"}
	}
	response = model.ResponseMessage{Status: "success", Msg: token, Data: model.UserData{Email: user.Email, Name: user.Name}}
	return response
}

func (m *AuthRepository) Login(input model.LoginUser) model.ResponseMessage {
	var response model.ResponseMessage
	query, err := m.Db.Query("SELECT * FROM users WHERE email = $1", input.Email)
	if err != nil {
		log.Println(err)
		response = model.ResponseMessage{Status: "failed", Msg: "user not found"}
	}
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
			err2 := utils.VerifyPassword(input.Password, password_hash)
			if err2 != nil {
				log.Println(err2)
				response = model.ResponseMessage{Status: "failed", Msg: "wrong password"}
			} else {
				token, err := utils.GenerateToken(id)
				if err != nil {
					log.Println(err)
					response = model.ResponseMessage{Status: "failed", Msg: "error"}
				}
				response = model.ResponseMessage{Status: "success", Msg: token, Data: model.UserData{Email: email, Name: name}}
			}

		}
	}

	return response
}
