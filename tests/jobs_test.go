	package tests

	import (
		"errors"
		"fmt"
		"github.com/thomascpowell/drive/jobs"
		"github.com/thomascpowell/drive/models"
		"github.com/thomascpowell/drive/utils"
		"testing"
		"time"
	)

	func TestGetUser(t *testing.T) {
		mock := &MockStore{
			GetUserByUsernameFunc: func(username string) (*models.User, error) {
				return &models.User{ID: 1}, nil
			},
		}
		dispatcher := jobs.NewDispatcher(mock, 1)
		dispatcher.StartWorkers(1)
		defer dispatcher.Stop()
		testUsername := "user"
		done := make(chan models.Result, 1)
		job := &models.Job{
			ID:      "JOB_TEST_0",
			Type:    models.GetUser,
			Payload: &testUsername,
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
			user, ok := res.Value.(*models.User)
			if !ok {
				t.Fatalf("expected *models.User, got %T", res.Value)
			}
			utils.Expect(t, user.ID, 1, "userID")
		case <-time.After(time.Second):
			t.Fatal("job did not complete")
		}
	}

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
			utils.Expect(t, file.ID, 123, "fileID")
		case <-time.After(time.Second):
			t.Fatal("job did not complete")
		}
	}

	func TestHandleGetFile_Fail(t *testing.T) {
		mock := &MockStore{
			GetFileByIDFunc: func(id uint) (*models.File, error) {
				return nil, errors.New("some error")
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
			utils.Expect(t, file, nil, "file")
			utils.Expect(t, res.Err != nil, true, "expected error but got nil")
		case <-time.After(time.Second):
			t.Fatal("job did not complete")
		}
	}

	func TestQueue(t *testing.T) {
		JOB_COUNT := 10
		WORKER_COUNT := 2
		TIME_LIMIT := 6 * time.Second

		mock := &MockStore{
			GetFileByIDFunc: func(id uint) (*models.File, error) {
				// simulate workload
				// (so that both workers get a chance)
				// time.Sleep(500 * time.Millisecond)
				return &models.File{ID: id, Filename: "example.txt"}, nil
			},
		}

		dispatcher := jobs.NewDispatcher(mock, JOB_COUNT)
		dispatcher.StartWorkers(WORKER_COUNT)
		results := make([]chan models.Result, JOB_COUNT)
		resultCh := make(chan models.Result) // accumulates all results

		for i := range JOB_COUNT {
			payload := uint(i)
			done := make(chan models.Result, 1)
			job := &models.Job{
				ID:      fmt.Sprintf("JOB_TEST_%d", i),
				Type:    models.GetFile,
				Payload: &payload,
				Done:    done,
			}
			if err := dispatcher.Dispatch(job); err != nil {
				t.Fatalf("dispatch failed: %v", err)
			}
			results[i] = done
		}

		for _, done := range results {
			go func(c chan models.Result) {
				select {
				case res := <-c:
					resultCh <- res
				case <-time.After(5 * time.Second):
					resultCh <- models.Result{Err: fmt.Errorf("job timeout")}
				}
			}(done)
	}

	received := 0
	for received < JOB_COUNT {
		select {
		case res := <-resultCh:
			utils.Expect(t, res.Err, nil, "job failed")
			file, ok := res.Value.(*models.File)
			if !ok {
				t.Fatalf("expected *models.File, got %T", res.Value)
			}
			fmt.Printf("got file: %d\n", file.ID)
			received++
		case <-time.After(TIME_LIMIT):
			t.Fatal("exceeded time limit")
		}
	}
}
