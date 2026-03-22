package model

import "time"

type Company struct {
	Id              string
	UserId          string
	Name            string
	Industry        *string
	JobType         *string
	PreferenceLevel *int
	Memo            *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
