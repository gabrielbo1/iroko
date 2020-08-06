package domain

//LoginCompany - Entity for access control of persons associated
//with client companies that may eventually access external
//functionalities by system or integrations by api.
type LoginCompany struct {
	ID       int     `json:"id"`
	Login    string  `json:"login"`
	Password string  `json:"password"`
	Company  Company `json:"company"`
}
