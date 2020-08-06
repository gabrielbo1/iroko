package permission

//Routine - Defines a system routine that can be defined
//with or without an associated Path (API).
type Routine struct {
	ID          int    `json:"id"`
	Path        string `json:"path"`
	RoutineName string `json:"routineName"`
}
