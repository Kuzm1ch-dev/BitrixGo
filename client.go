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
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type WebhookAuthData struct {
	UserID int    `valid:"required"`
	Secret string `valid:"alphanum,required"`
}

type Client struct {
	webhookAuth  *WebhookAuthData
	Url          *url.URL
	HttpClient   *http.Client
	HttpServer   *gin.Engine
	OnTaskCreate func(*gin.Context)
	OnTaskDelete func(*gin.Context)
	OnTaskEdit   func(*gin.Context)
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
		HttpClient:  httpClient,
		HttpServer:  gin.Default(),
	}, nil
}

//

func (c *Client) Run(host string, port string) {
	c.HttpServer.POST("/TaskCreate", c.OnTaskCreate)
	c.HttpServer.GET("/TaskCreate", c.OnTaskCreate)
	c.HttpServer.POST("/TaskDelete", c.OnTaskCreate)
	c.HttpServer.GET("/TaskDelete", c.OnTaskCreate)
	c.HttpServer.POST("/TaskEdit", c.OnTaskEdit)
	c.HttpServer.GET("/TaskEdit", c.OnTaskEdit)
	c.HttpServer.Run(fmt.Sprintf("%s:%s", host, port))
}

func (c *Client) AddTask(task Task) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.Url.String()+"tasks.task.add.json?", nil)
	if err != nil {
		log.Println(err)
	}
	AddParamsFromStruct(req, task)
	reqDump, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s", string(reqDump))
	return c.HttpClient.Do(req)
}

func (c *Client) GetTask(taskid int) (string, error) {
	req, err := http.NewRequest("POST", c.Url.String()+"tasks.task.get.json?", nil)
	if err != nil {
		log.Println(err)
	}
	AddParam(req, "taskId", strconv.Itoa(taskid))
	reqDump, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s", string(reqDump))

	response, err := c.HttpClient.Do(req)

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
	return c.HttpClient.Do(req)
}

func (c *Client) CheckTask(taskid int, task Task) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.Url.String()+"tasks.task.get.json?", nil)
	if err != nil {
		log.Println(err)
	}
	AddParamsFromStruct(req, task)
	AddParam(req, "taskId", strconv.Itoa(taskid))
	reqDump, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("REQUEST:\n%s", string(reqDump))
	return c.HttpClient.Do(req)
}
