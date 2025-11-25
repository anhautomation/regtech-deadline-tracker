package ports

import "regtech-backend/internal/core/domain"

type DeadlineRepository interface {
	GetAll() ([]domain.Deadline, error)
	Create(deadline domain.Deadline) (domain.Deadline, error)
	MarkCompleted(id string) (domain.Deadline, error)
}
