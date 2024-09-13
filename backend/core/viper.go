package core

import (
	"flag"
	"fmt"
	"gin-pro/core/internal"
	"gin-pro/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "shell to write config")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("当前gin模式:%s,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("当前gin模式:%s,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("当前gin模式:%s,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else {
				config = configEnv
				fmt.Printf("当前使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			}
		} else {
			fmt.Printf("当前使用的是命令行-c 传递的参数，config的路径为%s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("当前使用的是Viper()传递的参数，config的路径为%s\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取config的yaml环境变量配置失败:%s,\n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println("Error unmarshalling config:", err)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println("Error unmarshalling config:", err)
	}
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
