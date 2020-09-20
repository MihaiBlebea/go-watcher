package snapshot

import "os"

// Method represents the last action made on the file
type Method string

// Method constants
const (
	Created  Method = "created"
	Modified Method = "modified"
	Deleted  Method = "deleted"
)

// File represents all the data that we gathered on a file
type File struct {
	path   string
	info   os.FileInfo
	method Method
}

// Path returns the path as string
func (f *File) Path() string {
	return f.path
}

// Info returns the file information
func (f *File) Info() os.FileInfo {
	return f.info
}

// Method returns the file method type
func (f *File) Method() Method {
	return f.method
}
