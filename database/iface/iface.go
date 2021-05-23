package iface

import "context"

type IDatabase interface {
	Create(ctx context.Context, database, table string, entity interface{}) error
	GetByID(ctx context.Context, database, table string, entityId string, entity interface{}) error
	GetAll(ctx context.Context, database, table string, entities interface{}) error
	Update(ctx context.Context, enitty string) error
	Delete(ctx context.Context, entityId string) error
}
