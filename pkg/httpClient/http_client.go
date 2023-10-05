package httpClient

import (
	"bytes"
	"fmt"
	"github.com/sony/gobreaker"
	"io/ioutil"
	"net/http"
	"time"
)

type clientHttp struct {
	client *gobreaker.CircuitBreaker
}

type IClientHttp interface {
	Get(url string) ([]byte, error)
	Post(url string, body []byte) ([]byte, error)
	Put(url string, body []byte) ([]byte, error)
	Delete(url string) error
}

func NewClient() IClientHttp {
	st := gobreaker.Settings{
		Name:        "HTTPClient",
		MaxRequests: 3,
		Interval:    5 * time.Second,
		Timeout:     10 * time.Second,
	}

	return &clientHttp{
		client: gobreaker.NewCircuitBreaker(st),
	}
}

func (c *clientHttp) Get(url string) ([]byte, error) {
	resp, err := c.client.Execute(func() (interface{}, error) {
		return c.get(url)
	})
	if err != nil {
		return nil, err
	}

	return resp.([]byte), nil
}

func (c *clientHttp) Post(url string, body []byte) ([]byte, error) {
	resp, err := c.client.Execute(func() (interface{}, error) {
		return c.post(url, body)
	})

	if err != nil {
		return nil, err
	}
	return resp.([]byte), nil
}

func (c *clientHttp) Put(url string, body []byte) ([]byte, error) {
	resp, err := c.client.Execute(func() (interface{}, error) {
		return c.put(url, body)
	})

	if err != nil {
		return nil, err
	}

	return resp.([]byte), nil
}

func (c *clientHttp) Delete(url string) error {
	_, err := c.client.Execute(func() (interface{}, error) {
		return nil, c.delete(url)
	})

	return err
}

func (c *clientHttp) get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Received non-200 status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func (c *clientHttp) post(url string, body []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Received non-200 status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func (c *clientHttp) put(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Received non-200 status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func (c *clientHttp) delete(url string) error {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Received non-200 status code: %d", resp.StatusCode)
	}

	return nil
}
