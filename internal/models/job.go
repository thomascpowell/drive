package models

type Job struct {
	ID      string
	Type    JobType
	Payload any
	Done    chan error
}

type JobType int
const (
	Upload JobType = iota
	GetUserFiles
	GetFile
	DeleteFile
)
