package gnet

import (
	"errors"
	"github.com/golyu/vlog"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func InitDefaultEnv(c *cli.Context)error{
	env := c.String("env")
	if env == "" {
		return errors.New("config cannot empty")
	}
	var configPath string
	if env == "test" {
		configPath = "./config/config_test.toml"
	} else if env == "pro" {
		configPath = "./config/config_pro.toml"
	}

	viper.SetConfigType("toml")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(errors.New("config file not found"))
		}
	}
	return nil
}

func InitLog() *logrus.Logger{
	logDir := viper.GetString("log.dir")
	logMode := viper.GetString("log.level")
	day := viper.GetInt("log.day")
	intervalHour := viper.GetInt("log.intervalHour")
	log, err := vlog.Init(logDir, logMode, day, intervalHour)
	if err != nil {
		panic(err)
	}
	return log
}