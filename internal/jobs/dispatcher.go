package jobs

import (
	"errors"
	"fmt"
	"github.com/thomascpowell/drive/internal/models"
	"github.com/thomascpowell/drive/internal/store"
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
	case models.Upload: d.handleUpload(job)
	case models.GetUserFiles: d.handleGetUserFiles(job)
	case models.GetFile: d.handleGetFile(job)
	case models.DeleteFile: d.handleDeleteFile(job)
	default:
		job.Done <- errors.New("unknown job type")
	}
}

func (d *Dispatcher) handleUpload(job *models.Job) {
	file, ok := job.Payload.(*models.File)
	if !ok {
		job.Done <- fmt.Errorf("invalid payload type for upload")
		return
	}
	err := d.Store.CreateFile(file)
	job.Done <- err
}

func (d *Dispatcher) handleGetUserFiles(job *models.Job) {
 // todo
}

func (d *Dispatcher) handleGetFile(job *models.Job) {
 // todo
}

func (d *Dispatcher) handleDeleteFile(job *models.Job) {
 // todo
}
