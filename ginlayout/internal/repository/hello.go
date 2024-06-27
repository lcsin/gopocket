package repository

import (
	"context"

	"github.com/lcsin/ginlayout/internal/domain"
	"github.com/lcsin/ginlayout/internal/repository/cache"
	"github.com/lcsin/ginlayout/internal/repository/dao"
)

type IHelloRepository interface {
	Get(ctx context.Context, id int64) (domain.Hello, error)
}

type HelloRepository struct {
	dao   dao.IHelloDAO
	cache cache.IHelloCache
}

func NewHelloRepository(dao dao.IHelloDAO, cache cache.IHelloCache) IHelloRepository {
	return &HelloRepository{dao: dao, cache: cache}
}

func (h HelloRepository) Get(ctx context.Context, id int64) (domain.Hello, error) {
	return domain.Hello{}, nil
}
