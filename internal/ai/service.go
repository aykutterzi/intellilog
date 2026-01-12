package ai

import (
	"strings"

	"github.com/aykutterzi/intellilog/internal/models"
)

type LogAnalysis struct {
	IsAnomaly       bool   `json:"is_anomaly"`
	SeverityScore   int    `json:"severity_score"` // 1-10
	SuggestedFix    string `json:"suggested_fix,omitempty"`
	Prediction      string `json:"prediction,omitempty"`
}

type AIService interface {
	AnalyzeLog(log models.LogEntry) LogAnalysis
}

type SimpleRuleBasedAI struct{}

func NewSimpleRuleBasedAI() *SimpleRuleBasedAI {
	return &SimpleRuleBasedAI{}
}

func (ai *SimpleRuleBasedAI) AnalyzeLog(log models.LogEntry) LogAnalysis {
	analysis := LogAnalysis{
		IsAnomaly:     false,
		SeverityScore: 1,
	}

	// Simulate AI analysis with keywords
	msg := strings.ToLower(log.Message)

	if log.Level == models.LogLevelError {
		analysis.SeverityScore = 8
		analysis.IsAnomaly = true
		analysis.SuggestedFix = "Check system logs for stack trace and restart the service if stuck."
	}

	if strings.Contains(msg, "timeout") || strings.Contains(msg, "latency") {
		analysis.IsAnomaly = true
		analysis.SeverityScore = 7
		analysis.Prediction = "Potential cascading failure if latency persists."
		analysis.SuggestedFix = "Scale up the service or check database connection pool."
	}

	if strings.Contains(msg, "memory") || strings.Contains(msg, "oom") {
		analysis.SeverityScore = 10
		analysis.IsAnomaly = true
		analysis.Prediction = "System crash imminent within 5 minutes."
		analysis.SuggestedFix = "Urgent: Increase memory limit or investigate leak."
	}

	return analysis
}
