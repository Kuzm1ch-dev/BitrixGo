package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/Kuzm1ch-dev/BitrixGo/types"
)

func (c *Client) AddTask(task types.Task) (*http.Response, error) {
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

func (c *Client) UpdateTask(taskid int, task types.Task) (*http.Response, error) {
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

func (c *Client) CheckTask(taskid int, task types.Task) (*http.Response, error) {
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
