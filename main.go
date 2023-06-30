package main

import (
	"github.com/hokageCV/gotrack/db"
	"github.com/hokageCV/gotrack/utils"
)

func main() {
	dbInstance, err := db.InitializeDB()
	utils.CheckNilErr(err)
	defer dbInstance.Close()

	db.DisplayTasksFromDB(dbInstance)
}
