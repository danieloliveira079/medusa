package git

import (
	"github.com/pkg/errors"
)

const (
	RepoNoConfig GitRepoStatus = "unconfigured" // configuration is empty
	RepoNew                    = "new"          // no attempt made to clone it yet
	RepoCloned                 = "cloned"       // has been read (cloned); no attempt made to write
	RepoReady                  = "ready"        // has been written to, so ready to sync
)

type Repo struct {
	GitRemoteConfig
}

type GitRemoteConfig struct {
	URL    string `json:"url"`
	Branch string `json:"branch"`
}

type GitConfig struct {
	Remote GitRemoteConfig `json:"remote"`
	Status GitRepoStatus   `json:"status"`
}

type GitRepoStatus string

func NewGitRemoteConfig(url, branch string) (GitRemoteConfig, error) {
	if len(url) == 0 {
		return GitRemoteConfig{}, errors.New("The remote URL is empty!")
	}
	return GitRemoteConfig{
		URL:    url,
		Branch: branch,
	}, nil
}
