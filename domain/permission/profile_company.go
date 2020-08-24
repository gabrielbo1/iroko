package permission

import (
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/pkg"
)

//ProfileCompany - Access profile of client companies.
type ProfileCompany struct {
	ID          int            `json:"id"`
	Version     int            `json:"version"`
	Name        string         `json:"name"`
	Permissions []Permission   `json:"permissions"`
	Company     domain.Company `json:"company"`
}

//ValidProfileCompany - Valid ProfileCompany entity.
func ValidProfileCompany(profileCompany *ProfileCompany) (*ProfileCompany, *domain.Err) {
	if profileCompany.Name == "" {
		return nil, domain.NewErr().WithCode("profile10")
	}
	if len(profileCompany.Permissions) == 0 {
		return nil, domain.NewErr().WithCode("profile20")
	}
	for i := range profileCompany.Permissions {
		if _, err := ValidPermission(&profileCompany.Permissions[i]); err != nil {
			return nil, err
		}
	}
	if profileCompany.Company.ID == 0 {
		return nil, domain.NewErr().WithCode("profilecompany10")
	}
	return profileCompany, nil
}

type ProfileCompanyRepository interface {
	//Save - Save new ProfileCompany.
	Save(profile ProfileCompany) (int, *domain.Err)

	//Update - Update ProfileCompany.
	Update(profile ProfileCompany) *domain.Err

	//FindById - Find ProfileCompany with id.
	FindById(id int) (ProfileCompany, *domain.Err)

	//FindByName - Profile page search with name.
	FindByName(name string) (pkg.Page, *domain.Err)
}
