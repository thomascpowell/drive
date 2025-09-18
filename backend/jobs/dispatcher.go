package jobs

import (
	"errors"
	"fmt"
	"github.com/thomascpowell/drive/models"
	"github.com/thomascpowell/drive/store"
)

type Dispatcher struct {
	JobQueue chan *models.Job
	Store    store.StoreInterface
	RC       store.RedisInterface
}

func NewDispatcher(store store.StoreInterface, rc store.RedisInterface, size int) *Dispatcher {
	dispatcher := Dispatcher{
		JobQueue: make(chan *models.Job, size),
		Store:    store,
		RC: 			rc,
	}
	return &dispatcher
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
		close(job.Done)
	}
}

func (d *Dispatcher) process(job *models.Job) {
	switch job.Type {
	case models.Upload:
		payload := &job.Payload.Upload
		d.handleUpload(payload, job)
	case models.GetUserFiles:
		payload := &job.Payload.GetUserFiles
		d.handleGetUserFiles(payload, job)
	case models.GetFile:
		payload := &job.Payload.GetFile
		d.handleGetFile(payload, job)
	case models.DeleteFile:
		payload := &job.Payload.DeleteFile
		d.handleDeleteFile(payload, job)
	case models.GetUser:
		payload := &job.Payload.GetUser
		d.handleGetUser(payload, job)
	case models.RegisterUser:
		payload := &job.Payload.RegisterUser
		d.handleRegisterUser(payload, job)
	case models.AuthenticateUser:
		payload := &job.Payload.AuthenticateUser
		d.handleAuthenticateUser(payload, job)
	default:
		job.Done <- models.Result{Err: errors.New("unknown job type")}
	}
}
