package http

type Config struct {
	Port int `yaml:"port"`
}

const (
	DefaultHTTPPort = 8000
)

func (c *Config) setDefaultConfig() {
	if c.Port <= 0 {
		c.Port = DefaultHTTPPort
	}
}
