package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`
	UseMongo      bool   `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`
	LimitCount    int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTime     int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
}
