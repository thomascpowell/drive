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
}

func NewDispatcher(store store.StoreInterface, size int) *Dispatcher {
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
	case models.GetUser:
		d.handleGetUser(job)
	case models.RegisterUser:
		d.handleRegisterUser(job)
	case models.AuthenticateUser:
		d.handleAuthenticateUser(job)
	default:
		job.Done <- models.Result{Err: errors.New("unknown job type")}
	}
}

