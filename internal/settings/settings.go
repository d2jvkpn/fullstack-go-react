package settings

import (
	"fmt"
	"strings"

	"github.com/d2jvkpn/go-web/pkg/wrap"
	"github.com/d2jvkpn/x-ai/pkg/chatgpt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	_Project    *viper.Viper
	GPTCli      *chatgpt.Client
	Tls         *TlsConfig
	Logger      *wrap.Logger
	TransLogger *zap.Logger // transaction
	ReqLogger   *zap.Logger // request
	AppLogger   *zap.Logger // application
	// DebugLogger    *zap.Logger // debug
)

func init() {
	// StartTs = time.Now().Unix()
}

type TlsConfig struct {
	Enable bool   `mapstructure:"enable"`
	Crt    string `mapstructure:"crt"`
	Key    string `mapstructure:"key"`
}

func NewTlsConfig(fp, key string) (config *TlsConfig, err error) {
	vp := viper.New()
	vp.SetConfigType("yaml")

	vp.SetConfigFile(fp)
	if err = vp.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("ReadInConfig(): %q, %w", fp, err)
	}

	config = new(TlsConfig)
	if err = vp.UnmarshalKey(key, config); err != nil {
		return nil, err
	}

	return config, nil
}

func SetProject(str string) (err error) {
	_Project = viper.New()
	_Project.SetConfigType("yaml")
	if err = _Project.ReadConfig(strings.NewReader(str)); err != nil {
		return err
	}

	_Project.SetDefault("cors_origins", "*")
	return nil
}

func GetCorsOrigins() string {
	return _Project.GetString("cors_origins")
}

func GetProject() string {
	return _Project.GetString("project")
}

func GetVersion() string {
	return _Project.GetString("version")
}

func GetConfig() string {
	return _Project.GetString("config")
}

func SetupLoggers() {
	TransLogger = Logger.Named("transaction")
	ReqLogger = Logger.Named("request")
	AppLogger = Logger.Named("application")
}
