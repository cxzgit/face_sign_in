package globals

import "time"

// 和配置相关的结构体

// DatabaseConfig mysql配置
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	Password     string        `yaml:"password"`
	DB           int           `yaml:"db"`
	PoolSize     int           `yaml:"pool_size"`      // Redis 连接池中的最大连接数
	MinIdleConns int           `yaml:"min_idle_conns"` // Redis 连接池中的最小空闲连接数
	IdleTimeout  time.Duration `yaml:"idle_timeout"`   // Redis 连接池中空闲连接的最大超时时间
	DialTimeout  time.Duration `yaml:"dial_timeout"`   // Redis 连接超时
	ReadTimeout  time.Duration `yaml:"read_timeout"`   // Redis 读取数据超时
	WriteTimeout time.Duration `yaml:"write_timeout"`  // Redis 写入数据超时
	MaxRetries   int           `yaml:"max_retries"`    // Redis 最大重试次数
}

// App 配置
type App struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Domain string `yaml:"domain"` // 域名
}

// LogConfig 日志配置
type LogConfig struct {
	Level   string `yaml:"level"`
	LogPath string `yaml:"logPath"`
	AppName string `yaml:"appName"`
}

// Config 总配置
type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	App      App            `yaml:"app"`
	Log      LogConfig      `yaml:"log"`
}
