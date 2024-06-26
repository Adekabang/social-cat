package model

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PostUser struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type GetUser struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserUri struct {
	ID string `uri:"id" binding:"required"`
}

type RegisterUser struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserData struct {
	Email string
	Name  string
}
type ResponseMessage struct {
	Status string   `json:"status"`
	Data   UserData `json:"data"`
	Msg    string   `json:"msg"`
}
