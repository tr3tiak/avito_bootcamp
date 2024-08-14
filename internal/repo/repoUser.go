package repo

import (
	"avito_bootcamp/config"
	"avito_bootcamp/internal/entity"
	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlUserRepo struct {
	db *sql.DB
}

func InitRepo() (*MysqlUserRepo, error) {
	cfg := config.InitConfigDB()
	db, err := sql.Open("mysql", cfg.UserDB+":"+cfg.PasswordDB+"@/"+cfg.NameDB)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &MysqlUserRepo{
		db: db,
	}, nil
}

func (repo *MysqlUserRepo) Create(user *entity.User) error {
	_, err := repo.db.Exec("INSERT INTO Users(Name, Password, Role) VALUES (?,?,?)", user.Name, user.Password, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MysqlUserRepo) Get(userName string) (*entity.User, error) {
	var User entity.User
	rows, err := repo.db.Query("SELECT Name, Password, Role FROM Users WHERE Name = ?", userName)
	if err != nil {
		return &User, err
	}

	rows.Next()
	rows.Scan(&User.Name, &User.Password, &User.Role)
	return &User, nil

}
