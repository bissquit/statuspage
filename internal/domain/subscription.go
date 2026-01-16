package domain

import "time"

type Subscription struct {
	ID         string
	UserID     string
	ServiceIDs []string
	CreatedAt  time.Time
}
