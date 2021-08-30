package domain

import (
	"context"
	"sagara-test/src/media/domain/entity"
	"sagara-test/src/media/domain/interfaces"
)

func GetMedia(ctx context.Context, repo interfaces.IMediaRepository, request entity.StructQuery) (response []entity.ModelMedia, count int64, err error) {
	response, count, err = repo.GetMedia(ctx, request)
	return
}
