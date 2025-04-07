package a10n

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	headerContentType    = "Content-Type"
	headerFormUrlEncoded = "application/x-www-form-urlencoded"

	errAuthorizationPending = "authorization_pending"
	errSlowDown             = "slow_down"
)

type (
	deviceCodeResponse struct {
		DeviceCode              string `json:"device_code"`
		UserCode                string `json:"user_code"`
		VerificationURI         string `json:"verification_uri"`
		VerificationURIComplete string `json:"verification_uri_complete"`
		ExpiresIn               int    `json:"expires_in"`
		Interval                int    `json:"interval"`
	}

	TokenResponse struct {
		AccessToken string `json:"access_token"`
		IDToken     string `json:"id_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
		Error       string `json:"error"`
	}

	settings struct {
		AuthDevEndpoint string
		TokenEndpoint   string
		ClientID        string
		PollInterval    time.Duration
		Client          *http.Client
		GrantType       string
	}

	Setting func(*settings)

	Device struct {
		settings settings
	}
)

func newSettings(opt ...Setting) settings {
	s := settings{
		GrantType:    "device_code",
		PollInterval: time.Second,
		Client:       &http.Client{},
	}

	for _, o := range opt {
		o(&s)
	}

	return s
}

func WithKeycloak(baseURL, realm string) Setting {
	return func(s *settings) {
		s.AuthDevEndpoint = fmt.Sprintf("%s/realms/%s/protocol/openid-connect/auth/device", baseURL, realm)
		s.TokenEndpoint = fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", baseURL, realm)
		s.GrantType = "urn:ietf:params:oauth:grant-type:device_code"
	}
}

func WithAuthDeviceEndpoint(v string) Setting {
	return func(s *settings) {
		s.AuthDevEndpoint = v
	}
}

func WithTokenEndpoint(v string) Setting {
	return func(s *settings) {
		s.TokenEndpoint = v
	}
}

func WithClientID(v string) Setting {
	return func(s *settings) {
		s.ClientID = v
	}
}

func WithPollInterval(v time.Duration) Setting {
	return func(s *settings) {
		s.PollInterval = v
	}
}

func WithClient(v *http.Client) Setting {
	return func(s *settings) {
		s.Client = v
	}
}

func NewDevice(opt ...Setting) Device {
	return Device{
		settings: newSettings(opt...),
	}
}

func (d *Device) Auth(ctx context.Context, out io.Writer) (TokenResponse, error) {
	dev, err := d.requestDeviceCode(ctx)
	if err != nil {
		return TokenResponse{}, err
	}

	_, _ = out.Write([]byte(fmt.Sprintf("Login required! Statring device authentication flow...\n")))
	_, _ = out.Write([]byte(fmt.Sprintf("%s\n", dev.VerificationURIComplete)))

	interval := d.settings.PollInterval
	for {
		token, err := d.requestToken(ctx, dev.DeviceCode)
		if err != nil {
			return TokenResponse{}, err
		}

		switch token.Error {
		case "":
			return token, nil
		case errAuthorizationPending:
			break
		case errSlowDown:
			interval += d.settings.PollInterval
			break
		default:
			return TokenResponse{}, fmt.Errorf("unknown error: %s", token.Error)
		}

		select {
		case <-ctx.Done():
			return TokenResponse{}, ctx.Err()

		case <-time.After(interval):
			break
		}
	}
}

func (d *Device) requestToken(ctx context.Context, deviceCode string) (TokenResponse, error) {
	form := url.Values{}
	form.Set("grant_type", d.settings.GrantType)
	form.Set("device_code", deviceCode)
	form.Set("client_id", d.settings.ClientID)

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		d.settings.TokenEndpoint,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return TokenResponse{}, err
	}

	req.Header.Set(headerContentType, headerFormUrlEncoded)

	resp, err := d.settings.Client.Do(req)
	if err != nil {
		return TokenResponse{}, fmt.Errorf("token error: %v", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var token TokenResponse
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &token); err != nil {
		return TokenResponse{}, err
	}

	return token, nil
}

func (d *Device) requestDeviceCode(ctx context.Context) (deviceCodeResponse, error) {
	form := url.Values{}
	form.Set("client_id", d.settings.ClientID)

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		d.settings.AuthDevEndpoint,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return deviceCodeResponse{}, err
	}

	req.Header.Set(headerContentType, headerFormUrlEncoded)

	resp, err := d.settings.Client.Do(req)
	if err != nil {
		return deviceCodeResponse{}, fmt.Errorf("device code request error: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return deviceCodeResponse{}, fmt.Errorf(
			"device code response error - status: %d, body: %s",
			resp.StatusCode,
			body,
		)
	}

	var res deviceCodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return deviceCodeResponse{}, fmt.Errorf("device code error: %w", err)
	}
	return res, nil
}
