package jobs

import (
	"fmt"
	"github.com/thomascpowell/drive/models"
	"github.com/thomascpowell/drive/auth"
	"github.com/thomascpowell/drive/utils"
)

func (d *Dispatcher) handleAuthenticateUser(payload *models.AuthenticateUserPayload, job *models.Job) {
	user, err := d.Store.GetUserByUsername(payload.Username)
	if err != nil {
		job.Done <- models.Err(err)
		return
	}
	if !utils.CheckPasswordHash(payload.Password, user.Password) {
		job.Done <- models.Err(fmt.Errorf("invalid credentials"))
		return
	}
	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		job.Done <- models.Err(fmt.Errorf("token generation failed"))
		return
	}
	job.Done <- models.Result{Value: token}
}

func (d *Dispatcher) handleRegisterUser(payload *models.RegisterUserPayload, job *models.Job) {
	user := payload.User
	err := d.Store.CreateUser(user)
	job.Done <- models.Result{Value: user, Err: err}
}

func (d *Dispatcher) handleGetUser(payload *models.GetUserPayload, job *models.Job) {
	user, err := d.Store.GetUserByUsername(payload.Username)
	job.Done <- models.Result{Value: user, Err: err}
}

func (d *Dispatcher) handleUpload(payload *models.UploadPayload, job *models.Job) {
	err := d.Store.CreateFile(payload)
	job.Done <- models.Result{Err: err}
}

func (d *Dispatcher) handleGetUserFiles(payload *models.GetUserFilesPayload, job *models.Job) {
	files, err := d.Store.GetFilesByUserID(payload.UserID)
	job.Done <- models.Result{Value: files, Err: err}
}

func (d *Dispatcher) handleGetFile(payload *models.GetFilePayload, job *models.Job) {
	file, err := d.Store.GetFileByID(payload.FileID)
	job.Done <- models.Result{Value: file, Err: err}
}

func (d *Dispatcher) handleDeleteFile(payload *models.DeleteFilePayload, job *models.Job) {
	err := d.Store.DeleteFileByID(payload.FileID)
	job.Done <- models.Result{Err: err}
}
