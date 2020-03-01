package users

import (
	"os"
	"path/filepath"
)

type Repo interface {
	CreateUser(userID string, password []byte) error
	GetUser(userID string) (*UserInfo, error)
	UpdateUser(userID string, user *UserInfo) (*UserInfo, error)
	DeleteUser(userID string) error
}

type fileRepo struct {
	path string
}

func NewFileRepo(path string) *fileRepo {
	return &fileRepo{
		path: path,
	}
}

func (r fileRepo) CreateUser(userID string, password []byte) error {
	_, err := os.Create(filepath.Join(r.path, userID))
	if err != nil {
		return err
	}
	return nil
}

func (r fileRepo) GetUser(userID string) (*UserInfo, error) {
	file, err := os.Open(filepath.Join(r.path, userID))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return convertFileToUser(file)
}

func (r fileRepo) UpdateUser(userID string, user *UserInfo) (*UserInfo, error) {
	file, err := os.OpenFile(
		filepath.Join(r.path, userID),
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
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
	return r.GetUser(userID)
}

func (r fileRepo) DeleteUser(userID string) error {
	return os.Remove(filepath.Join(r.path, userID))
}