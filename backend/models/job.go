package models

type Job struct {
	ID      string
	Type    JobType
	Payload *Payload
	Done    chan Result
}

type JobType int

const (
	Upload JobType = iota
	GetUserFiles
	GetFile
	DeleteFile
	GetUser
	RegisterUser
	AuthenticateUser
	GetShareLink
)

type Payload struct {
	Upload           UploadPayload
	GetUserFiles     GetUserFilesPayload
	GetFile          GetFilePayload
	DeleteFile       DeleteFilePayload
	GetUser          GetUserPayload
	RegisterUser     RegisterUserPayload
	AuthenticateUser AuthenticateUserPayload
	GetShareLink     GetShareLinkPayload
}

type UploadPayload = File

type GetShareLinkPayload struct {
	FileID uint
	TTL    uint
}

type GetUserFilesPayload struct {
	UserID uint
}

type GetFilePayload struct {
	FileID uint
}

type DeleteFilePayload struct {
	UserID uint
	FileID uint
}

type GetUserPayload struct {
	Username string
}

type RegisterUserPayload struct {
	User *User
}

type AuthenticateUserPayload = Credentials

func NewGetShareLinkPayload(fileID uint, ttl uint) *Payload {
	return &Payload{
		GetShareLink: GetShareLinkPayload{
			FileID: fileID,
			TTL:    ttl,
		},
	}
}

func NewUploadPayload(file File) *Payload {
	return &Payload{
		Upload: UploadPayload(file),
	}
}

func NewGetUserFilesPayload(userID uint) *Payload {
	return &Payload{
		GetUserFiles: GetUserFilesPayload{
			UserID: userID,
		},
	}
}

func NewGetFilePayload(fileID uint) *Payload {
	return &Payload{
		GetFile: GetFilePayload{
			FileID: fileID,
		},
	}
}

func NewDeleteFilePayload(userID, fileID uint) *Payload {
	return &Payload{
		DeleteFile: DeleteFilePayload{
			UserID: userID,
			FileID: fileID,
		},
	}
}

func NewGetUserPayload(username string) *Payload {
	return &Payload{
		GetUser: GetUserPayload{
			Username: username,
		},
	}
}

func NewRegisterUserPayload(user User) *Payload {
	return &Payload{
		RegisterUser: RegisterUserPayload{
			User: &user,
		},
	}
}

func NewAuthenticateUserPayload(username, password string) *Payload {
	return &Payload{
		AuthenticateUser: AuthenticateUserPayload{
			Username: username,
			Password: password,
		},
	}
}
