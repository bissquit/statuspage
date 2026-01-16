package domain

import "time"

type EventTemplate struct {
	ID            string
	Slug          string
	Type          EventType
	TitleTemplate string
	BodyTemplate  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
