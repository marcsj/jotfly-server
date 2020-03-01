package controller

import (
	repo2 "github.com/marcsj/jotfly-server/repo"
	"github.com/marcsj/jotfly-server/users"
)

type Controller interface {
	Login(userID string, password string) (string, error)
	GetDirectories(userID string) ([]string, error)
	CreateUser(userID string, password string, role users.Role) error
}

type controller struct {
	repo repo2.Repo
	key  []byte
}

func NewController(repo repo2.Repo, key []byte) Controller {
	return &controller{
		repo: repo,
		key: key,
	}
}

func (c controller) Login(userID string, password string) (string, error) {
	user, err := c.checkPassword(userID, password)
	if err != nil {
		return "", err
	}
	return users.createToken(userID, user.GetRole(), c.key)
}

func (c controller) GetDirectories(userID string) ([]string, error) {
	user, err := c.repo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user.GetDirectories(), nil
}

func (c controller) addDirectory(userID string, directory string) error {
	user, err := c.repo.GetUser(userID)
	if err != nil {
		return err
	}
	user.Directories = append(user.Directories, directory)
	_, err = c.repo.UpdateUser(userID, user)
	return err
}

func (c controller) CreateUser(userID string, password string, role users.Role) error {
	salt, err := users.generateRandomBytes(users.saltLength)
	if err != nil {
		return err
	}
	hashedPass := users.generatePassword(password, salt)
	err = c.repo.CreateUser(userID, hashedPass)
	if err != nil {
		return err
	}
	_, err = c.repo.UpdateUser(userID, &users.UserInfo{
		Password: hashedPass,
		Salt: salt,
		Role: role,
		Directories: make([]string, 0),
	})
	return err
}

func (c controller) checkPassword(userID string, password string) (*users.UserInfo, error) {
	user, err := c.repo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	err = users.checkPassword(password, user.GetPassword(), user.GetSalt())
	if err != nil {
		return nil, err
	}
	return user, nil
}