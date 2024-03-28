package utils

import (
	"net/url"
	"strings"
)

// ParseGitRemote parses a remote URL and returns the owner and repository name.
func ParseGitRemote(remoteUrl string) (string, string) {
	u, err := url.Parse(strings.TrimSpace(remoteUrl))
	if err != nil {
		return "", ""
	}

	cleanPath := strings.TrimSuffix(u.Path, ".git")

	// TODO: Support other git hosting services,
	// such as Azure DevOps. How do we map orgs
	// and projects to owners and repos?
	parts := strings.Split(cleanPath, "/")
	if len(parts) != 3 {
		return "", ""
	}

	owner := parts[1]
	repo := parts[2]

	return owner, repo
}
