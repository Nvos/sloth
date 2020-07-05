package activity

import (
	"sloth/fragment"
)

type Activity struct {
	ID          int                    `json:"id"`
	Range       fragment.DateTimeRange `json:"range"`
	Description string                 `json:"description"`
}

type Activities = []Activity

func (Activity) IsNode() {}
func (Activity) IsActivityCreateResult() {}
