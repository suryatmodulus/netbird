package activity

import (
	"time"
)

const (
	SystemInitiator = "sys"
)

// Event represents a network/system activity event.
type Event struct {
	// Timestamp of the event
	Timestamp time.Time
	// Activity that was performed during the event
	Activity Activity
	// ID of the event (can be empty, meaning that it wasn't yet generated)
	ID uint64
	// InitiatorID is the ID of an object that initiated the event (e.g., a user)
	InitiatorID string
	// TargetID is the ID of an object that was effected by the event (e.g., a peer)
	TargetID string
	// AccountID is the ID of an account where the event happened
	AccountID string
	// Meta of the event, e.g. deleted peer information like name, IP, etc
	Meta map[string]any
}

// Copy the event
func (e *Event) Copy() *Event {

	meta := make(map[string]any, len(e.Meta))
	for key, value := range e.Meta {
		meta[key] = value
	}

	return &Event{
		Timestamp:   e.Timestamp,
		Activity:    e.Activity,
		ID:          e.ID,
		InitiatorID: e.InitiatorID,
		TargetID:    e.TargetID,
		AccountID:   e.AccountID,
		Meta:        meta,
	}
}
