package repo

import (
	"avito_bootcamp/config"
	"avito_bootcamp/internal/entity"
	"database/sql"

	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type FlatRepo struct {
	db *sql.DB
}

func InitFlatRepo() (*FlatRepo, error) {
	cfg := config.InitConfigDB()
	db, err := sql.Open("mysql", cfg.UserDB+":"+cfg.PasswordDB+"@/"+cfg.NameDB)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &FlatRepo{
		db: db,
	}, nil
}

func (fr FlatRepo) Post(flat *entity.Flat) error {
	logrus.Info("Post flat repo started")
	result, err := fr.db.Exec("insert into Flats(HouseID, Price, Rooms, Status) values (?, ?, ?, ?)", flat.HouseId, flat.Price, flat.Rooms, flat.Status)
	if err != nil {
		logrus.Error(err)
		return err
	}
	Id, err := result.LastInsertId()
	if err != nil {
		logrus.Error(err)
		return err
	}
	flat.Id = int(Id)
	flat.Status = "on moderate"
	logrus.Info("Post Status flat repo complete")
	return nil
}

func (fr FlatRepo) UpdateStatus(flat *entity.Flat) error {
	logrus.Info("Update flat repo started")

	tx, err := fr.db.Begin()
	if err != nil {
		logrus.Error(err)
		return err
	}

	_, err = tx.Exec("update Flats set Status = ? where ID = ?", flat.Status, flat.Id)
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
		return err
	}
	rows, err := tx.Query("select HouseID, Price, Rooms from Flats where ID = ?", flat.Id)
	if err != nil {
		tx.Rollback()
		logrus.Error(err)
		return err
	}
	for rows.Next() {
		logrus.Info("rows is not empty")
		err := rows.Scan(&flat.HouseId, &flat.Price, &flat.Rooms)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Update Status flat repo complete")
	return nil
}
