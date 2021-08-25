package output

import (
	"io/ioutil"
)

func CreateFile(path string) *File {
	return &File{
		path: path,
	}
}

type File struct {
	path string
}

func (f *File) Write(content []byte) (int, error) {
	return len(content), ioutil.WriteFile(f.path, content, 0644)
}
