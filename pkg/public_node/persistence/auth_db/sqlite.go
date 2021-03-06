package auth_db

import (
	"github.com/vladimish/soo/pkg/logger"
	"github.com/vladimish/soo/pkg/public_node/auth/models"
	node_models "github.com/vladimish/soo/pkg/public_node/node/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
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
	glog := gormLogger.New(log.New(os.Stdout, "\n", log.LstdFlags), gormLogger.Config{
		SlowThreshold:             0,
		Colorful:                  true,
		IgnoreRecordNotFoundError: false,
		LogLevel:                  gormLogger.Info,
	})
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{
		Logger: glog,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.AuthData{})
	db.AutoMigrate(node_models.Node{})
	return db, nil
}

func (a *DB) GetNode(nickName string) (*node_models.Node, error) {
	n := &node_models.Node{}
	tx := a.db.Where(node_models.Node{NickName: nickName}).Take(n)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return n, nil
}

func (a *DB) GetNodesLikeOrNil(nickName string, limit int) ([]node_models.Node, error) {
	var nodes []node_models.Node
	tx := a.db.Select("*").Table("nodes").Where("nick_name LIKE ?", "%"+nickName+"%").Limit(limit).Take(&nodes)
	if tx.Error != nil {
		return nil, tx.Error
	}

	logger.L.Sugar().Info(nodes)

	return nodes, nil
}

func (a *DB) SaveNode(node *node_models.Node) error {
	tx := a.db.Save(node)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (a *DB) UpdateNode(nodeId int, column string, value interface{}) error {
	tx := a.db.Model(&node_models.Node{}).Where("id = ?", nodeId).Update(column, value)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (a *DB) GetAuthData(message string) (*models.AuthData, error) {
	ad := &models.AuthData{}
	tx := a.db.Where(models.AuthData{CheckoutMessage: message}).Last(ad)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return ad, nil
}

func (a *DB) GetLastAuthData(node node_models.Node) (*models.AuthData, error) {
	ad := &models.AuthData{}
	tx := a.db.Where(models.AuthData{Node: node}).Last(ad)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return ad, nil
}

func (a *DB) SaveAuth(d *models.AuthData) error {
	tx := a.db.Save(d)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
