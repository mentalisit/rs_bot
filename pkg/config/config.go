package config

import (
	"fmt"
	"os"
)

type ConfigEnv struct {
	TokenT, TokenD                       string
	Username, Password, Hostname, Dbname string
}

func InitEnv() (ConfigEnv, bool) {
	fmt.Println("init", os.Getenv("TOKENT"))
	TokenT := os.Getenv("TOKENT")
	fmt.Println("init2", TokenT)
	TokenD := os.Getenv("TokenD")
	fmt.Println("init3", TokenD)
	username := os.Getenv("dbUsername")
	fmt.Println("init4", username)
	password := os.Getenv("dbPassword")
	hostname := os.Getenv("dbHostname")
	dbname := os.Getenv("Dbname")
	return ConfigEnv{
		TokenT:   TokenT,
		TokenD:   TokenD,
		Username: username,
		Password: password,
		Hostname: hostname,
		Dbname:   dbname,
	}, true
}
