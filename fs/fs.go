package fs

import (
	"errors"
	"io/ioutil"
	"os"
)

// GetFileFromVirtualFSError vfs.ReadFileSync fail
const GetFileFromVirtualFSError = "open file from vfs error"

// GetFileFromFSError ioutil.ReadFile fail
const GetFileFromFSError = "open file from user fs error"

type fs struct {
	vfs VFS
}

// VFS VirtualFS main object
type VFS interface {
	ReadFile(path string) ([]byte, error)
}

// FS file system interface
type FS interface {
	GetFileFromVirtualFS(path string) (string, error)
	GetFileFromFS(path string) (string, error)
	GetCurrentDirectory() (string, error)
}

// NewFs return new file system with virtual file system
func NewFs(vfs VFS) FS {
	return &fs{
		vfs: vfs,
	}
}

// GetFileFromVirtualFS return a file from virtual fs
func (vfs *fs) GetFileFromVirtualFS(path string) (string, error) {
	file, err := vfs.vfs.ReadFile(path)
	if err != nil {
		return "", errors.New(GetFileFromVirtualFSError)
	}

	return string(file), nil
}

// GetFileFromFS return a file from user fs
func (vfs *fs) GetFileFromFS(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.New(GetFileFromFSError)
	}

	return string(file), nil
}

// GetCurrentDirectory return user current directory
func (vfs *fs) GetCurrentDirectory() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", errors.New(GetFileFromFSError)
	}

	return path, nil
}
