package output

import (
	"io/ioutil"
)

//CreateFile creates a File instance.
func CreateFile(path string) *File {
	return &File{
		path: path,
	}
}

//File is a simple struct to wrap the file interaction methods from the standard library.
type File struct {
	path string
}

//Write writes a list of bytes on the file system.
func (f *File) Write(content []byte) (int, error) {
	return len(content), ioutil.WriteFile(f.path, content, 0644)
}
