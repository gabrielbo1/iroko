package domain

import "github.com/gabrielbo1/iroko/pkg"

// System - Define basic system information.
type System struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Rotines []Rotine `json:"rotines"`
}

// NewSystem - Create valid new system entity.
func NewSystem(system System) (*System, *pkg.Err) {
	if !pkg.NameIsValid(system.Name) {
		return nil, pkg.NewErr().WithCode("PERMISSION_SYSTEM_10")
	}
	if len(system.Rotines) > 0 {
		for _, r := range system.Rotines {
			if _, err := NewRotine(r); err != nil {
				return nil, err
			}
			if system.Id != r.SystemId {
				return nil, pkg.NewErr().WithCode("PERMISSION_SYSTEM_20")
			}
		}
	}
	return &System{
		Name:    system.Name,
		Rotines: system.Rotines,
	}, nil
}

// Permission - Define permission entity,
// one to many relation between user and permission.
type Permission struct {
	Id       string `json:"id"`
	RotineId string `json:"rotine_id"`
	UserId   string `json:"user_id"`
	Create   bool   `json:"create"`
	Read     bool   `json:"read"`
	Update   bool   `json:"update"`
	Delete   bool   `json:"delete"`
}

//NewPermission - Create new permission valid.
func NewPermission(permission *Permission) (*Permission, *pkg.Err) {
	if !pkg.UuidIsValid(permission.RotineId) {
		return nil, pkg.NewErr().WithCode("PERMISSION_PERMISSION_10")
	}
	if !pkg.UuidIsValid(permission.UserId) {
		return nil, pkg.NewErr().WithCode("PERMISSION_PERMISSION_20")
	}
	return &Permission{
		RotineId: permission.RotineId,
		UserId:   permission.UserId,
		Create:   permission.Create,
		Read:     permission.Read,
		Update:   permission.Update,
		Delete:   permission.Delete,
	}, nil
}

// Rotine - Define rotine enitty,
// path is optinal.
type Rotine struct {
	Id       string `json:"id"`
	SystemId string `json:"system_id"`
	Rotine   string `json:"rotiine"`
	Path     string `json:"path"`
}

// NewRotine - Create a new valid rotine.
func NewRotine(rotine Rotine) (*Rotine, *pkg.Err) {
	if !pkg.UuidIsValid(rotine.SystemId) {
		return nil, pkg.NewErr().WithCode("PERMISSION_ROTINE_10")
	}
	if !pkg.NameIsValid(rotine.Rotine) {
		return nil, pkg.NewErr().WithCode("PERMISSION_ROTINE_20")
	}
	if rotine.Path != "" && !pkg.NameIsValid(rotine.Path) {
		return nil, pkg.NewErr().WithCode("PERMISSION_ROTINE_30")
	}
	return &Rotine{
		SystemId: rotine.SystemId,
		Rotine:   rotine.Rotine,
		Path:     rotine.Path,
	}, nil
}
