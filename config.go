package nametagprinter

const VERSION = "v0.0.3"

type Config struct {
	Server struct {
		Address string
		Port    int
	}
	Tag struct {
		Template string
	}
	Printer string
}

func NewConfig() (c *Config) {
	c = new(Config)
	c.Server.Port = 8088
	c.Server.Address = "0.0.0.0"
	c.Tag.Template = "./upload/nametag.svg"
	c.Printer = "Brother_QL-720NW_USB"
	return
}
