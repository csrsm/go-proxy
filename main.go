package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"webSocks5/client"
	"webSocks5/server"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Mode                        string `yaml:"Mode"`
	WsServer                    string `yaml:"WsServer"`
	ClientSocks5Port            int    `yaml:"ClientSocks5Port"`
	ClientJwtPrivateKeyFilePath string `yaml:"ClientJwtPrivateKeyFilePath"`
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "config.yaml", "Path to the configuration file")
	flag.Parse()

	// 读取配置文件
	configData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(configData, &config); err != nil { // 使用 yaml 包
		log.Fatalf("Error unmarshaling YAML data: %v", err)
	}
	fmt.Println(config)
	if config.Mode == "s" {
		server.Listen()
	} else if config.Mode == "c" {
		client.Listen(client.Config{WsServerAddr: config.WsServer, Socks5Port: config.ClientSocks5Port, JwtPrivateKeyFilePath: config.ClientJwtPrivateKeyFilePath})
	} else {
		fmt.Println("不支持该运行模式")
	}
}
