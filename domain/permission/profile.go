package permission

//Profile - Access profile of technology company operators.
type Profile struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}
