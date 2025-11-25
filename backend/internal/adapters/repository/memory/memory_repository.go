package memory

import (
	"sync"
	"time"

	"regtech-backend/internal/core/domain"
)

type MemoryDeadlineRepository struct {
	mu        sync.RWMutex
	deadlines []domain.Deadline
}

func NewMemoryDeadlineRepository() *MemoryDeadlineRepository {
	now := time.Now()
	return &MemoryDeadlineRepository{
		deadlines: []domain.Deadline{
			{
				ID:       "sample-asic",
				Title:    "ASIC Annual Review",
				Category: "ASIC",
				DueDate:  now.AddDate(0, 0, 14),
				Status:   domain.StatusPending,
				Notes:    "Check company details & pay fee.",
			},
			{
				ID:       "sample-tax",
				Title:    "Quarterly BAS Lodgement",
				Category: "ATO_TAX",
				DueDate:  now.AddDate(0, 1, 0),
				Status:   domain.StatusPending,
				Notes:    "Coordinate with accountant.",
			},
			{
				ID:       "payroll-super",
				Title:    "Payroll Superannuation",
				Category: "PAYROLL",
				DueDate:  now.AddDate(0, 0, -2),
				Status:   domain.StatusOverdue,
				Notes:    "Submit SG payment to all employees.",
			},
			{
				ID:       "ato-payg",
				Title:    "PAYG Withholding Report",
				Category: "ATO_TAX",
				DueDate:  now, // today
				Status:   domain.StatusOverdue,
				Notes:    "Upload monthly withholding statement.",
			},
			{
				ID:       "tax-return",
				Title:    "Company Tax Return",
				Category: "ATO_TAX",
				DueDate:  now.AddDate(0, 0, 7),
				Status:   domain.StatusPending,
				Notes:    "Prepare financial report for accountant.",
			},
			{
				ID:       "payroll-stp",
				Title:    "STP Finalisation",
				Category: "PAYROLL",
				DueDate:  now.AddDate(0, 0, 1),
				Status:   domain.StatusPending,
				Notes:    "Validate EOFY payroll summary.",
			},
			{
				ID:       "workcover",
				Title:    "WorkCover Declaration",
				Category: "INSURANCE",
				DueDate:  now.AddDate(0, 2, 0),
				Status:   domain.StatusPending,
				Notes:    "Verify total wages for the year.",
			},
			{
				ID:       "tpar",
				Title:    "TPAR Annual Report",
				Category: "ATO_TAX",
				DueDate:  now.AddDate(0, -1, 0),
				Status:   domain.StatusOverdue,
				Notes:    "Submit contractor payment report.",
			},
			{
				ID:       "director-id",
				Title:    "Director ID Verification",
				Category: "ASIC",
				DueDate:  now.AddDate(0, 0, -10),
				Status:   domain.StatusCompleted,
				Notes:    "Already submitted to ASIC.",
			},
			{
				ID:       "insurance-premium",
				Title:    "Business Insurance Premium",
				Category: "INSURANCE",
				DueDate:  now.AddDate(0, 0, 3),
				Status:   domain.StatusPending,
				Notes:    "Renew public liability & cyber policy.",
			},
		},
	}
}

func (r *MemoryDeadlineRepository) GetAll() ([]domain.Deadline, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]domain.Deadline, len(r.deadlines))
	copy(out, r.deadlines)
	return out, nil
}

func (r *MemoryDeadlineRepository) Create(deadline domain.Deadline) (domain.Deadline, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.deadlines = append(r.deadlines, deadline)
	return deadline, nil
}

func (r *MemoryDeadlineRepository) MarkCompleted(id string) (domain.Deadline, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, d := range r.deadlines {
		if d.ID == id {
			r.deadlines[i].Status = domain.StatusCompleted
			return r.deadlines[i], nil
		}
	}
	return domain.Deadline{}, ErrNotFoundMemory()
}

var errNotFound = &notFoundError{}

type notFoundError struct{}

func (e *notFoundError) Error() string { return "deadline not found" }

func ErrNotFoundMemory() error { return errNotFound }
