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

// Returns the version that has been detected in the provided source code.
func (m *Launchpad) SourceVersion(ctx context.Context) string {
	stdout, err := dag.Container().
		From("cgr.dev/chainguard/go").
		WithMountedDirectory("/src", m.Source).
		WithWorkdir("/src").
		WithEntrypoint([]string{"sh", "-c"}).
		WithExec([]string{"git fetch && git describe --tags --always --exclude main"}).
		Stdout(ctx)
	if err != nil {
		return err.Error()
	}

	return stdout
}
