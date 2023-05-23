package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nandiheath/spacetraders/pkg/api"
	"github.com/nandiheath/spacetraders/pkg/log"
)

const APIAddress = "https://api.spacetraders.io/v2"
const STHeaderRateLimit = "x-ratelimit-limit"
const STHeaderRateLimitBurst = "x-ratelimit-limit-burst"
const STHeaderRateLimitRemaining = "x-ratelimit-remaining"
const STHeaderRateLimitReset = "x-ratelimit-reset"
const STHeaderRateLimitPerSecond = "x-ratelimit-limit-per-second"

const ClientParallelLimit = 2
const ClientTokenTimeout = 10 * time.Second

const STRateLimitTimeFormat = "2006-01-02T15:04:05.000Z"

type HttpRateLimitedClient struct {
	client       api.HttpRequestDoer
	tokenChannel chan bool
	tokenTimeout time.Duration
}

// NewRateLimitClient returns HttpRateLimitedClient with parallel limit
func NewRateLimitClient(parallel int) (*HttpRateLimitedClient, error) {
	tokenC := make(chan bool, parallel)
	// prefill the tokenChannel
	for i := 0; i < parallel; i++ {
		tokenC <- true
	}

	return &HttpRateLimitedClient{
		client:       &http.Client{},
		tokenChannel: tokenC,
		tokenTimeout: ClientTokenTimeout,
	}, nil
}

func (c *HttpRateLimitedClient) Do(req *http.Request) (*http.Response, error) {
	select {
	case <-c.tokenChannel:
		attempt := 0
		defer func() {
			c.tokenChannel <- true
		}()
		for attempt < 3 {
			resp, err := c.client.Do(req)
			if err != nil {
				return nil, err
			}
			// catch only the rate limit
			if resp.StatusCode == http.StatusTooManyRequests {
				attempt++
				timeStr := resp.Header.Get(STHeaderRateLimitReset)
				t, err := time.Parse(STRateLimitTimeFormat, timeStr)
				if err != nil {
					log.Logger().Errorf("unable to parse time for rate limit. time:%s", timeStr)
					continue
				}
				// sleep until rate limit is reset
				d := t.Sub(time.Now())
				log.Logger().Debugf("hitting rate limit. paused for %+v.", d)
				time.Sleep(d)
				continue
			}
			return resp, err
		}
		return nil, errors.New("too many attempts for hitting rate limit. aborting request")
	case <-time.After(c.tokenTimeout):
		return nil, errors.New("unable to obtain token to request")
	}
	return nil, nil
}

// NewAPIClient returns the APIClient with bearer token and handles rate limiting
func NewAPIClient() *api.ClientWithResponses {
	c, _ := api.NewClientWithResponses(APIAddress, func(client *api.Client) error {
		client.RequestEditors = append(client.RequestEditors, func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("ACCESS_TOKEN")))
			return nil
		})

		return nil
	}, func(client *api.Client) error {
		c, err := NewRateLimitClient(ClientParallelLimit)
		if err != nil {
			return err
		}
		client.Client = c
		return nil
	})
	return c
}

type APIResponse interface {
	Status() string
	StatusCode() int
	ResponseBody() []byte
}

type ResponseError struct {
	DataError struct {
		Message string
		Code    int
	} `json:"error"`
}

func (e *ResponseError) Error() string {
	return e.DataError.Message
}

func newResponseError(err error) *ResponseError {
	if err == nil {
		return nil
	}
	return &ResponseError{DataError: struct {
		Message string
		Code    int
	}{Message: err.Error(), Code: 0}}
}

func TryParseError(resp APIResponse, err error) *ResponseError {
	if resp.StatusCode() != 200 && resp.StatusCode() != 201 {
		errorResp := ResponseError{}
		e := json.Unmarshal(resp.ResponseBody(), &errorResp)
		if e != nil {
			return newResponseError(e)
		}
		return &errorResp
	}
	return newResponseError(err)
}
