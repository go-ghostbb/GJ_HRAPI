package model

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}
