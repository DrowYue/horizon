package oauth

import (
	"net/http"
	"time"

	herrors "g.hz.netease.com/horizon/core/errors"
	perror "g.hz.netease.com/horizon/pkg/errors"
	"g.hz.netease.com/horizon/pkg/oauth/generate"
	"g.hz.netease.com/horizon/pkg/oauth/manager"
	oauthmodel "g.hz.netease.com/horizon/pkg/oauth/models"
	"g.hz.netease.com/horizon/pkg/util/wlog"
	"golang.org/x/net/context"
)

type AuthorizeReq struct {
	ClientID     string
	Scope        string
	RedirectURL  string
	State        string
	UserIdentity string

	Request *http.Request
}

type AuthorizeCodeResponse struct {
	Code        string
	RedirectURL string
	State       string
}

type AccessTokenReq struct {
	ClientID     string
	ClientSecret string
	Code         string
	RedirectURL  string

	Request *http.Request
}

type AccessTokenResponse struct {
	AccessToken string        `json:"access_token"`
	ExpiresIn   time.Duration `json:"expires_in"`
	Scope       string        `json:"scope"`
	TokenType   string        `json:"token_type"`
}

type Controller interface {
	GenAuthorizeCode(ctx context.Context, req *AuthorizeReq) (*AuthorizeCodeResponse, error)
	GenAccessToken(ctx context.Context, req *AccessTokenReq) (*AccessTokenResponse, error)
}

func NewController(authManager manager.Manager) Controller {
	return &controller{oauthManager: authManager}
}

var _ Controller = &controller{}

type controller struct {
	oauthManager manager.Manager
}

func (c *controller) GenAuthorizeCode(ctx context.Context, req *AuthorizeReq) (*AuthorizeCodeResponse, error) {
	const op = "oauth controller: GenAuthorizeCode"
	defer wlog.Start(ctx, op).StopPrint()

	// 1. TODO: check if the scope is ok (now horizon app do not need provide scope)
	// 2. gen authorization Code
	authToken, err := c.oauthManager.GenAuthorizeCode(ctx, &manager.AuthorizeGenerateRequest{
		ClientID:     req.ClientID,
		RedirectURL:  req.RedirectURL,
		State:        req.State,
		Scope:        req.Scope,
		UserIdentify: req.UserIdentity,
		Request:      req.Request,
	})
	if err != nil {
		return nil, err
	}
	resp := &AuthorizeCodeResponse{
		Code:        authToken.Code,
		RedirectURL: authToken.RedirectURI,
		State:       req.State,
	}
	return resp, nil
}

func (c *controller) GenAccessToken(ctx context.Context, req *AccessTokenReq) (*AccessTokenResponse, error) {
	app, err := c.oauthManager.GetOAuthApp(ctx, req.ClientID)
	if err != nil {
		return nil, err
	}
	var generator generate.AccessTokenCodeGenerate
	switch app.AppType {
	case oauthmodel.HorizonOAuthAPP:
		generator = generate.NewHorizonAppUserToServerAccessGenerate()
	case oauthmodel.DirectOAuthAPP:
		generator = generate.NewOauthAccessGenerate()
	default:
		return nil, perror.Wrapf(herrors.ErrOAuthInternal,
			"appType Not Supported, appType = %d", app.AppType)
	}

	token, err := c.oauthManager.GenAccessToken(ctx, &manager.AccessTokenGenerateRequest{
		ClientID:     req.ClientID,
		ClientSecret: req.ClientSecret,
		Code:         req.Code,
		RedirectURL:  req.RedirectURL,
		Request:      req.Request,
	}, generator)
	if err != nil {
		return nil, err
	}
	return &AccessTokenResponse{
		AccessToken: token.Code,
		ExpiresIn:   token.ExpiresIn,
		Scope:       token.Scope,
		TokenType:   "bearer",
	}, err
}
