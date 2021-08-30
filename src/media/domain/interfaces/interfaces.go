package interfaces

import (
	"context"
	"sagara-test/src/media/domain/entity"
)

type IMediaRepository interface {
	InsertMedia(ctx context.Context, request entity.ModelMedia) (response entity.ModelMedia, err error)
	DeleteMedia(ctx context.Context, data string) (response entity.ModelMedia, err error)
	UpdateMedia(ctx context.Context, request entity.ModelMedia) (response entity.ModelMedia, err error)
	GetMedia(ctx context.Context, request entity.StructQuery) (response []entity.ModelMedia, count int64, err error)
}
