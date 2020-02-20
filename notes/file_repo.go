package notes

type NoteRepo interface {
	CreateNote(userID string, directory string, id string) error
	GetNote(userID string, directory string, id string) (*Note, error)
	UpdateNote(userID string, note *Note) (*Note, error)
	DeleteNote(userID string, directory string, id string) error
	GetUserDirectories(userID string) ([]string, error)
	GetNotesInDirectory(userID string, directory string) ([]string, error)
}

type fileRepo struct {

}

func NewFileRepo() *fileRepo {
	return &fileRepo{

	}
}

func (r fileRepo) CreateNote(userID string, directory string, id string) error {
	return nil
}

func (r fileRepo) GetNote(userID string, directory string, id string) (*Note, error) {
	return nil, nil
}

func (r fileRepo) UpdateNote(userID string, note *Note) (*Note, error) {
	return nil, nil
}

func (r fileRepo) DeleteNote(userID string, directory string, id string) error {
	return nil
}

func (r fileRepo) GetUserDirectories(userID string) ([]string, error) {
	return nil, nil
}

func (r fileRepo) GetNotesInDirectory(userID string, directory string) ([]string, error) {
	return nil, nil
}

