package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"sloth/feature/activity"
	"sloth/fragment"
	"sloth/transport/graph/model"
	"strconv"
)

func (r *mutationResolver) ActivityCreate(ctx context.Context, input activity.ActivityCreateInput) (model.ActivityCreateResult, error) {
	result, err := r.ActivityService.Insert(ctx, input)
	if err != nil {
		return &model.ValidationError{
			Message: err.Error(),
		}, nil
	}

	return result, nil
}

func (r *queryResolver) Activities(ctx context.Context, first *int, last *int, after *string, before *string) (*model.ActivityConnection, error) {
	result, err := r.ActivityService.All(ctx, fragment.PaginationInput{
		First:  first,
		Last:   last,
		After:  after,
		Before: before,
	})

	if err != nil {
		return nil, err
	}

	var edges []*model.ActivityEdge

	for i := range result {
		edges = append(edges, &model.ActivityEdge{
			Node:   &result[i],
			Cursor: strconv.Itoa(result[i].ID),
		})
	}

	return &model.ActivityConnection{
		Edges: edges,
	}, nil
}
