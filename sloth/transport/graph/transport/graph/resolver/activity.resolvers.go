package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	model1 "sloth/transport/graph/model"
)

func (r *mutationResolver) ActivityCreate(ctx context.Context, input model1.ActivityCreateInput) (*model1.ActivityCreatePayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Activities(ctx context.Context, first *int, last *int, after *string, before *string) (*model1.ActivityConnection, error) {
	panic(fmt.Errorf("not implemented"))
}
