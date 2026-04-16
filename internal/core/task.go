package core

import "time"

type Priority int
type Status int

const (
	Low Priority = iota
	Medium
	High
)

const (
	Todo Status = iota
	InProgress
	Done
)

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Priority    Priority  `json:"priority"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	StartDate   time.Time `json:"start_date"` // Para o Calendário
	EndDate     time.Time `json:"end_date"`   // Para o Calendário
	Tags        []string  `json:"tags"`
}
