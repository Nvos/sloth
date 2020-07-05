package resolver

import (
	"sloth/feature/activity"
	"sloth/infra"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Dependencies struct {
	Logger infra.Logger
	ActivityService activity.Service
}

type Resolver struct{
	ActivityService activity.Service
}

func New(deps Dependencies) *Resolver {
	return &Resolver{
		ActivityService: deps.ActivityService,
	}
}