package database

import "context"

type Config struct {
	Driver   string
	Source   string
	Host     string
	Password string
	Db       string
}

func Init(ctx context.Context, config Config) interface{} {
	switch config.Driver {
	case "mysql":
		return InitMysql(config.Driver, config.Source)
	case "redis":
		return InitRedis(ctx, config.Host, config.Password, config.Db)
	default:
		return nil
	}
}
