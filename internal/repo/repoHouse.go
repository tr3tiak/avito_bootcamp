package repo

import (
	"avito_bootcamp/config"
	"avito_bootcamp/internal/entity"
	"database/sql"

	"time"

	"github.com/sirupsen/logrus"
)

const (
	dbformat     string = "2006-01-02 15:04:05"
	structformat string = "2006-01-02T15:04:05Z"
)

type HouseRepo struct {
	db *sql.DB
}

func InitHouseRepo() (*HouseRepo, error) {
	cfg := config.InitConfigDB()
	db, err := sql.Open("mysql", cfg.UserDB+":"+cfg.PasswordDB+"@/"+cfg.NameDB)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &HouseRepo{
		db: db,
	}, nil
}

func (repo HouseRepo) CreateHouse(house *entity.House) (*entity.House, error) {
	logrus.Info("house repo create started")
	now := time.Now()

	query := "INSERT INTO Houses(Address, Year, Developer, CreatedAt, UpdateAt) values (?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(query, house.Address, house.Year, house.Developer, now.Format(dbformat), now.Format(dbformat))
	if err != nil {
		logrus.Error(err)
		return house, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		logrus.Error(err)
		return house, err
	}
	house.HouseId = int(id)
	house.CreateAt = now.Format(structformat)
	house.UpdateAt = now.Format(structformat)
	logrus.Info("house repo create complete")
	return house, nil
}
