package interfaces

import (
	"context"
	"sagara-test/src/product/domain/entity"
)

type IProductRepository interface {
	InsertProduct(ctx context.Context, request entity.ModelProduct) (response entity.ModelProduct, err error)
	DeleteProduct(ctx context.Context, data string) (response entity.ModelProduct, err error)
	UpdateProduct(ctx context.Context, request entity.ModelProduct) (response entity.ModelProduct, err error)
	GetProduct(ctx context.Context, request entity.StructQuery) (response []entity.ModelProduct, count int64, err error)
}
