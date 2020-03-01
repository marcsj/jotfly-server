package users

import (
	"os"
	"path/filepath"
	"sync"
)

type Repo interface {
	CreateUser(userID string, password []byte) error
	GetUser(userID string) (*UserInfo, error)
	UpdateUser(userID string, user *UserInfo) (*UserInfo, error)
	DeleteUser(userID string) error
}

type fileRepo struct {
	path string
	mx sync.Mutex
}

func NewFileRepo(path string) *fileRepo {
	return &fileRepo{
		path: path,
	}
}

func (r fileRepo) CreateUser(userID string, password []byte) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	_, err := os.Create(filepath.Join(r.path, userID))
	if err != nil {
		return err
	}
	return nil
}

func (r fileRepo) GetUser(userID string) (*UserInfo, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	file, err := os.Open(filepath.Join(r.path, userID))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return convertFileToUser(file)
}

func (r fileRepo) UpdateUser(userID string, user *UserInfo) (*UserInfo, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	file, err := os.OpenFile(
		filepath.Join(r.path, userID),
		os.O_RDWR|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	noteContents, err := convertUserToFileContent(user)
	if err != nil {
		return nil, err
	}
	_, err = file.WriteString(noteContents)
	if err != nil {
		return nil, err
	}

	file, err = os.Open(filepath.Join(r.path, userID))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return convertFileToUser(file)
}

func (r fileRepo) DeleteUser(userID string) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	return os.Remove(filepath.Join(r.path, userID))
}