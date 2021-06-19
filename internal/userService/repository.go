package userService

import (
	"github.com/jameycribbs/hare"
	"github.com/jameycribbs/hare/datastores/disk"
	"log"
)

func InitDB() *hare.Database {
	ds, err := disk.New("./data", ".json")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := hare.New(ds)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.CreateTable("users")
	return db
}
