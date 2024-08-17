package repo

import (
	"avito_bootcamp/config"
	"avito_bootcamp/internal/entity"
	"database/sql"

	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlUserRepo struct {
	db *sql.DB
}

func InitRepo() (*MysqlUserRepo, error) {
	cfg := config.InitConfigDB()
	db, err := sql.Open("mysql", cfg.UserDB+":"+cfg.PasswordDB+"@/"+cfg.NameDB)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &MysqlUserRepo{
		db: db,
	}, nil
}

func (repo MysqlUserRepo) Create(user *entity.User) error {
	logrus.Info("repo create started")
	_, err := repo.db.Exec("INSERT INTO Users(Name, Password, Role) VALUES (?,?,?)", user.Name, user.Password, user.Role)

	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Info("repo create complete")
	return nil
}

func (repo MysqlUserRepo) Get(userName string) (*entity.User, error) {
	logrus.Info("repo user started")
	var User entity.User
	rows, err := repo.db.Query("SELECT Name, Password, Role FROM Users WHERE Name = ?", userName)
	if err != nil {
		logrus.Error(err)
		return &User, err
	}

	if !rows.Next() {
		logrus.Info("table users is empty")
	}
	err = rows.Scan(&User.Name, &User.Password, &User.Role)
	if err != nil {
		logrus.Error(err)
		return &User, err
	}
	logrus.Info("repo get complete")
	return &User, nil

}
