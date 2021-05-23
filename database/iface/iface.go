package iface

import "context"

type IDatabase interface {
	Create(ctx context.Context, database, table string, entity interface{}) error
	Get(ctx context.Context, entityId string) (interface{}, error)
	Update(ctx context.Context, enitty string) error
	Delete(ctx context.Context, entityId string) error
}
