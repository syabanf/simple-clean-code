package domain

import (
	"context"
	"sagara-test/src/product/domain/entity"
	"sagara-test/src/product/domain/interfaces"
)

func GetProduct(ctx context.Context, repo interfaces.IProductRepository, request entity.StructQuery) (response []entity.ModelProduct, count int64, err error) {
	response, count, err = repo.GetProduct(ctx, request)
	return
}
