package utils

import (
	"fmt"
	"os"
	"os/user"
	"path"
)

func GetCurrentUserHomeDir() string {
	currentUser, _ := user.Current()
	homeDir := currentUser.HomeDir
	return homeDir
}

func GetResultPath(createFolder bool) string {
	homeDir := GetCurrentUserHomeDir()
	resultPath := path.Join(homeDir, "dsfetcher", "result.json")
	if createFolder {
		err := mkdirIfNotExist(path.Join(homeDir, "dsfetcher"))
		if err != nil {
			fmt.Println("Error creating result directory:", err)
		}
	}
	return resultPath
}

func mkdirIfNotExist(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	return err
}
