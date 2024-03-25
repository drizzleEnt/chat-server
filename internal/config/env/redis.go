package env

import (
	"errors"
	"net"
	"os"
)

const (
	rdsName     = "REDIS_DATABASE_NAME"
	rdsuser     = "REDIS_USER"
	rdspassword = "REDIS_PASSWORD"
	rdsport     = "REDIS_PORT"
	rdshost     = "REDIS_HOST"
)

type redisConfig struct {
	adrr       string
	clientName string
	userName   string
	password   string
	db         int
}

func NewRedisConfig() (*redisConfig, error) {
	host := os.Getenv(rdshost)

	if len(host) == 0 {
		return nil, errors.New("redis host not found")
	}

	port := os.Getenv(rdsport)

	if len(port) == 0 {
		return nil, errors.New("redis port not found")
	}

	name := os.Getenv(rdsName)

	if len(name) == 0 {
		return nil, errors.New("redis name not found")
	}

	user := os.Getenv(rdsuser)

	if len(user) == 0 {
		return nil, errors.New("redis user not found")
	}

	password := os.Getenv(rdspassword)

	if len(password) == 0 {
		return nil, errors.New("redis password not found")
	}

	adr := net.JoinHostPort(host, port)

	return &redisConfig{
		adrr:       adr,
		clientName: name,
		userName:   user,
		password:   password,
		db:         0,
	}, nil
}

func (cfg *redisConfig) Address() string {
	return cfg.adrr
}

func (cfg *redisConfig) ClientName() string {
	return cfg.clientName
}

func (cfg *redisConfig) UserName() string {
	return cfg.userName
}

func (cfg *redisConfig) Password() string {
	return cfg.password
}
