// Package burrow implements a simple burrow client.
package burrow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Client is a simple burrow client.
type Client struct {
	url string
}

// New returns a new *Client with the given endpoint.
func New(endpoint string) (*Client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	return &Client{
		url: fmt.Sprintf("%s://%s/v3", u.Scheme, u.Host),
	}, nil
}

// Clusters returns a list of clusters.
func (c *Client) Clusters() (clusters Clusters, err error) {
	err = c.get("/kafka", &clusters)
	return
}

// Cluster returns cluster info by name.
func (c *Client) Cluster(name string) (cluster Cluster, err error) {
	err = c.get("/kafka/"+name, &cluster)
	return
}

// Consumers returns a list of consumers for a cluster `name`.
func (c *Client) Consumers(name string) (consumers Consumers, err error) {
	err = c.get("/kafka/"+name+"/consumer", &consumers)
	return
}

// ConsumerTopics returns a list of topics for a `cluster` and `group`.
func (c *Client) ConsumerTopics(cluster, group string) (topics Topics, err error) {
	err = c.get("/kafka/"+cluster+"/consumer/"+group+"/topic", &topics)
	return
}

// ConsumerTopic returns consumer topic info.
func (c *Client) ConsumerTopic(cluster, group, topic string) (t Topic, err error) {
	err = c.get("/kafka/"+cluster+"/consumer/"+group+"/topic/"+topic, &t)
	return
}

// ConsumerLag returns consumer lag.
func (c *Client) ConsumerLag(cluster, group string) (lag Lag, err error) {
	res := struct {
		Lag Lag `json:"status"`
	}{}

	err = c.get("/kafka/"+cluster+"/consumer/"+group+"/lag", &res)
	if err != nil {
		return
	}

	return res.Lag, nil
}

// Topics returns a list of topics for a `cluster`.
func (c *Client) Topics(cluster string) (topics Topics, err error) {
	err = c.get("/kafka/"+cluster+"/topic", &topics)
	return
}

// Topic returns info for a `topic` in `cluster`.
func (c *Client) Topic(cluster, topic string) (t Topic, err error) {
	err = c.get("/kafka/"+cluster+"/topic/"+topic, &t)
	return
}

// Get will request and unmarshal the response of `path`
// into the given `v`, in case of an API error `.Error` is returned.
func (c *Client) get(path string, v interface{}) error {
	url := c.url + path

	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrapf(err, "GET %s", url)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var dec = json.NewDecoder(resp.Body)
		var e *Error

		if err = dec.Decode(&e); err == nil {
			return e
		}

		return fmt.Errorf("burrow: %s %s", url, resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		return errors.Wrapf(err, "GET %s", url)
	}

	return nil
}
