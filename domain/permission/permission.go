package permission

import "github.com/gabrielbo1/iroko/domain"

//Permission - Defines a specific permission associated
//with a certain routine, the set of these permissions
//formed a profile associated with one or more users
//that can be people or other systems.
type Permission struct {
	ID      int     `json:"id"`
	Version int     `json:"version"`
	Routine Routine `json:"routine"`
	Create  bool    `json:"create"`
	Read    bool    `json:"read"`
	Update  bool    `json:"update"`
	Delete  bool    `json:"delete"`
}

//ValidPermission - Valid Permission entity.
func ValidPermission(permission *Permission) (*Permission, *domain.Err) {
	if !(permission.Create ||
		permission.Read ||
		permission.Update ||
		permission.Delete) {
		return nil, domain.NewErr().WithCode("permission10")
	}
	return permission, nil
}
