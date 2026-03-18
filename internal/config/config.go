package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	DefaultStatusInbox = "Backlog"
	DefaultStatusPlan  = "Plan"
	DefaultStatusReady = "Ready"
	DefaultStatusDoing = "In progress"
	DefaultPlanLimit   = 3
)

// getEnvOrDefault returns the value of the environment variable named by the key,
// or the default value if the variable is not set or empty.
func getEnvOrDefault(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

// Config holds application configuration loaded from environment variables.
type Config struct {
	Token         string
	Owner         string
	ProjectNumber int
	StatusInbox   string
	StatusPlan    string
	StatusReady   string
	StatusDoing   string
	PlanLimit     int
}

// Load reads environment variables and returns a Config.
// Required variables: GH_TOKEN, GHPP_OWNER, GHPP_PROJECT_NUMBER.
func Load() (*Config, error) {
	token := os.Getenv("GH_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("failed to load config: GH_TOKEN is required")
	}

	owner := os.Getenv("GHPP_OWNER")
	if owner == "" {
		return nil, fmt.Errorf("failed to load config: GHPP_OWNER is required")
	}

	projectNumberStr := os.Getenv("GHPP_PROJECT_NUMBER")
	if projectNumberStr == "" {
		return nil, fmt.Errorf("failed to load config: GHPP_PROJECT_NUMBER is required")
	}

	projectNumber, err := strconv.Atoi(projectNumberStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse GHPP_PROJECT_NUMBER: %w", err)
	}

	planLimit := DefaultPlanLimit
	if v := os.Getenv("GHPP_PLAN_LIMIT"); v != "" {
		planLimit, err = strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("failed to parse GHPP_PLAN_LIMIT: %w", err)
		}
	}

	return &Config{
		Token:         token,
		Owner:         owner,
		ProjectNumber: projectNumber,
		StatusInbox:   getEnvOrDefault("GHPP_STATUS_INBOX", DefaultStatusInbox),
		StatusPlan:    getEnvOrDefault("GHPP_STATUS_PLAN", DefaultStatusPlan),
		StatusReady:   getEnvOrDefault("GHPP_STATUS_READY", DefaultStatusReady),
		StatusDoing:   getEnvOrDefault("GHPP_STATUS_DOING", DefaultStatusDoing),
		PlanLimit:     planLimit,
	}, nil
}
