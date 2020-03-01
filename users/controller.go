package users

type Controller interface {
	Login(userID string, password string) (string, error)
	GetDirectories(userID string) ([]string, error)
	CreateUser(userID string, password string, role Role) error
}

type controller struct {
	repo Repo
	key []byte
}

func NewController(repo Repo, key []byte) Controller {
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
	return createToken(userID, user.GetRole(), c.key)
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

func (c controller) CreateUser(userID string, password string, role Role) error {
	salt, err := generateRandomBytes(saltLength)
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

func (c controller) checkPassword(userID string, password string) (*UserInfo, error) {
	user, err := c.repo.GetUser(userID)
	if err != nil {
		return nil, err
	}
	err = checkPassword(password, user.GetPassword(), user.GetSalt())
	if err != nil {
		return nil, err
	}
	return user, nil
}