package controller

import (
	repo "github.com/marcsj/jotfly-server/repo"
	"github.com/marcsj/jotfly-server/users"
	"github.com/marcsj/jotfly-server/util"
)

type Controller interface {
	Login(userID string, password string) (string, error)
	GetDirectories(userID string) ([]string, error)
	CreateUser(userID string, password string, role users.Role) error
}

type controller struct {
	repo repo.UsersRepo
	key  []byte
}

func NewController(repo repo.UsersRepo, key []byte) Controller {
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
	return util.CreateToken(userID, user.GetRole(), c.key)
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
	salt, err := util.GenerateRandomBytes(util.SaltLength)
	if err != nil {
		return err
	}
	hashedPass := util.GeneratePassword(password, salt)
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
	err = util.CheckPassword(password, user.GetPassword(), user.GetSalt())
	if err != nil {
		return nil, err
	}
	return user, nil
}