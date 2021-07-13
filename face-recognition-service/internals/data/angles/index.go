package angles

import "context"

type Repository interface {
	Get(ctx context.Context, key string) (map[string]float64, bool, error)
	Set(ctx context.Context, key string, value *map[string]float64) error
}
