package jobs

import (
	"errors"
	"fmt"
	"github.com/thomascpowell/drive/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/thomascpowell/drive/auth"
	"reflect"
)

func (d *Dispatcher) handleAuthenticateUser(job *models.Job) {
	credentials, err := validate[models.Credentials](job.Payload)
	if err != nil {
		job.Done <- models.Err(err)
		return
	}
	user, err := d.Store.GetUserByUsername(credentials.Username)
	if err != nil {
		job.Done <- models.Err(err)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
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

func (d *Dispatcher) handleRegisterUser(job *models.Job) {
	user, err := validate[models.User](job.Payload)
	if err != nil {
		job.Done <- models.Err(err)
		return
	}
	err = d.Store.CreateUser(user)
	job.Done <- models.Result{Value: user, Err: err}
}

func (d *Dispatcher) handleGetUser(job *models.Job) {
	username, err := validate[string](job.Payload)
	if err != nil {
		job.Done <- models.Err(err)
		return
	}
	user, err := d.Store.GetUserByUsername(*username)
	job.Done <- models.Result{Value: user, Err: err}
}

func (d *Dispatcher) handleUpload(job *models.Job) {
	file, err := validate[models.File](job.Payload)
	if err != nil {
		job.Done <- models.Err(err)
		return
	}
	err = d.Store.CreateFile(file)
	job.Done <- models.Result{Err: err}
}

func (d *Dispatcher) handleGetUserFiles(job *models.Job) {
	userID, err := validate[uint](job.Payload)
	if err != nil {
		job.Done <- models.Result{Err: err}
		return
	}
	files, err := d.Store.GetFilesByUserID(*userID)
	job.Done <- models.Result{Value: files, Err: err}
}

func (d *Dispatcher) handleGetFile(job *models.Job) {
	fileID, err := validate[uint](job.Payload)
	if err != nil {
		job.Done <- models.Result{Err: err}
		return
	}
	file, err := d.Store.GetFileByID(*fileID)
	job.Done <- models.Result{Value: file, Err: err}
}

func (d *Dispatcher) handleDeleteFile(job *models.Job) {
	fileID, err := validate[uint](job.Payload)
	if err != nil {
		job.Done <- models.Result{Err: err}
		return
	}
	err = d.Store.DeleteFileByID(*fileID)
	job.Done <- models.Result{Err: err}
}

func validate[T any](payload any) (*T, error) {
	if payload == nil {
		return nil, errors.New("payload is nil")
	}
	isPointer := reflect.ValueOf(payload).Kind() == reflect.Ptr
	if !isPointer {
		return nil, errors.New("payload must be a pointer")
	}
	ptr, ok := payload.(*T)
	if !ok {
		return nil, fmt.Errorf("unexpected type in payload: got %T, want *%T", payload, new(T))
	}
	return ptr, nil
}
