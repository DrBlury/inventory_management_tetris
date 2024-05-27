package usecase

import "context"

type Cache interface {
	GetString(ctx context.Context, key string) (string, error)
	SetString(ctx context.Context, key, value string) error
	Invalidate(ctx context.Context, key string) error
}
