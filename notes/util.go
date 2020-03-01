package notes

import (
	"encoding/json"
	"os"
)

func convertFileToNote(
	userID string, directory string, id string, file *os.File) (*Note, error) {
	fileNote := &FileNote{}
	decoder := json.NewDecoder(file)
	err := decoder.Decode(fileNote)
	if err != nil {
		return nil, err
	}
	return &Note{
		OwnerId: userID,
		Directory: directory,
		Id: id,
		Content: fileNote.GetContent(),
		Labels: fileNote.GetLabels(),
		 }, nil
}

func convertNoteToFileContent(note *Note) (string, error) {
	fileNote := FileNote{
		Labels: note.GetLabels(),
		Content: note.GetContent(),
	}
	content, err := json.Marshal(fileNote)
	if err != nil {
		return "", err
	}
	return string(content), nil
}