package services

import (
	"strings"
	"time"

	"github.com/google/uuid"

	"regtech-backend/internal/core/contract"
	"regtech-backend/internal/core/domain"
	"regtech-backend/internal/core/ports"
)

type DeadlineService struct {
	repo ports.DeadlineRepository
}

func NewDeadlineService(repo ports.DeadlineRepository) *DeadlineService {
	return &DeadlineService{repo: repo}
}

func (s *DeadlineService) ListDeadlines() (string, []domain.Deadline) {
	items, err := s.repo.GetAll()
	if err != nil {
		return contract.INTERNAL, nil
	}
	return contract.SUCCESS, items
}

type CreateDeadlineInput struct {
	Title    string
	Category string
	DueDate  time.Time
	Notes    string
}

func (s *DeadlineService) CreateDeadline(input CreateDeadlineInput) (string, domain.Deadline) {
	if strings.TrimSpace(input.Title) == "" || strings.TrimSpace(input.Category) == "" {
		return contract.INVALID, domain.Deadline{}
	}

	existing, err := s.repo.GetAll()
	if err != nil {
		return contract.INTERNAL, domain.Deadline{}
	}
	for _, d := range existing {
		if strings.EqualFold(d.Title, input.Title) && strings.EqualFold(d.Category, input.Category) {
			return contract.ALREADY_EXISTS, domain.Deadline{}
		}
	}

	d := domain.Deadline{
		ID:       uuid.NewString(),
		Title:    input.Title,
		Category: input.Category,
		DueDate:  input.DueDate,
		Status:   domain.StatusPending,
		Notes:    input.Notes,
	}

	created, err := s.repo.Create(d)
	if err != nil {
		return contract.INTERNAL, domain.Deadline{}
	}

	return contract.SUCCESS, created
}

func (s *DeadlineService) MarkCompleted(id string) (string, domain.Deadline) {
	if strings.TrimSpace(id) == "" {
		return contract.INVALID, domain.Deadline{}
	}

	updated, err := s.repo.MarkCompleted(id)
	if err != nil {
		return contract.NOT_FOUND, domain.Deadline{}
	}

	return contract.SUCCESS, updated
}
