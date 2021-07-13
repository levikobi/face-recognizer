package persons

import (
	"context"
	"face-recognition-service/internals/models"
)

//go:generate mockery --name Repository
type Repository interface {
	InsertOne(context context.Context, document interface{}) error
	Find(ctx context.Context, filter interface{}) (*[]models.Person, error)
	Count(ctx context.Context, filter interface{}) (int64, error)
}
