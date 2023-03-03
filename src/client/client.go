package client

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

type WebhookAuthData struct {
	UserID int    `valid:"required"`
	Secret string `valid:"alphanum,required"`
}

type Client struct {
	webhookAuth *WebhookAuthData
	Url         *url.URL
	httpClient  *http.Client
}

func NewClientWithWebhookAuth(intranetUrl string, userId int, secret string) (*Client, error) {
	u, err := url.Parse(fmt.Sprintf("%s/rest/%d/%s/", intranetUrl, userId, secret))
	if err != nil {
		return nil, errors.Wrap(err, "Error parsing B24 URL")
	}

	auth := &WebhookAuthData{
		UserID: userId,
		Secret: secret,
	}

	_, err = govalidator.ValidateStruct(auth)
	if err != nil {
		return nil, errors.Wrap(err, "Auth params validation failed")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}

	return &Client{
		Url:         u,
		webhookAuth: auth,
		httpClient:  httpClient,
	}, nil
}
