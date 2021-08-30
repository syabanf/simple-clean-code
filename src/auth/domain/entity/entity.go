package entity

type ModelUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
