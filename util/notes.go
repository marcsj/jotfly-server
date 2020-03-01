package util

import (
	"encoding/json"
	"github.com/marcsj/jotfly-server/notes"
	"os"
)

func ConvertFileToNote(
	userID string, directory string, id string, file *os.File) (*notes.Note, error) {
	fileNote := &notes.FileNote{}
	decoder := json.NewDecoder(file)
	err := decoder.Decode(fileNote)
	if err != nil {
		return nil, err
	}
	return &notes.Note{
		OwnerId: userID,
		Directory: directory,
		Id: id,
		Content: fileNote.GetContent(),
		Labels: fileNote.GetLabels(),
		 }, nil
}

func ConvertNoteToFileContent(note *notes.Note) (string, error) {
	fileNote := notes.FileNote{
		Labels: note.GetLabels(),
		Content: note.GetContent(),
	}
	content, err := json.Marshal(fileNote)
	if err != nil {
		return "", err
	}
	return string(content), nil
}