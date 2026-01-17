package domain

import "time"

// Subscription represents a user's subscription to service notifications.
type Subscription struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	ServiceIDs []string  `json:"service_ids"`
	CreatedAt  time.Time `json:"created_at"`
}
