package db

import (
	"context"
	"fmt"
	"sloth/ent"
	entActivity "sloth/ent/activity"
	"sloth/feature/activity"
	"sloth/fragment"
	"strconv"
)

type activityStore struct {
	client *ent.Client
}

func NewActivityStore(client *ent.Client) activity.Store {
	return &activityStore {
		client: client,
	}
}

func (s *activityStore) Insert(ctx context.Context, input activity.ActivityCreateInput) (activity.Activity, error) {
	result, err := s.client.Activity.Create().
		SetDescription(input.Description).
		SetEndedAt(input.EndedAt).
		SetStartedAt(input.StartedAt).
		Save(ctx)

	if err != nil {
		return activity.Activity{}, fmt.Errorf("failed to insert activity, err=%w", err)
	}

	return activity.Activity{
		ID:          result.ID,
		Range: fragment.DateTimeRange{
			StartedAt: result.StartedAt,
			EndedAt:  result.EndedAt,
		},
		Description: result.Description,
	}, nil
}

// TODO, 04.07.2020: verify if directions are correct
func (s *activityStore) All(ctx context.Context, input fragment.PaginationInput) (activity.Activities, error) {
	query := s.client.Activity.Query()

	if input.First != nil {
		after, err := strconv.Atoi(*input.After)
		if err != nil {
			return nil, fmt.Errorf("failed to convert after=%v to int, err=%w", input.After, err)
		}

		query.
			Limit(*input.First).
			Order(ent.Asc(entActivity.FieldID)).
			Where(entActivity.IDLT(after))
	}

	if input.Last != nil {
		query.Limit(*input.Last).Order(ent.Desc(entActivity.FieldID))

		before, err := strconv.Atoi(*input.Before)
		if err != nil {
			return nil, fmt.Errorf("failed to convert before=%v to int, err=%w", input.After, err)
		}

		query.
			Limit(*input.First).
			Order(ent.Asc(entActivity.FieldID)).
			Where(entActivity.IDGTE(before))
	}

	result, err := query.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all activities, err=%w", err)
	}

	var activities activity.Activities
	for i := range result {
		activities = append(activities, activity.Activity{
			ID:          result[i].ID,
			Range:       fragment.DateTimeRange{
				StartedAt: result[i].StartedAt,
				EndedAt:   result[i].EndedAt,
			},
			Description: result[i].Description,
		})
	}

	return activities, nil
}