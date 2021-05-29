package constants

import (
	"fmt"
	"log"
	"os"
)

var DbFileName string
var DbFilePath string
var DbFullFilePath string

func SetDbFileName(dbName string) {
	DbFileName = fmt.Sprintf("%s.db", dbName)
}

func SetDbFilePath() {
	var err error
	DbFilePath, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	SetDbFilePath()
}

func SetDbFullFilePath() {
	DbFullFilePath = fmt.Sprintf("%s%s", DbFilePath, DbFileName)
}
