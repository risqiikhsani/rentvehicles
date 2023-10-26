package controllers

import (
	"github.com/risqiikhsani/rentvehicles/configs"
)

var AppConf configs.MainConfig
var SecretConf configs.SecretsConfig

func SetAppConfig(config configs.MainConfig) {
	AppConf = config
}

func SetSecretConfig(config configs.SecretsConfig) {
	SecretConf = config
}
