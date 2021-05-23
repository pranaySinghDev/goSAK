package mongodb

import (
	"github.com/pranaySinghDev/goSAK/database/config"
	"github.com/pranaySinghDev/goSAK/database/iface"
)

type MongoFactory struct{}

// build mongodb
func (f *MongoFactory) Build(config *config.DBConfig) (iface.IDatabase, error) {
	//standard configuration
	db, err := Connect(config.URL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
