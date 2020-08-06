package permission

import "github.com/gabrielbo1/iroko/domain"

//ProfileCompany - Access profile of client companies.
type ProfileCompany struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Permissions []Permission   `json:"permissions"`
	Company     domain.Company `json:"company"`
}
