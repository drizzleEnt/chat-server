package config

import "github.com/subosito/gotenv"

func Load(path string) error {
	err := gotenv.Load(path)

	if err != nil {
		return err
	}

	return nil
}

type RedisConfig interface {
	Address() string
	ClientName() string
	UserName() string
	Password() string
}

type PgConfig interface {
	DSN() string
}

type GRPCConfig interface {
	Address() string
}
