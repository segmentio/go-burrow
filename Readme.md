# burrow
--
    import "github.com/segmentio/go-burrow"

Package burrow implements a simple burrow client.

## Usage

#### type Client

```go
type Client struct {
}
```

Client is a simple burrow client.

#### func  New

```go
func New(endpoint string) (*Client, error)
```
New returns a new *Client with the given endpoint.

#### func (*Client) Cluster

```go
func (c *Client) Cluster(name string) (cluster Cluster, err error)
```
Cluster returns cluster info by name.

#### func (*Client) Clusters

```go
func (c *Client) Clusters() (clusters Clusters, err error)
```
Clusters returns a list of clusters.

#### func (*Client) ConsumerLag

```go
func (c *Client) ConsumerLag(cluster, group string) (lag Lag, err error)
```
ConsumerLag returns consumer lag.

#### func (*Client) ConsumerTopic

```go
func (c *Client) ConsumerTopic(cluster, group, topic string) (t Topic, err error)
```
ConsumerTopic returns consumer topic info.

#### func (*Client) ConsumerTopics

```go
func (c *Client) ConsumerTopics(cluster, group string) (topics Topics, err error)
```
ConsumerTopics returns a list of topics for a `cluster` and `group`.

#### func (*Client) Consumers

```go
func (c *Client) Consumers(name string) (consumers Consumers, err error)
```
Consumers returns a list of consumers for a cluster `name`.

#### func (*Client) Topic

```go
func (c *Client) Topic(cluster, topic string) (t Topic, err error)
```
Topic returns info for a `topic` in `cluster`.

#### func (*Client) Topics

```go
func (c *Client) Topics(cluster string) (topics Topics, err error)
```
Topics returns a list of topics for a `cluster`.

#### type Cluster

```go
type Cluster struct {
	Zookeepers    []string `json:"zookeepers"`
	ZookeeperPort int      `json:"zookeeper_port"`
	ZookeeperPath string   `json:"zookeeper_path"`
	Brokers       []string `json:"brokers"`
	BrokerPort    int      `json:"broker_port"`
	OffsetsTopic  string   `json:"offsets_topic"`
}
```

Cluster represents a cluster response.

#### type Clusters

```go
type Clusters struct {
	Names []string `json:"clusters"`
}
```

Clusters represents a clusters response.

#### type Consumers

```go
type Consumers struct {
	Names []string `json:"consumers"`
}
```

Consumers represents a list of consumers.

#### type Error

```go
type Error struct {
	Message string `json:"message"`
}
```

Error reprsents a burrow api error.

#### func (*Error) Error

```go
func (e *Error) Error() string
```
Error implements the error interface.

#### type Lag

```go
type Lag struct {
	Cluster    string      `json:"cluster"`
	Group      string      `json:"group"`
	Status     string      `json:"status"`
	Complete   bool        `json:"complete"`
	Partitions []Partition `json:"partitions"`
}
```

Lag represents a lag.

#### type Partition

```go
type Partition struct {
	Topic     string `json:"topic"`
	Partition int    `json:"partition"`
	Status    string `json:"status"`
	Start     struct {
		Offset    int `json:"offset"`
		Timestamp int `json:"timestamp"`
		Lag       int `json:"lag"`
	}
	End struct {
		Offset    int `json:"offset"`
		Timestamp int `json:"timestamp"`
		Lag       int `json:"lag"`
	}
}
```

Partition represents a partition lag info.

#### type Topic

```go
type Topic struct {
	Offsets []int `json:"offsets"`
}
```

Topic represents a topic.

#### type Topics

```go
type Topics struct {
	Names []string `json:"topics"`
}
```

Topics represents a list of topics.
