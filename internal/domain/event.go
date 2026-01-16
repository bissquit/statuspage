package domain

import "time"

type EventType string

const (
	EventTypeIncident    EventType = "incident"
	EventTypeMaintenance EventType = "maintenance"
)

type EventStatus string

const (
	EventStatusInvestigating EventStatus = "investigating"
	EventStatusIdentified    EventStatus = "identified"
	EventStatusMonitoring    EventStatus = "monitoring"
	EventStatusResolved      EventStatus = "resolved"
	EventStatusScheduled     EventStatus = "scheduled"
	EventStatusInProgress    EventStatus = "in_progress"
	EventStatusCompleted     EventStatus = "completed"
)

type Severity string

const (
	SeverityMinor    Severity = "minor"
	SeverityMajor    Severity = "major"
	SeverityCritical Severity = "critical"
)

type Event struct {
	ID                string
	Title             string
	Type              EventType
	Status            EventStatus
	Severity          *Severity
	Description       string
	StartedAt         *time.Time
	ResolvedAt        *time.Time
	ScheduledStartAt  *time.Time
	ScheduledEndAt    *time.Time
	NotifySubscribers bool
	TemplateID        *string
	CreatedBy         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ServiceIDs        []string
}

type EventUpdate struct {
	ID                string
	EventID           string
	Status            EventStatus
	Message           string
	NotifySubscribers bool
	CreatedBy         string
	CreatedAt         time.Time
}
