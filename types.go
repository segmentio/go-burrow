package burrow

// Error reprsents a burrow api error.
type Error struct {
	Message string `json:"message"`
}

// Error implements the error interface.
func (e *Error) Error() string {
	return e.Message
}

// Clusters represents a clusters response.
type Clusters struct {
	Names []string `json:"clusters"`
}

// Cluster represents a cluster response.
type Cluster struct {
	Zookeepers    []string `json:"zookeepers"`
	ZookeeperPort int      `json:"zookeeper_port"`
	ZookeeperPath string   `json:"zookeeper_path"`
	Brokers       []string `json:"brokers"`
	BrokerPort    int      `json:"broker_port"`
	OffsetsTopic  string   `json:"offsets_topic"`
}

// Consumers represents a list of consumers.
type Consumers struct {
	Names []string `json:"consumers"`
}

// Topics represents a list of topics.
type Topics struct {
	Names []string `json:"topics"`
}

// Topic represents a topic.
type Topic struct {
	Offsets []int `json:"offsets"`
}

// Lag represents a lag.
type Lag struct {
	Cluster    string      `json:"cluster"`
	Group      string      `json:"group"`
	Status     string      `json:"status"`
	Complete   bool        `json:"complete"`
	Partitions []Partition `json:"partitions"`
	MaxLag     Partition   `json:"maxlag"`
	TotalLag   int         `json:"totallag"`
}

// Partition represents a partition lag info.
type Partition struct {
	Topic     string `json:"topic"`
	Partition int    `json:"partition"`
	Status    string `json:"status"`
	Start     struct {
		Offset    int `json:"offset"`
		Timestamp int `json:"timestamp"`
		Lag       int `json:"lag"`
	} `json:"start"`
	End struct {
		Offset    int `json:"offset"`
		Timestamp int `json:"timestamp"`
		Lag       int `json:"lag"`
	} `json:"end"`
}
