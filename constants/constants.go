package constants

import (
	"fmt"
)

var DbFileName string
var DbFilePath string
var DbFullFilePath string

func SetDbFileName(dbName string) {
	DbFileName = fmt.Sprintf("%s.db", dbName)
}

func SetDbFilePath(path string) {
	DbFilePath = path
}

func SetDbFullFilePath() {
	DbFullFilePath = fmt.Sprintf("%s/%s", DbFilePath, DbFileName)
}
