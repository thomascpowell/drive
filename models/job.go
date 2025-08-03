package models

type Job struct {
	ID      string
	Type    JobType
	Payload any
	Done    chan Result
}

type JobType int

const (
	Upload JobType = iota
	GetUserFiles
	GetFile
	DeleteFile
	GetUser

	// TODO: implement
	RegisterUser // payload: User
	AuthenticateUser // payload: {username, password}
)
