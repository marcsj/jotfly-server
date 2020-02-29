package notes

import (
	"io/ioutil"
	"os"
)

func convertFileToNote(
	userID string, directory string, id string, file *os.File) (*Note, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return &Note{
		OwnerId: userID,
		Directory: directory,
		Id: id,
		Content: string(data),
		 }, nil
}

func convertNoteToFileContent(note *Note) (string, error) {
	content := note.GetContent()
	return content, nil
}