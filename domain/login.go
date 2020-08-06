package domain

//Login - Entity login with all members of the technology company,
//who in turn are the direct operators of the system.
type Login struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
