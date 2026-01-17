package domain

import "time"

// EventType represents the type of an event.
type EventType string

// Event types.
const (
	EventTypeIncident    EventType = "incident"
	EventTypeMaintenance EventType = "maintenance"
)

// EventStatus represents the current status of an event.
type EventStatus string

// Event statuses.
const (
	EventStatusInvestigating EventStatus = "investigating"
	EventStatusIdentified    EventStatus = "identified"
	EventStatusMonitoring    EventStatus = "monitoring"
	EventStatusResolved      EventStatus = "resolved"
	EventStatusScheduled     EventStatus = "scheduled"
	EventStatusInProgress    EventStatus = "in_progress"
	EventStatusCompleted     EventStatus = "completed"
)

// Severity represents the severity level of an event.
type Severity string

// Severity levels.
const (
	SeverityMinor    Severity = "minor"
	SeverityMajor    Severity = "major"
	SeverityCritical Severity = "critical"
)

// Event represents an incident or maintenance event.
type Event struct {
	ID                string       `json:"id"`
	Title             string       `json:"title"`
	Type              EventType    `json:"type"`
	Status            EventStatus  `json:"status"`
	Severity          *Severity    `json:"severity"`
	Description       string       `json:"description"`
	StartedAt         *time.Time   `json:"started_at"`
	ResolvedAt        *time.Time   `json:"resolved_at"`
	ScheduledStartAt  *time.Time   `json:"scheduled_start_at"`
	ScheduledEndAt    *time.Time   `json:"scheduled_end_at"`
	NotifySubscribers bool         `json:"notify_subscribers"`
	TemplateID        *string      `json:"template_id"`
	CreatedBy         string       `json:"created_by"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	ServiceIDs        []string     `json:"service_ids"`
}

// EventUpdate represents a status update for an event.
type EventUpdate struct {
	ID                string      `json:"id"`
	EventID           string      `json:"event_id"`
	Status            EventStatus `json:"status"`
	Message           string      `json:"message"`
	NotifySubscribers bool        `json:"notify_subscribers"`
	CreatedBy         string      `json:"created_by"`
	CreatedAt         time.Time   `json:"created_at"`
}
