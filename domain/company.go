package domain

//Company - Company Entity company with all related information,
//and the other relationships belong to a company and technology
//company in turn have many client companies.
type Company struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
