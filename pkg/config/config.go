package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Bind string `yaml:"bind"`

	Db *Db `yaml:"db"`

	DSN string `yaml:"-"`

	Qiniu *Qiniu `yaml:"qiniu"`
}

type Db struct {
	Protocol string `yaml:"protocol"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

const hAO_DB_PASSWORD = "HAO_DB_PASSWORD"

type Qiniu struct {
	Ak       string `yaml:"ak"`
	Sk       string `yaml:"sk"`
	Domain   string `yaml:"domain"`
	Bucket   string `yaml:"bucket"`
	Region   string `yaml:"region"`
	Endpoint string `yaml:"endpoint"`
}

// NewConfig 初始化 yaml 格式的配置文件
func NewConfig(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("fail to read config file: %v, err: %v", path, err)
	}

	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("fail to parse yaml file: %v, err: %v", string(data), err)
	}

	constructDSN(&cfg)

	return &cfg
}

func constructDSN(cfg *Config) {
	var password string
	passwordEnv := os.Getenv(hAO_DB_PASSWORD)
	if passwordEnv != "" {
		password = passwordEnv
	} else {
		password = cfg.Db.Password
	}

	dsn := fmt.Sprintf("%v://%v:%v@%v:%v/%v", cfg.Db.Protocol, cfg.Db.User, password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Database)

	cfg.DSN = dsn
}
