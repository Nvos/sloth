package activity

import (
	"context"
	"sloth/fragment"
)

type Store interface {
	Insert(ctx context.Context, input ActivityCreateInput) (Activity, error)
	All(ctx context.Context, input fragment.PaginationInput) (Activities, error)
}

type Service interface {
	Insert(ctx context.Context, input ActivityCreateInput) (Activity, error)
	All(ctx context.Context, input fragment.PaginationInput) (Activities, error)
}