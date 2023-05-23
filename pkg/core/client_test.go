package core

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
)

type MockClientWithHeader struct {
	ResponseCode   int
	ResponseHeader http.Header
}

func (c *MockClientWithHeader) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: c.ResponseCode,
		Header:     c.ResponseHeader,
	}, nil
}

func TestHttpRateLimitedClient_Do(t *testing.T) {
	mockClient := &MockClientWithHeader{
		ResponseHeader: map[string][]string{},
	}
	c, _ := NewRateLimitClient(2)
	c.client = mockClient

	mockClient.ResponseHeader.Set(STHeaderRateLimit, "2")
	mockClient.ResponseCode = http.StatusOK

	req := &http.Request{}

	resp, err := c.Do(req)
	assert.Equal(t, err, nil, "should not return error")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "should return 200")

	resp, err = c.Do(req)
	assert.Equal(t, err, nil, "should not return error")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "should able to execute another call")

	resp, err = c.Do(req)
	assert.Equal(t, err, nil, "should not return error")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "should able to execute 3rd call")

	mockClient.ResponseHeader.Set(STHeaderRateLimit, "2")
	mockClient.ResponseHeader.Set(STHeaderRateLimitReset, time.Now().Add(time.Second).UTC().Format(STRateLimitTimeFormat))
	mockClient.ResponseCode = http.StatusTooManyRequests
	resp, err = c.Do(req)
	assert.Equal(t, err, errors.New("too many attempts for hitting rate limit. aborting request"), "should error out because of rate limit")

}
