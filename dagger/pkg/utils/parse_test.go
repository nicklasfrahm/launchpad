package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRemoteUrl(t *testing.T) {
	t.Run("should parse GitHub URL", func(t *testing.T) {
		// Act.
		owner, repo := ParseGitRemote("https://github.com/OWNER/REPO.git\n")

		// Assert.
		assert.Equal(t, "OWNER", owner)
		assert.Equal(t, "REPO", repo)
	})

	t.Run("should parse GitHub URL without .git suffix", func(t *testing.T) {
		// Act.
		owner, repo := ParseGitRemote("https://github.com/OWNER/REPO")

		// Assert.
		assert.Equal(t, "OWNER", owner)
		assert.Equal(t, "REPO", repo)
	})

	t.Run("should parse GitLab URL", func(t *testing.T) {
		// Act.
		owner, repo := ParseGitRemote("https://gitlab.com/OWNER/REPO")

		// Assert.
		assert.Equal(t, "OWNER", owner)
		assert.Equal(t, "REPO", repo)
	})

	t.Run("should ignore Azure DevOps URL", func(t *testing.T) {
		// Act.
		owner, repo := ParseGitRemote("https://dev.azure.com/ORG/PROJECT/_git/REPO")

		// Assert.
		assert.Equal(t, "", owner)
		assert.Equal(t, "", repo)
	})

	t.Run("should ignore Visual Studio URL", func(t *testing.T) {
		// Act.
		owner, repo := ParseGitRemote("https://ORG.visualstudio.com/DefaultCollection/PROJECT/_git/REPO")

		// Assert.
		assert.Equal(t, "", owner)
		assert.Equal(t, "", repo)
	})
}
