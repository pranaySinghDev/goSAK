package iface

import "context"

type IDatabase interface {
	Create(ctx context.Context, database, table string, entity interface{}) (string, error)
	GetByID(ctx context.Context, database, table string, entityId string, entity interface{}) error
	GetAll(ctx context.Context, database, table string, entities interface{}, limit int64, index int64) error
	Update(ctx context.Context, enitty string) error
	Delete(ctx context.Context, entityId string) error
}
