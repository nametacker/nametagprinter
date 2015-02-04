package nametagprinter

const VERSION = "v0.0.1"

type Config struct {
	Server struct {
		Address string
		Port    int
	}
}

func NewConfig() (c *Config) {
	c = new(Config)
	c.Server.Port = 8088
	c.Server.Address = "0.0.0.0"
	return
}
