package auth_db

import (
	"github.com/telf01/soo/pkg/public_node/auth/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDB(connectionString string) (*DB, error) {
	db, err := initializeDB(connectionString)
	if err != nil {
		return nil, err
	}
	data := DB{
		db: db,
	}

	return &data, nil
}

func initializeDB(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.AuthData{})
	return db, nil
}

func (a *DB) GetAuthData(nickName string) (*models.AuthData, error){
	ad := &models.AuthData{}
	tx := a.db.First(ad, "NickName = ?", nickName)
	if tx.Error != nil{
		return nil, tx.Error
	}
	return ad, nil
}

func (a *DB) SaveAuth(d *models.AuthData) error{
	panic("NOT YET IMPLEMENTED")
}