package db

import (
	"FasuhaStats/models"
)

func SyncDatabase() {
	err := DBCon.AutoMigrate(&models.Player{}, &models.Match{}, &models.Player{})
	if err != nil {
		return
	}

}
