package config

import (
	"fmt"
	"os"
	"strconv"
)

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

	planLimit := 0
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
		StatusInbox:   os.Getenv("GHPP_STATUS_INBOX"),
		StatusPlan:    os.Getenv("GHPP_STATUS_PLAN"),
		StatusReady:   os.Getenv("GHPP_STATUS_READY"),
		StatusDoing:   os.Getenv("GHPP_STATUS_DOING"),
		PlanLimit:     planLimit,
	}, nil
}
