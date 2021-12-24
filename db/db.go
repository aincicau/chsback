package db

import (
	"chsback/common"
	"chsback/config"

	"github.com/golang-migrate/migrate"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func InitDatabase() {
	migrateConnection, err := migrate.New("file://db/migrate", config.GetConfig().Database.URL)

	if err != nil {
		logrus.WithError(err).Error("Error connecting to database")
		return
	}

	version := config.GetConfig().Database.Version
	currentVersion, _, _ := migrateConnection.Version()

	if version != currentVersion {
		err = migrateConnection.Migrate(version)
		if err != nil {
			logrus.WithError(err).Error("Error creating tables")
			return
		}
	}

	migrateConnection.Close()

	db, err = gorm.Open("postgres", config.GetConfig().Database.URL)

	if err != nil {
		logrus.WithError(err).Error(common.DB_ERROR_MESSAGE)
		return
	}

	db.LogMode(config.GetConfig().Database.LogMode)
}
