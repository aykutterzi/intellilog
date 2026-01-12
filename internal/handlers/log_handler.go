package handlers

import (
	"net/http"
	"time"

	"github.com/aykutterzi/intellilog/internal/ai"
	"github.com/aykutterzi/intellilog/internal/models"
	"github.com/aykutterzi/intellilog/internal/store"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LogHandler struct {
	store store.LogStore
	ai    ai.AIService
}

func NewLogHandler(store store.LogStore, ai ai.AIService) *LogHandler {
	return &LogHandler{store: store, ai: ai}
}

func (h *LogHandler) CreateLog(c echo.Context) error {
	var log models.LogEntry
	if err := c.Bind(&log); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if log.ID == "" {
		log.ID = uuid.New().String()
	}
	if log.Timestamp.IsZero() {
		log.Timestamp = time.Now()
	}

	// AI Analysis
	analysis := h.ai.AnalyzeLog(log)
	log.Analysis = analysis

	if err := h.store.AddLog(log); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save log"})
	}

	return c.JSON(http.StatusCreated, log)
}

func (h *LogHandler) GetLogs(c echo.Context) error {
	logs, err := h.store.GetLogs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve logs"})
	}
	return c.JSON(http.StatusOK, logs)
}
