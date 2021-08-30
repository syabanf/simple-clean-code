package domain

import (
	"context"
	"sagara-test/src/product/domain/entity"
	"sagara-test/src/product/domain/interfaces"
	"time"
)

func InsertProduct(ctx context.Context, repo interfaces.IProductRepository, request entity.ModelProduct) (response entity.ModelProduct, err error) {

	createTime := time.Now()
	request.CreatedAt = createTime
	response, err = repo.InsertProduct(ctx, request)
	return
}

func UpdateProduct(ctx context.Context, repo interfaces.IProductRepository, request entity.ModelProduct) (err error) {
	updateTime := time.Now()
	request.UpdatedAt = &updateTime
	_, err = repo.UpdateProduct(ctx, request)
	return
}

func RemoveProduct(ctx context.Context, repo interfaces.IProductRepository, request string) (err error) {
	_, err = repo.DeleteProduct(ctx, request)
	return
}
