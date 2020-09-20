package snapshot

import (
	"fmt"
	"time"
)

var last *SnapShot

// Service snapshot service
type Service struct {
	path string
	C    chan WatchResult
}

// WatchResult is the type of the watch channel
type WatchResult struct {
	Files []File
	Err   error
}

// New returns a pointer to a new service
func New(path string) *Service {
	return &Service{
		path: path,
		C:    make(chan WatchResult),
	}
}

// Latest returns the latest SnapShot available or a new one
func (s *Service) Latest() (*SnapShot, error) {
	if last == nil {
		snap, err := Take(s.path)
		if err != nil {
			return last, err
		}

		last = snap
	}

	return last, nil
}

// ChangedFiles returns all the files that were created, deleted or modified
// It also stores the latest snapsot
func (s *Service) ChangedFiles() ([]File, error) {
	snap, err := Take(s.path)
	if err != nil {
		return make([]File, 0), err
	}

	latest, err := s.Latest()
	if err != nil {
		return make([]File, 0), err
	}

	files, err := snap.Compare(latest)
	if err != nil {
		return make([]File, 0), err
	}

	last = snap

	return files, nil
}

// Watch watches a folder in a new go routine
func (s *Service) Watch(interval time.Duration) error {
	_, err := s.Latest()
	if err != nil {
		return err
	}

	// resChan := make(chan WatchResult)

	// for {
	go func() {
		for {
			files, err := s.ChangedFiles()
			// if err != nil {
			// 	fmt.Println(err)
			// }
			if err != nil {
				s.C <- WatchResult{Err: err}
			}

			if len(files) > 0 {
				fmt.Println("Sending on channel")
				s.C <- WatchResult{Files: files}
			}

			fmt.Println("FUNNING", files, err)
			time.Sleep(interval * time.Second)
		}

		// fmt.Println("FUNNING", files, err)
	}()

	return nil
}
