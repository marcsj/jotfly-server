package controller

import (
	repo "github.com/marcsj/jotfly-server/repo"
	"github.com/marcsj/jotfly-server/users"
	"github.com/marcsj/jotfly-server/util"
)

type UsersController interface {
	Login(userID string, password string) (string, error)
	GetDirectories(userID string) ([]string, error)
	CreateUser(userID string, password string, role users.Role) error
}

type usersController struct {
	repo repo.UsersRepo
	key  []byte
}

func NewUsersController(repo repo.UsersRepo, key []byte) UsersController {
	return &usersController{
		repo: repo,
		key: key,
	}
}

func (c usersController) Login(userID string, password string) (string, error) {
	user, err := c.checkPassword(userID, password)
	if err != nil {
		return "", err
	}
	return util.CreateToken(userID, user.GetRole(), c.key)
}

func (c usersController) GetDirectories(userID string) ([]string, error) {
	user, err := c.repo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user.GetDirectories(), nil
}

func (c usersController) CreateUser(userID string, password string, role users.Role) error {
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

func (c usersController) checkPassword(userID string, password string) (*users.UserInfo, error) {
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