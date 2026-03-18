package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		want    *Config
		wantErr bool
	}{
		{
			name: "all environment variables set correctly",
			env: map[string]string{
				"GH_TOKEN":            "ghp_test_token",
				"GHPP_OWNER":          "my-org",
				"GHPP_PROJECT_NUMBER": "42",
				"GHPP_STATUS_INBOX":   "Inbox",
				"GHPP_STATUS_PLAN":    "Plan",
				"GHPP_STATUS_READY":   "Ready",
				"GHPP_STATUS_DOING":   "Doing",
				"GHPP_PLAN_LIMIT":     "5",
			},
			want: &Config{
				Token:         "ghp_test_token",
				Owner:         "my-org",
				ProjectNumber: 42,
				StatusInbox:   "Inbox",
				StatusPlan:    "Plan",
				StatusReady:   "Ready",
				StatusDoing:   "Doing",
				PlanLimit:     5,
			},
		},
		{
			name: "only required variables set",
			env: map[string]string{
				"GH_TOKEN":            "ghp_token",
				"GHPP_OWNER":          "owner",
				"GHPP_PROJECT_NUMBER": "1",
			},
			want: &Config{
				Token:         "ghp_token",
				Owner:         "owner",
				ProjectNumber: 1,
			},
		},
		{
			name:    "missing GH_TOKEN",
			env:     map[string]string{},
			wantErr: true,
		},
		{
			name: "missing GHPP_OWNER",
			env: map[string]string{
				"GH_TOKEN": "ghp_token",
			},
			wantErr: true,
		},
		{
			name: "missing GHPP_PROJECT_NUMBER",
			env: map[string]string{
				"GH_TOKEN":   "ghp_token",
				"GHPP_OWNER": "owner",
			},
			wantErr: true,
		},
		{
			name: "GHPP_PROJECT_NUMBER is not a number",
			env: map[string]string{
				"GH_TOKEN":            "ghp_token",
				"GHPP_OWNER":          "owner",
				"GHPP_PROJECT_NUMBER": "abc",
			},
			wantErr: true,
		},
		{
			name: "GHPP_PLAN_LIMIT is not a number",
			env: map[string]string{
				"GH_TOKEN":            "ghp_token",
				"GHPP_OWNER":          "owner",
				"GHPP_PROJECT_NUMBER": "1",
				"GHPP_PLAN_LIMIT":     "xyz",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range []string{
				"GH_TOKEN", "GHPP_OWNER", "GHPP_PROJECT_NUMBER",
				"GHPP_STATUS_INBOX", "GHPP_STATUS_PLAN",
				"GHPP_STATUS_READY", "GHPP_STATUS_DOING",
				"GHPP_PLAN_LIMIT",
			} {
				t.Setenv(key, "")
			}

			for k, v := range tt.env {
				t.Setenv(k, v)
			}

			got, err := Load()
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got.Token != tt.want.Token {
				t.Errorf("Token = %q, want %q", got.Token, tt.want.Token)
			}
			if got.Owner != tt.want.Owner {
				t.Errorf("Owner = %q, want %q", got.Owner, tt.want.Owner)
			}
			if got.ProjectNumber != tt.want.ProjectNumber {
				t.Errorf("ProjectNumber = %d, want %d", got.ProjectNumber, tt.want.ProjectNumber)
			}
			if got.StatusInbox != tt.want.StatusInbox {
				t.Errorf("StatusInbox = %q, want %q", got.StatusInbox, tt.want.StatusInbox)
			}
			if got.StatusPlan != tt.want.StatusPlan {
				t.Errorf("StatusPlan = %q, want %q", got.StatusPlan, tt.want.StatusPlan)
			}
			if got.StatusReady != tt.want.StatusReady {
				t.Errorf("StatusReady = %q, want %q", got.StatusReady, tt.want.StatusReady)
			}
			if got.StatusDoing != tt.want.StatusDoing {
				t.Errorf("StatusDoing = %q, want %q", got.StatusDoing, tt.want.StatusDoing)
			}
			if got.PlanLimit != tt.want.PlanLimit {
				t.Errorf("PlanLimit = %d, want %d", got.PlanLimit, tt.want.PlanLimit)
			}
		})
	}
}
