package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Cfg Config

func InitConfig() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取配置路径发生错误 err =", err)
	}
	v := viper.New()
	v.AddConfigPath(dir + "/config")
	v.SetConfigType("yaml")
	v.SetConfigName("config.yaml")

	err = v.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置发生错误 err =", err)
		panic(err)
	}

	err = v.Unmarshal(&Cfg)
	if err != nil {
		fmt.Println("配置解析发生错误 err =", err)
	}
	fmt.Println(Cfg)
}

type Config struct {
	Server struct {
		Port string //web端口
		Name string //项目名
	}
	Mysql struct {
		Host     string
		Port     string
		User     string
		Password string
		DataBase string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
		Db       int
	}
}

//
//func banner() {
//	logger.SLog.Info("\n" + "\n" +
//		"  ___  ___  ________   ___  ________  ________  ________  ________           ________  ________  _____ ______   ________  ________  ________       ___    ___ \n|\\  \\|\\  \\|\\   ___  \\|\\  \\|\\   ____\\|\\   __  \\|\\   __  \\|\\   ___  \\        |\\   ____\\|\\   __  \\|\\   _ \\  _   \\|\\   __  \\|\\   __  \\|\\   ___  \\    |\\  \\  /  /|\n\\ \\  \\\\\\  \\ \\  \\\\ \\  \\ \\  \\ \\  \\___|\\ \\  \\|\\  \\ \\  \\|\\  \\ \\  \\\\ \\  \\       \\ \\  \\___|\\ \\  \\|\\  \\ \\  \\\\\\__\\ \\  \\ \\  \\|\\  \\ \\  \\|\\  \\ \\  \\\\ \\  \\   \\ \\  \\/  / /\n \\ \\  \\\\\\  \\ \\  \\\\ \\  \\ \\  \\ \\  \\    \\ \\  \\\\\\  \\ \\   _  _\\ \\  \\\\ \\  \\       \\ \\  \\    \\ \\  \\\\\\  \\ \\  \\\\|__| \\  \\ \\   ____\\ \\   __  \\ \\  \\\\ \\  \\   \\ \\    / / \n  \\ \\  \\\\\\  \\ \\  \\\\ \\  \\ \\  \\ \\  \\____\\ \\  \\\\\\  \\ \\  \\\\  \\\\ \\  \\\\ \\  \\       \\ \\  \\____\\ \\  \\\\\\  \\ \\  \\    \\ \\  \\ \\  \\___|\\ \\  \\ \\  \\ \\  \\\\ \\  \\   \\/  /  /  \n   \\ \\_______\\ \\__\\\\ \\__\\ \\__\\ \\_______\\ \\_______\\ \\__\\\\ _\\\\ \\__\\\\ \\__\\       \\ \\_______\\ \\_______\\ \\__\\    \\ \\__\\ \\__\\    \\ \\__\\ \\__\\ \\__\\\\ \\__\\__/  / /    \n    \\|_______|\\|__| \\|__|\\|__|\\|_______|\\|_______|\\|__|\\|__|\\|__| \\|__|        \\|_______|\\|_______|\\|__|     \\|__|\\|__|     \\|__|\\|__|\\|__| \\|__|\\___/ /     \n                                                                                                                                                \\|___|/      \n",
//	)
//}
