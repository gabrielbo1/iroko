package permission

import (
	"github.com/gabrielbo1/iroko/domain"
	"github.com/gabrielbo1/iroko/pkg"
)

//Profile - Access profile of technology company operators.
type Profile struct {
	ID          int          `json:"id"`
	Version     int          `json:"version"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

//ValidProfile - Valid entity Profile.
func ValidProfile(profile *Profile) (*Profile, *domain.Err) {
	if profile.Name == "" {
		return nil, domain.NewErr().WithCode("profile10")
	}
	if len(profile.Permissions) == 0 {
		return nil, domain.NewErr().WithCode("profile20")
	}
	for i := range profile.Permissions {
		if _, err := ValidPermission(&profile.Permissions[i]); err != nil {
			return nil, err
		}
	}
	return profile, nil
}

//ProfileRepository - Define all operations entity Profile.
type ProfileRepository interface {

	//Save - Save new Profile.
	Save(profile Profile) (int, *domain.Err)

	//Update - Update Profile.
	Update(profile Profile) *domain.Err

	//FindById - Find Profile with id.
	FindById(id int) (Profile, *domain.Err)

	//FindByName - Profile page search with name.
	FindByName(name string) (pkg.Page, *domain.Err)
}
