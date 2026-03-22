package urlutil

import (
	"fmt"
	"net/url"
	"strings"
)

// ExtractKey builds a key string "{phase}-{owner}-{repo}-{number}" from a GitHub URL and phase name.
func ExtractKey(rawURL string, phase string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	parts := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")
	if len(parts) < 4 {
		return ""
	}
	return fmt.Sprintf("%s-%s-%s-%s", phase, parts[0], parts[1], parts[3])
}

// ExtractRepo extracts "owner/repo" from a GitHub issue/PR URL.
func ExtractRepo(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	parts := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")
	if len(parts) < 2 {
		return ""
	}
	return parts[0] + "/" + parts[1]
}
