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
	return nil, nil
}

func (c controller) CreateUser(userID string, password []byte, role UserInfo_Role) error {
	return nil
}

func (c controller) checkPassword(userID string, password []byte) error {
	return nil
}