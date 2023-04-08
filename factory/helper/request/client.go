package request

import (
	"io"
	"net/http"
	"strings"
	"time"

	"tdp-aiart/helper/logman"
)

type Client struct {
	Method  string
	Url     string
	Data    string
	Headers map[string]string
	Timeout time.Duration
}

func (c *Client) Request() ([]byte, error) {

	var (
		err  error
		body io.Reader
		req  *http.Request
		resp *http.Response
	)

	if c.Data != "" {
		body = strings.NewReader(c.Data)
	}

	// 创建请求
	if req, err = http.NewRequest(c.Method, c.Url, body); err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	logman.Info("Http request", "method", c.Method, "url", c.Url)

	// 发起请求
	client := http.Client{Timeout: c.Timeout}
	if resp, err = client.Do(req); err != nil {
		return nil, err
	}

	// 读取数据
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)

}

func (c *Client) JsonRequest() ([]byte, error) {

	if c.isBodyMethod() && c.Headers["Content-Type"] == "" {
		c.Headers["Content-Type"] = "application/json"
	}

	return c.Request()

}

func (c *Client) TextRequest() (string, error) {

	if c.isBodyMethod() && c.Headers["Content-Type"] == "" {
		c.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	body, err := c.Request()
	return string(body), err

}

func (c *Client) isBodyMethod() bool {

	return strings.Contains("POST,PUT,PATCH", c.Method)

}
