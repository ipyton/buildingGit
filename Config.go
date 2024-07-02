package main

import (
	"os"
)

type Config struct {
	gitPath       string
	GLOBAL_CONFIG string
	SYSTEM_CONFIG string
	storage       map[string]string
}

func (*Config) readConfig(file *os.File) {
	//buf := bufio.NewReader(file)
	//for
}

func (config *Config) NewConfig(path string) *Config {
	global_config := "~/.gitconfig"
	system_config := "/etc/getconfig"
	system, _ := os.Open(system_config)
	global, _ := os.Open(global_config)
	local, _ := os.Open(path)
	config.readConfig(system)
	config.readConfig(global)
	config.readConfig(local)
	return &Config{}
}

func (config *Config) Get(key string) string {
	return config.storage[key]
}
