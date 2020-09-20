package snapshot

import (
	"os"
	"path/filepath"
	"time"
)

// SnapShot struct that holds a map of every file in the root folder
type SnapShot struct {
	folder  string
	files   map[string]File
	takenAt time.Time
}

// Take returns a pointer to a Snapshoot of the folder specified as param
func Take(path string) (*SnapShot, error) {
	snap := SnapShot{
		folder:  path,
		files:   make(map[string]File),
		takenAt: time.Now(),
	}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() == false {
			snap.files[path] = File{
				path: path,
				info: info,
				// method: Created,
			}
		}

		return nil
	})

	return &snap, err
}

// Compare returns a slice of all the modified files, new files, or deleted files
func (s *SnapShot) Compare(old *SnapShot) ([]File, error) {
	var files []File

	for path, f := range s.files {

		// File was just created
		if _, ok := old.files[path]; ok == false {
			f.method = Created
			files = append(files, f)
			continue
		}

		// File was modified
		if old.files[path].info.ModTime().Before(f.info.ModTime()) {
			f.method = Modified
			files = append(files, f)
		}
	}

	for path := range old.files {

		// File was deleted
		if f, ok := s.files[path]; ok == false {
			f.method = Deleted
			files = append(files, f)
			continue
		}
	}

	return files, nil
}

// Files returns all the files in the snapshot
func (s *SnapShot) Files() []File {
	var files []File
	for _, f := range s.files {
		files = append(files, f)
	}

	return files
}
