// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"time"

	"gopkg.in/mgo.v2"
)

// Config is the main configuration struct passed to Mongobeat.New() method at startup
type Config struct {
	Period time.Duration `config:"period"`
	// Name of your target
	Name string `config:"name"`
	// DbStats toggles whether to monintor db.stats()
	DbStats bool `config:"db_stats"`
	// ServerStatus toggles whether to monitor db.serverStatus()
	ServerStatus bool `config:"server_status"`
	// ConnectionInfo are the params used to connect to mongo instance
	ConnectionInfo *mgo.DialInfo `config:"connection_info"`
	// DiscoveryNodes specifies whether mongobeat should discovery and report on additional nodes
	// in a cluster
	DiscoverNodes bool `config:"discover_nodes"`
}

// DefaultConfig is the default configuration parameters used to run mongobeat
var DefaultConfig = Config{
	Period:       1 * time.Second,
	Name:         "Mongobeat default local target 1",
	DbStats:      true,
	ServerStatus: true,
	ConnectionInfo: &mgo.DialInfo{
		Addrs: []string{
			"localhost:27017",
		},
	},
	DiscoverNodes: false,
}
