package activity

import (
	"context"
	"fmt"
	"sloth/fragment"
)

type activityService struct {
	store Store
}

func NewService(store Store) Service {
	return &activityService{
		store: store,
	}
}

func (s *activityService) Insert(ctx context.Context, input ActivityCreateInput) (Activity, error) {
	if input.StartedAt.After(input.EndedAt) {
		return Activity{}, fmt.Errorf("startedAt=%v is before endedAt=%v", input.StartedAt, input.EndedAt)
	}

	result, err := s.store.Insert(ctx, input)
	if err != nil {
		return Activity{}, err
	}

	return result, nil
}

func (s *activityService) All(ctx context.Context, input fragment.PaginationInput) (Activities, error) {
	all, err := s.store.All(ctx, input)
	if err != nil {
		return all, err
	}

	return all, err
}