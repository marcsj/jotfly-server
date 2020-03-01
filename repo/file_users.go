package repo

import (
	"github.com/marcsj/jotfly-server/users"
	"github.com/marcsj/jotfly-server/util"
	"os"
	"path/filepath"
	"sync"
)

type UsersRepo interface {
	CreateUser(userID string, password []byte) error
	GetUser(userID string) (*users.UserInfo, error)
	UpdateUser(userID string, user *users.UserInfo) (*users.UserInfo, error)
	DeleteUser(userID string) error
}

type fileUsersRepo struct {
	path string
	mx sync.Mutex
}

func NewFileUsersRepo(path string) *fileUsersRepo {
	return &fileUsersRepo{
		path: path,
	}
}

func (r fileUsersRepo) CreateUser(userID string, password []byte) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	_, err := os.Create(filepath.Join(r.path, userID))
	if err != nil {
		return err
	}
	return nil
}

func (r fileUsersRepo) GetUser(userID string) (*users.UserInfo, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	file, err := os.Open(filepath.Join(r.path, userID))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return util.ConvertFileToUser(file)
}

func (r fileUsersRepo) UpdateUser(userID string, user *users.UserInfo) (*users.UserInfo, error) {
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

	noteContents, err := util.ConvertUserToFileContent(user)
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
	return util.ConvertFileToUser(file)
}

func (r fileUsersRepo) DeleteUser(userID string) error {
	r.mx.Lock()
	defer r.mx.Unlock()
	return os.Remove(filepath.Join(r.path, userID))
}