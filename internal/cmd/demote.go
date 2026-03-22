package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/douhashi/gh-project-promoter/internal/config"
	"github.com/douhashi/gh-project-promoter/internal/demote"
	"github.com/douhashi/gh-project-promoter/internal/github"
)

// RunDemote fetches project items via the API, runs the demotion logic, and prints the results as JSON.
func RunDemote(ctx context.Context, cfg *config.Config, demoter github.ItemPromoter) error {
	items, err := demoter.FetchProjectItems(ctx, cfg.Owner, cfg.ProjectNumber)
	if err != nil {
		return fmt.Errorf("failed to fetch project items: %w", err)
	}

	resp, err := demote.Run(ctx, cfg, items, demoter)
	if err != nil {
		return fmt.Errorf("failed to run demote: %w", err)
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal results: %w", err)
	}

	fmt.Println(string(out))
	return nil
}
