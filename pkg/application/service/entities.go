package service

type newUserIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type newEventIn struct {
	Organiser   string `json:"organiser"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"date"`
}

type tokenOut struct {
	Token string `json:"token"`
}
