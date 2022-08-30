package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/latoken/bridge-backend-service/src/models"
	"github.com/latoken/bridge-backend-service/src/service/storage"

	"github.com/spf13/viper"
)

// Config ...
type Config interface {
	ReadServiceConfig() string
	ReadWorkersConfig() []*models.WorkerConfig
	ReadLachainConfig() *models.WorkerConfig
	ReadFetcherConfig() []*models.FetcherConfig
	ReadDBConfig() *models.StorageConfig
	ReadResourceIDs() []*storage.ResourceId
	ReadChains() []string
	GetString(key string) string
	GetStringMap(key string) map[string]string
	GetInt64(key string) int64
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetStringSlice(key string) []string
	Init()
}

type viperConfig struct {
}

func (v *viperConfig) Init() {
	viper.AutomaticEnv()
	// viper.AddConfigPath(os.Getenv("FILE_PATH"))
	viper.AddConfigPath("../config-files")
	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	// viper.SetConfigName(os.Getenv("FILE_NAME"))
	viper.SetConfigName(`config.json`)
	if _, err := os.Stat("./config.json.local"); !os.IsNotExist(err) {
		viper.SetConfigFile(`config.json.local`)
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (v *viperConfig) GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func (v *viperConfig) GetStringMap(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func (v *viperConfig) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// NewViperConfig creates new viper for reading config.json
func NewViperConfig() Config {
	v := &viperConfig{}
	v.Init()
	return v
}
