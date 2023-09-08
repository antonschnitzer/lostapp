package postgres

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend/env"
	"backend/store/postgres/models"
)

const maxRetries, retryTimeout = 10, time.Second * 15

func NewClient() (*gorm.DB, error) {
	client, err := newGormClient(env.POSTGRES_DB_LOSTSTATS.Value())
	if err != nil {
		return nil, err
	}

	if err = client.AutoMigrate(
		&models.LostClan{},
		&models.Member{},
		&models.Kickpoint{},
		&models.LostClan{},
		&models.LostClanSettings{},
		&models.Notification{},
		&models.NotificationReceiver{},
	); err != nil {
		return nil, err
	}

	log.Printf("Auto-migrated postgres models.")
	return client, nil
}

func newGormClient(dbName string) (client *gorm.DB, err error) {
	dsn := env.POSTGRES_URL.Value() + dbName

	for i := 0; i < maxRetries; i++ {
		if client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
			log.Printf("Failed to connect to database '%s': %v\nRetrying in %s...", dbName, err, retryTimeout.String())
			time.Sleep(retryTimeout)
			continue
		}

		log.Printf("Connected to postgres database '%s'.", dbName)
		return
	}

	return nil, errors.New("max retries reached")
}
