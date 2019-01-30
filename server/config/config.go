package config

import (
    "fmt"
    "github.com/spf13/viper"
)

func InitConfig(path string, name string) {
    fmt.Println(path)
    viper.SetConfigName(name)
    viper.AddConfigPath(path)
    viper.AddConfigPath(".")
    viper.SetEnvPrefix("MAGNUX_")
    viper.AutomaticEnv()
    viper.SetConfigType("yaml")
    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    fmt.Println(viper.IsSet("test"))
}
