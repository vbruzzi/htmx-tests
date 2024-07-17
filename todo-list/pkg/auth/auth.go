package auth

import (
	"encoding/json"
	"net/http"
	"net/url"
	"vbruzzi/todo-list/pkg/config"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Realm       string
	ClientId    string
	Authority   string
	Scheme      string
	RedirectUri string
	Path        string
	CookieName  string
}

type TokenResult struct {
	AccessToken string `json:"access_token"`
}

func (h *AuthHandler) validateAuthorizationCode(code string) (string, error) {
	authUrl := url.URL{
		Scheme: h.Scheme,
		Host:   "identity:8080",
		Path:   h.Path + "/token",
	}

	data := url.Values{}
	data.Add("client_id", h.ClientId)
	data.Add("grant_type", "authorization_code")
	data.Add("code", code)
	data.Add("redirect_uri", h.RedirectUri)

	res, err := http.PostForm(authUrl.String(), data)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	v := TokenResult{}
	err = json.NewDecoder(res.Body).Decode(&v)

	if err != nil {
		return "", err
	}

	return v.AccessToken, nil
}

func (h *AuthHandler) IsAuthenticated(c echo.Context) bool {
	cookie, err := c.Cookie(h.CookieName)
	return err == nil && cookie.Value != ""
}

func (h *AuthHandler) generateAuthCodeUrl() string {
	generationUrl := url.URL{
		Scheme: h.Scheme,
		Host:   h.Authority,
		Path:   h.Path + "/auth",
	}

	query := generationUrl.Query()
	query.Add("response_type", "code")
	query.Add("scope", "openid email")
	query.Add("state", "random")
	query.Add("client_id", h.ClientId)
	query.Add("redirect_uri", h.RedirectUri)

	generationUrl.RawQuery = query.Encode()

	return generationUrl.String()
}

func (h *AuthHandler) Authenticate(c echo.Context) error {
	if code := c.Request().URL.Query().Get("code"); code != "" {
		accessToken, err := h.validateAuthorizationCode(code)

		if err != nil {
			return err
		}

		c.SetCookie(&http.Cookie{
			Name:  h.CookieName,
			Value: accessToken,
		})

		return c.Redirect(http.StatusSeeOther, "/")
	}

	if _, err := c.Cookie(h.CookieName); err != nil {
		redirUrl := h.generateAuthCodeUrl()
		return c.Redirect(http.StatusSeeOther, redirUrl)
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func New(config config.OidcConfig) AuthHandler {
	return AuthHandler{
		Realm:       config.Realm,
		ClientId:    config.ClientId,
		Authority:   config.Authority,
		RedirectUri: config.RedirectUrl,
		Scheme:      config.Scheme,
		CookieName:  config.CookieName,
		Path:        "/realms/" + config.Realm + "/protocol/openid-connect",
	}
}
