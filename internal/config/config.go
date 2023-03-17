package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// Build information -ldflags .
const (
	version    string = "dev"
	commitHash string = "-"
)

var cfg *Config

// GetConfigInstance returns service config
func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

// Database - contains all parameters database connection.
type Database struct {
	Host       string  `yaml:"host"`
	Port       string  `yaml:"port"`
	User       string  `yaml:"user"`
	Password   string  `yaml:"password"`
	Migrations string  `yaml:"migrations"`
	Name       string  `yaml:"name"`
	SslMode    string  `yaml:"sslmode"`
	Driver     string  `yaml:"driver"`
	DBConns    DBConns `yaml:"connection"`
}

type DBConns struct {
	DSN             string        `yaml:"dsn"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	ConnMaxIdleTime time.Duration `yaml:"connMaxIdleTime"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}

func (d Database) GetDSN() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		d.Host,
		d.Port,
		d.User,
		d.Password,
		d.Name,
		d.SslMode,
	)
}

func (d Database) GetMaxOpenConns() int {
	return d.DBConns.MaxOpenConns
}

func (d Database) GetMaxIdleConns() int {
	return d.DBConns.MaxIdleConns
}

func (d Database) GetConnMaxIdleTime() time.Duration {
	return d.DBConns.ConnMaxIdleTime
}

func (d Database) GetConnMaxLifetime() time.Duration {
	return d.DBConns.ConnMaxLifetime
}

// Grpc - contains parameter address grpc.
type Grpc struct {
	Port              int    `yaml:"port"`
	MaxConnectionIdle int64  `yaml:"maxConnectionIdle"`
	Timeout           int64  `yaml:"timeout"`
	MaxConnectionAge  int64  `yaml:"maxConnectionAge"`
	Host              string `yaml:"host"`
}

// Rest - contains parameter rest json connection.
type Rest struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	Version     string
	CommitHash  string
}

// Metrics - contains all parameters metrics information.
type Metrics struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

// Jaeger - contains all parameters metrics information.
type Jaeger struct {
	Service string `yaml:"service"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
}

// Kafka - contains all parameters kafka information.
type Kafka struct {
	Capacity uint64   `yaml:"capacity"`
	Topic    string   `yaml:"topic"`
	GroupID  string   `yaml:"groupId"`
	Brokers  []string `yaml:"brokers"`
}

// Status config for service.
type Status struct {
	Port          int    `yaml:"port"`
	Host          string `yaml:"host"`
	VersionPath   string `yaml:"versionPath"`
	LivenessPath  string `yaml:"livenessPath"`
	ReadinessPath string `yaml:"readinessPath"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project  Project  `yaml:"project"`
	Grpc     Grpc     `yaml:"grpc"`
	Rest     Rest     `yaml:"rest"`
	Database Database `yaml:"database"`
	Metrics  Metrics  `yaml:"metrics"`
	Jaeger   Jaeger   `yaml:"jaeger"`
	Kafka    Kafka    `yaml:"kafka"`
	Status   Status   `yaml:"status"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(filePath string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash

	return nil
}
