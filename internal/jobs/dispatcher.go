package jobs

import (
	"errors"
	"fmt"
	"github.com/thomascpowell/drive/internal/models"
	"github.com/thomascpowell/drive/internal/store"
	"reflect"
)

type Dispatcher struct {
	JobQueue chan *models.Job
	Store    *store.Store
}

func NewDispatcher(store *store.Store, size int) *Dispatcher {
	return &Dispatcher{
		JobQueue: make(chan *models.Job, size),
		Store:    store,
	}
}

func (d *Dispatcher) Stop() {
	close(d.JobQueue)
}

func (d *Dispatcher) Dispatch(job *models.Job) error {
	select {
	case d.JobQueue <- job:
		return nil
	default:
		return errors.New("job queue is full")
	}
}

func (d *Dispatcher) StartWorkers(number int) {
	for i := range number {
		go d.startWorker(i)
	}
}

func (d *Dispatcher) startWorker(id int) {
	for job := range d.JobQueue {
		if job == nil {
			break
		}
		fmt.Printf("worker %d: processing job %s\n", id, job.ID)
		d.process(job)
	}
}

func (d *Dispatcher) process(job *models.Job) {
	switch job.Type {
	case models.Upload:
		d.handleUpload(job)
	case models.GetUserFiles:
		d.handleGetUserFiles(job)
	case models.GetFile:
		d.handleGetFile(job)
	case models.DeleteFile:
		d.handleDeleteFile(job)
	default:
		job.Done <- models.Result{Err: errors.New("unknown job type")}
	}
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
	files, err := d.Store.GetFilesByUserID(userID)
	job.Done <- models.Result{Value: files, Err: err}
}

func (d *Dispatcher) handleGetFile(job *models.Job) {
	fileID, err := validate[uint](job.Payload)
	if err != nil {
		job.Done <- models.Result{Err: err}
		return
	}
	file, err := d.Store.GetFileByID(fileID)
	job.Done <- models.Result{Value: file, Err: err}
}

func (d *Dispatcher) handleDeleteFile(job *models.Job) {
	fileID, err := validate[uint](job.Payload)
	if err != nil {
		job.Done <- models.Result{Err: err}
		return
	}
	err = d.Store.DeleteFileByID(fileID)
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
