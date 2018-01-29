package twitter

import (
	"os"
)

func GetFileSize(file *os.File) (int, error) {
	fileStatus, err := file.Stat()
	if err != nil {
		return 0, nil
	}
	return int(fileStatus.Size()), nil
}
