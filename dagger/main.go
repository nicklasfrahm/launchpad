// A generated module for Launchpad functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"

	"dagger/launchpad/pkg/utils"
)

// New creates a new instance of the Launchpad module.
func New(
	// +required
	source *Directory,
) *Launchpad {
	return &Launchpad{
		Source: source,
	}
}

// Launchpad is the entrypoint for the entire pipeline.
type Launchpad struct {
	// Source contains the source code to be processed.
	Source *Directory
}

// SourceMetadata returns metadata about the source code, such as the version and commit hash.
type SourceMetadata struct {
	// Version is the version of the source code.
	Version string `json:"version"`
	// Owner is the organization that owns the source code.
	Owner string `json:"owner"`
	// Repo is the name of the repository.
	Repo string `json:"repo"`
}

// Returns metadata about the source code, such as the version and commit hash.
func (m *Launchpad) SourceMetadata(ctx context.Context) *SourceMetadata {
	metadataContainer := dag.Container().
		From("cgr.dev/chainguard/go").
		WithMountedDirectory("/src", m.Source).
		WithWorkdir("/src").
		WithEntrypoint([]string{"sh", "-c"})

	version, err := metadataContainer.
		WithExec([]string{"git fetch && git describe --tags --always --exclude main"}).
		Stdout(ctx)
	if err != nil {
		return nil
	}

	remoteUrl, err := metadataContainer.
		WithExec([]string{"git remote get-url origin"}).
		Stdout(ctx)
	if err != nil {
		return nil
	}

	owner, repo := utils.ParseGitRemote(remoteUrl)

	return &SourceMetadata{
		Version: version,
		Owner:   owner,
		Repo:    repo,
	}
}
