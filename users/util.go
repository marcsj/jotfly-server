package users

import (
	"encoding/json"
	"os"
)

func convertFileToUser(file *os.File) (*UserInfo, error) {
	userInfo := &UserInfo{}
	decoder := json.NewDecoder(file)
	err := decoder.Decode(userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func convertUserToFileContent(user *UserInfo) (string, error) {
	content, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
