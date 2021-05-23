package database

import (
	"github.com/pranaySinghDev/goSAK/database/config"
	"github.com/pranaySinghDev/goSAK/database/iface"
	"github.com/pranaySinghDev/goSAK/database/mongodb"
)

type databaseFactory interface {
	Build(*config.DBConfig) (iface.IDatabase, error)
}

var databaseFactoryMap = map[config.DBType]databaseFactory{
	config.Mongodb:     &mongodb.MongoFactory{},
	config.PostgresSQL: nil,
}

func Build(config *config.DBConfig) (iface.IDatabase, error) {
	return databaseFactoryMap[config.Type].Build(config)
}
