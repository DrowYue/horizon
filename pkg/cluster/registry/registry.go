package registry

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	harbormodels "g.hz.netease.com/horizon/pkg/harbor/models"

	"github.com/hashicorp/go-retryablehttp"
)

// Registry ...
type Registry interface {
	// CreateProject create a project, if the project is already exists, return true, or return project's ID
	CreateProject(ctx context.Context, project string) (bool, int, error)
	// AddMembers add members for project
	AddMembers(ctx context.Context, projectID int) error
	// DeleteRepository delete repository
	DeleteRepository(ctx context.Context, project string, repository string) error
	// ListImage list images for a repository
	ListImage(ctx context.Context, project string, repository string) ([]string, error)
	GetServer(ctx context.Context) string
}

type HarborMember struct {
	// harbor role 1:manager，2:developer，3:guest
	Role int `yaml:"role"`
	// harbor user name
	Username string `yaml:"username"`
}

// HarborRegistry implement Registry
type HarborRegistry struct {
	// harbor server address
	server string
	// harbor token
	token string
	// the member to add to projects
	members []*HarborMember
	// http client
	client *http.Client
	// retryableClient retryable client
	retryableClient *retryablehttp.Client
}

// default params
const (
	_backoffDuration = 1 * time.Second
	_retry           = 3
	_timeout         = 4 * time.Second
)

// harbor member to add to harbor project
var members = []*HarborMember{
	{
		Role:     3,
		Username: "musiccloudnative",
	},
}

// NewHarborRegistry new a HarborRegistry
func NewHarborRegistry(harbor *harbormodels.Harbor) Registry {
	transport := http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return &HarborRegistry{
		server:  harbor.Server,
		token:   harbor.Token,
		members: members,
		client: &http.Client{
			Transport: &transport,
		},
		retryableClient: &retryablehttp.Client{
			HTTPClient: &http.Client{
				Transport: &transport,
				Timeout:   _timeout,
			},
			RetryMax:   _retry,
			CheckRetry: retryablehttp.DefaultRetryPolicy,
			Backoff: func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
				// wait for this duration if failed
				return _backoffDuration
			},
		},
	}
}

var _ Registry = (*HarborRegistry)(nil)
