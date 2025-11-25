package domain

import "time"

type DeadlineStatus string

const (
	StatusPending   DeadlineStatus = "PENDING"
	StatusOverdue   DeadlineStatus = "OVERDUE"
	StatusCompleted DeadlineStatus = "COMPLETED"
)

type Deadline struct {
	ID       string         `json:"id"`
	Title    string         `json:"title"`
	Category string         `json:"category"`
	DueDate  time.Time      `json:"dueDate"`
	Status   DeadlineStatus `json:"status"`
	Notes    string         `json:"notes,omitempty"`
}
