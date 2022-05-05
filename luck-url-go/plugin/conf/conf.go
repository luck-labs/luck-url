package conf

/**
 * @brief 加载配置项
 */

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
)

var (
	Viper        = viper.New()
	GlobalConfig *Config // GlobalConfig 全局对象，主要是配置文件配置的数据
)

func Init(confPath string) {
	if confPath == "" {
		Viper.SetConfigName("app") // 默认文件配置文件为app.toml
		_, fn, _, _ := runtime.Caller(0)
		confDir := filepath.Dir(fn)
		confPath = filepath.Join(confDir, "../../../conf")
		Viper.AddConfigPath(confPath)
	} else {
		Viper.SetConfigFile(confPath)
	}
	err := Viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("tag=InitConfError | fatal error config files: %s", err))
	}

	_, err = toml.DecodeFile(confPath, &GlobalConfig)
	GlobalConfig.HostName, _ = os.Hostname()
	pwd, _ := os.Getwd()
	GlobalConfig.AppPath = pwd
	if err != nil {
		panic(fmt.Errorf("tag=DecodeTomlFileError | fatal error config files: %s", err))
	}
}

type Config struct {
	// MySQL Config
	Mysql struct {
		Dsn             string `toml:"dsn"`
		MaxIdleConns    int    `toml:"max_idle_conns"`
		MaxOpenConns    int    `toml:"max_open_conns"`
		ConnMaxLifetime int    `toml:"conn_max_lifetime"`
		Shard           int    `toml:"shard"`
	} `toml:"mysql"`

	// Redis Config
	Redis struct {
		Address       []string `toml:"addrs"`
		Auth          string   `toml:"auth"`
		PoolSize      int      `toml:"pool_size"`
		MaxPoolSize   int64    `toml:"max_con"`
		ConnTimeOutMS int      `toml:"conn_timeout"`
		ReadTimeOut   int      `toml:"read_timeout"`
		WriteTimeOut  int      `toml:"write_timeout"`
	} `toml:"redis"`

	// Http Server
	HttpServer struct {
		Address     string `toml:"addr"`
		MaxConn     int    `toml:"max_conn"`
		ReadTimeout int    `toml:"read_timeout"`
	} `toml:"http_server"`

	// Pprof
	Pprof struct {
		Address string `toml:"addr"`
	} `toml:"pprof"`

	// RPC Server
	RpcServer struct {
		Address  string `toml:"addr"`
		MaxTotal int    `toml:"max_total"`
		MaxIdle  int    `toml:"max_idle"`
		MinIdle  int    `toml:"min_idle"`
	} `toml:"rpc_server"`

	// Log
	Log struct {
		InfoLogPath   string `toml:"info_log_path"`
		ErrorLogPath  string `toml:"error_log_path"`
		LogPath       string `toml:"log_path"`
		LogFormat     string `toml:"log_format"`
		LogMaxAge     string `toml:"log_max_age"`
		LogRotateTime string `toml:"log_rotate_time"`
	} `toml:"log"`

	// ShortUrl
	ShortUrl struct {
		BasePrefix string `toml:"base_prefix"`
	} `toml:"short_url"`

	// Snowflake
	Snowflake struct {
		TimeLength     int `toml:"time_length"`
		SequenceLength int `toml:"sequence_length"`
		MachineLength  int `toml:"machine_length"`
	} `toml:"snowflake"`

	HostName string `toml:"host_name"` // 机器名称
	AppPath  string `toml:"app_path"`  // 应用目录
}
