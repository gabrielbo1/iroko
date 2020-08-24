package permission

import "github.com/gabrielbo1/iroko/domain"

//Routine - Defines a system routine that can be defined
//with or without an associated Path (API).
type Routine struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	Path        string `json:"path"`
	RoutineName string `json:"routineName"`
}

//ValidRoutine - Valid Routine entity
func ValidRoutine(routine *Routine) (*Routine, *domain.Err) {
	if routine.Path == "" {
		return nil, domain.NewErr().WithCode("routine10")
	}
	if routine.RoutineName == "" {
		return nil, domain.NewErr().WithCode("routine20")
	}
	return routine, nil
}

//RoutineRepository - Define all operations with Routine entity.
type RoutineRepository interface {

	//Save - Save new Routine entity.
	Save(profile Routine) (int, *domain.Err)

	//Update - Update Routine entity.
	Update(profile Routine) *domain.Err

	//FindById - Find Routine with id.
	FindById(id int) (Routine, *domain.Err)

	//FindAll - Find all Routine entities.
	FindAll() ([]Routine, *domain.Err)
}
