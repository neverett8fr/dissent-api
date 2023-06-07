package service

type newUserIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type tokenOut struct {
	Token string `json:"token"`
}
