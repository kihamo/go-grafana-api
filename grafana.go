package go_grafana_api // import "github.com/kihamo/go-grafana-api"

//go:generate /bin/bash -c "goimports -w `find . -type f -name '*.go' -not -path './vendor/*'`"

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

type ErrorOutput struct {
	Message   string `json:"message"`
	RealError string `json:"error"`
}

func (e *ErrorOutput) Error() error {
	if e.RealError != "" {
		return errors.New(e.RealError)
	}

	return errors.New(e.Message)
}

type ErrorsListOutput []struct {
	FieldNames     []string `json:"fieldNames"`
	Classification string   `json:"classification"`
	Message        string   `json:"message"`
}

func (e *ErrorsListOutput) Error() error {
	messages := []string{}

	for _, f := range *e {
		messages = append(messages, strings.Join(f.FieldNames, ", ")+" "+f.Message)
	}

	return errors.New(strings.Join(messages, ". "))
}

type Logger interface {
	Println(v ...interface{})
}

type Client struct {
	client        *http.Client
	logger        Logger
	address       string
	authorization string
}

func New(address string) *Client {
	address = strings.TrimRight(address, "/") + "/api/"

	return &Client{
		client:  &http.Client{},
		address: address,
	}
}

func (c *Client) WithLogger(logger Logger) *Client {
	c.logger = logger
	return c
}

func (c *Client) WithBasicAuth(username, password string) *Client {
	key := username + ":" + password
	key = base64.StdEncoding.EncodeToString([]byte(key))

	c.authorization = "Basic " + key

	return c
}

func (c *Client) WithApiKey(key string) *Client {
	c.authorization = "Bearer " + key

	return c
}

func (c *Client) newRequest(ctx context.Context, method, path string, input interface{}) (*http.Request, error) {
	u, err := url.Parse(c.address + strings.TrimLeft(path, "/"))
	if err != nil {
		return nil, err
	}

	var body []byte
	if input != nil {
		if method == http.MethodPost || method == http.MethodPut {
			if body, err = json.Marshal(input); err != nil {
				return nil, err
			}
		} else {
			values, err := query.Values(input)
			if err != nil {
				return nil, err
			}

			u.RawQuery = values.Encode()
		}
	}

	request, err := http.NewRequest(method, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		request = request.WithContext(ctx)
	}

	request.Header.Set("User-Agent", "Go-grafana-api-client/1.0")
	request.Header.Set("Content-Type", "application/json")

	if c.authorization != "" {
		request.Header.Set("Authorization", c.authorization)
	}

	return request, nil
}

func (c *Client) send(request *http.Request, output interface{}) error {
	if c.logger != nil {
		dumpRequest, _ := httputil.DumpRequestOut(request, true)
		c.logger.Println("--- REQUEST ---\n", string(dumpRequest), "\n--- /REQUEST ---")
	}

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if c.logger != nil {
		dumpResponse, _ := httputil.DumpResponse(response, true)
		c.logger.Println("--- RESPONSE ---\n", string(dumpResponse), "\n--- /RESPONSE ---")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusUnprocessableEntity {
		errorOut := &ErrorsListOutput{}
		if err := json.Unmarshal(body, errorOut); err == nil {
			return errorOut.Error()
		}
	} else if response.StatusCode >= 400 {
		errorOut := &ErrorOutput{}
		if err := json.Unmarshal(body, errorOut); err == nil {
			return errorOut.Error()
		}
	}

	return json.Unmarshal(body, output)
}

func Int64(v int64) *int64 {
	return &v
}

func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}

	return 0
}

func String(v string) *string {
	return &v
}

func StringSlice(src []string) []*string {
	dst := make([]*string, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

func TimeValue(v *time.Time) time.Time {
	if v != nil {
		return *v
	}

	return time.Time{}
}

func MillisecondsTimeValue(v *int64) time.Time {
	if v != nil {
		return time.Unix(0, (*v * 1000000))
	}
	return time.Time{}
}
