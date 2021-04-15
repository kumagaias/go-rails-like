package config

import (
	"github.com/spf13/viper"
	"runtime"
	"strings"
)

var v *viper.Viper

func Init(env string) {
	v = viper.New()
	v.SetConfigFile("yaml")
	v.SetConfigName(env)

	_, filePath, _, _ := runtime.Caller(1)
	projectPath := filePath[:strings.Index(filePath, "opb/")]
	v.AddConfigPath(projectPath + "opb/config/environments/")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
}

func Get(key string) string {
	return v.GetString(key)
}
