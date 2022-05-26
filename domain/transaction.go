package domain

//SGBD - Define type to SGDB consts.
type SGBD int

const (
	PostgreSQL SGBD = iota
	MySQL
	Oracle
	SQLServer
)

//Transaction - Command write or update data base,
// this transction is a unique command to exect
type Transction struct {
	//Command - SGBD instructions.
	Command string `json:"command"`

	//Base - SGBD identifier.
	Base SGBD `json:"base"`

	//App - App name to origin trasaction.
	App string `json:"app"`

	//CreateMoment - Timestamp moment to create commad origin.
	CreateMoment string `json:"create_moment"`

	//ExecutedMoment - Timestamp moment to executed instruction.
	ExecutedMoment string `json:"executed_moment"`
}

//Block - Block of transactions to be executed.
type Block struct {

	//Uuid - Unique ID identify to block the transactions.
	Uuid string `json:"uuid"`

	//Transactions - Transactions block.
	Transactions []Transction `json:"transactions"`
}
