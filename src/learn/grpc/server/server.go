// main
package main

import (
	"flag"

	"ppzzl.com/tinyblog-go/learn/grpc/hello"
)

// Config 服务端配置。
type Config struct {
	host string
	port int
}

func (cfg *Config) initialize() {
	cfg.port = *flag.Int("p", 20801, "listening port")
	cfg.host = *flag.String("b", "localhost", "binding host")
	flag.Parse()
}

func main() {
	cfg := Config{}
	cfg.initialize()
	s := hello.NewServer(cfg.host, cfg.port)
	s.Run()
}
