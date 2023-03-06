package BitrixGo

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

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

func (c *Client) AddTask(task Task) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.Url.String()+"tasks.task.add.json?", nil)
	if err != nil {
		log.Println(err)
	}
	AddParamsFromStruct(req, task)
	reqDump, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s", string(reqDump))
	return c.httpClient.Do(req)
}

func (c *Client) GetTask(taskid int) (string, error) {
	req, err := http.NewRequest("POST", c.Url.String()+"tasks.task.get.json?", nil)
	if err != nil {
		log.Println(err)
	}
	AddParam(req, "taskId", strconv.Itoa(taskid))
	reqDump, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s", string(reqDump))

	response, err := c.httpClient.Do(req)

	answer, _ := ioutil.ReadAll(response.Body)
	return string(answer), err
}

func (c *Client) UpdateTask(taskid int, task Task) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.Url.String()+"tasks.task.update.json?", nil)
	if err != nil {
		log.Println(err)
	}
	AddParamsFromStruct(req, task)
	AddParam(req, "taskId", strconv.Itoa(taskid))
	reqDump, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s", string(reqDump))
	return c.httpClient.Do(req)
}

func (c *Client) CheckTask(taskid int, task Task) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.Url.String()+"tasks.task.update.json?", nil)
	if err != nil {
		log.Println(err)
	}
	AddParamsFromStruct(req, task)
	AddParam(req, "taskId", strconv.Itoa(taskid))
	reqDump, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s", string(reqDump))
	return c.httpClient.Do(req)
}
