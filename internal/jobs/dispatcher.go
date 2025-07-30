package jobs

import (
	"errors"
	"github.com/thomascpowell/drive/internal/models"
	"github.com/thomascpowell/drive/internal/store"
)

type Dispatcher struct {
	JobQueue chan models.Job
	Store    *store.Store
}

func NewDispatcher(store *store.Store, size int) *Dispatcher {
	return &Dispatcher{
		JobQueue: make(chan models.Job, size),
		Store:    store,
	}
}

func (d *Dispatcher) Dispatch(job models.Job) error {
	select {
	case d.JobQueue <- job:
		return nil
	default:
		return errors.New("job queue is full")
	}
}

func (d *Dispatcher) StartWorker() {
	go func() {
		for job := range d.JobQueue {
			d.process(&job)
		}
	}()
}

func (d *Dispatcher) process(job *models.Job) {
	switch job.Type {
	case models.Upload: handleUpload(job)
	case models.GetUserFiles: handleGetUserFiles(job)
	case models.GetFile: handleGetFile(job)
	case models.DeleteFile: handleDeleteFile(job)
	default:
		job.Done <- errors.New("unknown job type")
	}
}
