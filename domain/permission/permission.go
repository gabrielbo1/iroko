package permission

//Permission - Defines a specific permission associated
//with a certain routine, the set of these permissions
//formed a profile associated with one or more users
//that can be people or other systems.
type Permission struct {
	ID      int     `json:"id"`
	Routine Routine `json:"routine"`
	Create  bool    `json:"create"`
	Read    bool    `json:"read"`
	Update  bool    `json:"update"`
	Delete  bool    `json:"delete"`
}
