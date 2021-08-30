package domain

import (
	"context"
	"sagara-test/src/media/domain/entity"
	"sagara-test/src/media/domain/interfaces"
	"time"
)

func InsertMedia(ctx context.Context, repo interfaces.IMediaRepository, request entity.ModelMedia) (response entity.ModelMedia, err error) {

	createTime := time.Now()
	request.CreatedAt = createTime
	response, err = repo.InsertMedia(ctx, request)
	return
}

func UpdateMedia(ctx context.Context, repo interfaces.IMediaRepository, request entity.ModelMedia) (err error) {
	updateTime := time.Now()
	request.UpdatedAt = &updateTime
	_, err = repo.UpdateMedia(ctx, request)
	return
}

func RemoveMedia(ctx context.Context, repo interfaces.IMediaRepository, request string) (err error) {
	_, err = repo.DeleteMedia(ctx, request)
	return
}
