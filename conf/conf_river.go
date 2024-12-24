package conf

import (
	"blogv2/service/river_service/rule"
)

type River struct {
	Enable   bool          `yaml:"enable"`
	ServerID uint32        `yaml:"server_id"`
	Flavor   string        `yaml:"flavor"`
	DataDir  string        `yaml:"data_dir"`
	Sources  []RiverSource `yaml:"source"`
	Rules    []*rule.Rule  `yaml:"rule"`
	BulkSize int           `yaml:"bulk_size"`
}

type RiverSource struct {
	Schema string   `yaml:"schema"`
	Tables []string `yaml:"tables"`
}
