package domain

import "time"

type ChannelType string

const (
	ChannelTypeEmail    ChannelType = "email"
	ChannelTypeTelegram ChannelType = "telegram"
)

type NotificationChannel struct {
	ID         string
	UserID     string
	Type       ChannelType
	Target     string
	IsEnabled  bool
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
