package user

import "context"

type Store interface {
	Insert(ctx context.Context, input UserCreateInput)
}

type Service interface {
	Insert(ctx context.Context, input UserCreateInput)
}