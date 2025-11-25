package http

import (
	"time"

	"github.com/gin-gonic/gin"

	"regtech-backend/internal/adapters/repository/memory"
	"regtech-backend/internal/core/contract"
	"regtech-backend/internal/core/services"
)

type DeadlineHandler struct {
	svc *services.DeadlineService
}

func NewDeadlineHandler(svc *services.DeadlineService) *DeadlineHandler {
	return &DeadlineHandler{svc: svc}
}

type createDeadlineRequest struct {
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"required"`
	DueDate  string `json:"dueDate" binding:"required"`
	Notes    string `json:"notes"`
}

// @Summary      List regulatory deadlines
// @Description  Returns all regulatory deadlines
// @Tags         deadlines
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "Example: {\"code\":\"A00\",\"message\":\"success\",\"data\":[]}"
// @Router       /deadlines [get]
func (h *DeadlineHandler) ListDeadlines(c *gin.Context) {
	code, items := h.svc.ListDeadlines()
	if code != contract.SUCCESS {
		Fail(c, code)
		return
	}
	Success(c, items)
}

// @Summary      Create a new regulatory deadline
// @Description  Example payload:
// @Description  {
// @Description    "title": "Quarterly BAS Lodgement",
// @Description    "category": "ATO_TAX",
// @Description    "dueDate": "2025-01-31",
// @Description    "notes": "Coordinate with accountant"
// @Description  }
// @Tags         deadlines
// @Accept       json
// @Produce      json
// @Param        request  body  object  true  "Deadline payload"
// @Success      200  {object}  map[string]interface{}  "Example: {\"code\":\"A00\",\"message\":\"success\",\"data\":{}}"
// @Router       /deadlines [post]
func (h *DeadlineHandler) CreateDeadline(c *gin.Context) {
	var req createDeadlineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Fail(c, contract.INVALID)
		return
	}

	due, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		Fail(c, contract.INVALID)
		return
	}

	code, item := h.svc.CreateDeadline(services.CreateDeadlineInput{
		Title:    req.Title,
		Category: req.Category,
		DueDate:  due,
		Notes:    req.Notes,
	})
	if code != contract.SUCCESS {
		Fail(c, code)
		return
	}

	Success(c, item)
}

// @Summary      Mark a deadline as completed
// @Tags         deadlines
// @Produce      json
// @Param        id   path  string  true  "Deadline ID"
// @Success      200  {object}  map[string]interface{}  "Example: {\"code\":\"A00\",\"message\":\"success\",\"data\":{}}"
// @Router       /deadlines/{id}/complete [put]
func (h *DeadlineHandler) MarkCompleted(c *gin.Context) {
	id := c.Param("id")

	code, item := h.svc.MarkCompleted(id)
	if code != contract.SUCCESS {
		Fail(c, code)
		return
	}

	Success(c, item)
}

func WireDeadlineHandler() *DeadlineHandler {
	repo := memory.NewMemoryDeadlineRepository()
	svc := services.NewDeadlineService(repo)
	return NewDeadlineHandler(svc)
}
