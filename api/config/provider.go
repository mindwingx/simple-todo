package config

import "time"

type GRPC struct {
	Host    string        `mapstructure:"GRPC_CLIENT_HOST"`
	Port    string        `mapstructure:"GRPC_CLIENT_PORT"`
	Time    time.Duration `mapstructure:"GRPC_PING_TIME"`
	Timeout time.Duration `mapstructure:"GRPC_PING_TIMEOUT"`
}
