package users

type Controller interface {

}

type controller struct {
	repo Repo
}

func NewController(repo Repo) Controller {
	return &controller{
		repo: repo,
	}
}

func (c controller) Login(userID string, password []byte) (*Session, error) {
	return nil, nil
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

func (c controller) CreateUser(userID string, password string, role UserInfo_Role) error {
	salt, err := generateRandomBytes(16)
	if err != nil {
		return err
	}
	hashedPass := generatePassword(password, salt)
	err = c.repo.CreateUser(userID, hashedPass)
	if err != nil {
		return err
	}
	_, err = c.repo.UpdateUser(userID, &UserInfo{
		Password: hashedPass,
		Salt: salt,
		Role: role,
		Directories: make([]string, 0),
	})
	return err
}

func (c controller) checkPassword(userID string, password string) error {
	user, err := c.repo.GetUser(userID)
	if err != nil {
		return err
	}
	return checkPassword(password, user.GetPassword(), user.GetSalt())
}