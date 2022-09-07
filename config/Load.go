package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (

	// Private variables
	config *configStruct
)

type configStruct struct {
	Server struct {
		Mysql struct {
			Sqluser     string `json:"SQLUser"`
			Sqlhost     string `json:"SQLHost"`
			Sqlpass     string `json:"SQLPass"`
			Sqldatabase string `json:"SQLDatabase"`
		} `json:"Mysql"`
		Master struct {
			Enabled bool `json:"Enabled"`
			Auth    struct {
				Maxtries int `json:"MaxTries"`
				Captcha struct {
					Enabled  bool `json:"Enabled"`
					Mfa      bool `json:"mfa"`
					Maxtries int  `json:"MaxTries"`
				} `json:"capcha"`
			} `json:"Auth"`
			Host string `json:"Host"`
		} `json:"Master"`
	} `json:"Server"`
}

func ReadConfig() error {
	log.Println(" [Config Worker] | [Started]")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		return err
	}

	MasterHost = config.Server.Master.Host
	MasterEnabled = config.Server.Master.Enabled
	MaxLoginTries = config.Server.Master.Auth.Maxtries
	CapchaEnabled = config.Server.Master.Auth.Captcha.Enabled
	SQLHost = config.Server.Mysql.Sqlhost
	SQLUser = config.Server.Mysql.Sqluser
	SQLPassword = config.Server.Mysql.Sqlpass
	SQLDatabase = config.Server.Mysql.Sqldatabase

	if config.Server.Master.Auth.Captcha.Enabled == true {
		log.Println(" [Capcha worker] | [Started]")
	}

	return nil
}
