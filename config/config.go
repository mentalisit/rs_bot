package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	TokenT  string
	TokenD	string

	config *configStruct
)

type configStruct struct {
	TokenT     string `json : "TokenT"`
	TokenD	   string `json : "TokenD"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	TokenT = config.TokenT
	TokenD = config.TokenD

	return nil

}
