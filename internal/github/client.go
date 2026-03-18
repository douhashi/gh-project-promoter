package github

import (
	gh "github.com/google/go-github/v82/github"
)

// Client wraps the GitHub API client.
type Client struct {
	inner *gh.Client //nolint:unused // will be used when API methods are implemented
}
