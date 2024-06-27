package dao

import (
	"context"

	"gorm.io/gorm"
)

type IHelloDAO interface {
	GetHelloById(ctx context.Context, id int64) (Hello, error)
}

type HelloDAO struct {
	db *gorm.DB
}

func NewHelloDAO(db *gorm.DB) IHelloDAO {
	return &HelloDAO{db: db}
}

func (h HelloDAO) GetHelloById(ctx context.Context, id int64) (Hello, error) {
	return Hello{}, nil
}
