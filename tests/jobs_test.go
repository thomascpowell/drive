package tests

import(
	"github.com/thomascpowell/drive/models"
	"github.com/thomascpowell/drive/jobs"
	"time"
	"errors"
	"testing"
)
func TestHandleGetFile_Success(t *testing.T) {
	mock := &MockStore{
		GetFileByIDFunc: func(id uint) (*models.File, error) {
			return &models.File{ID: id, Filename: "example.txt"}, nil
		},
	}
	dispatcher := jobs.NewDispatcher(mock, 1)
	dispatcher.StartWorkers(1)
	defer dispatcher.Stop()
	testFileID := uint(123)
	done := make(chan models.Result, 1)
	job := &models.Job{
		ID:      "JOB_TEST_1",
		Type:    models.GetFile,
		Payload: &testFileID,
		Done:    done,
	}
	if err := dispatcher.Dispatch(job); err != nil {
		t.Fatalf("dispatch failed: %v", err)
	}
	select {
	case res := <-done:
		if res.Err != nil {
			t.Fatalf("unexpected error: %v", res.Err)
		}
		file, ok := res.Value.(*models.File)
		if !ok {
			t.Fatalf("expected *models.File, got %T", res.Value)
		}
		Expect(t, file.ID, 123, "fileID")
	case <-time.After(time.Second):
		t.Fatal("job did not complete")
	}
}

func TestHandleGetFile_Fail(t *testing.T) {
	mock := &MockStore{
		GetFileByIDFunc: func(id uint) (*models.File, error) {
			return nil, errors.New("some error that would have been from db")
		},
	}
	dispatcher := jobs.NewDispatcher(mock, 1)
	dispatcher.StartWorkers(1)
	defer dispatcher.Stop()
	testFileID := uint(123)
	done := make(chan models.Result, 1)
	job := &models.Job{
		ID:      "JOB_TEST_2",
		Type:    models.GetFile,
		Payload: &testFileID,
		Done:    done,
	}
	if err := dispatcher.Dispatch(job); err != nil {
		t.Fatalf("dispatch failed: %v", err)
	}
	select {
	case res := <-done:
		file, _ := res.Value.(*models.File)
		Expect(t, file, nil, "file")
		Expect(t, res.Err != nil, true, "expected error but got nil")		
	case <-time.After(time.Second):
		t.Fatal("job did not complete")
	}
}

