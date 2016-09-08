package config

import (
	"github.com/jinzhu/configor"
)

var Variables = struct {
	Application struct {
		Email    string `required:"true"`
		Password string `required:"true"`
		Token    string `required:"true"`
		Debug    bool   `required:"true"`
	}
}{}

func init() {
	configor.Load(&Variables, "config/config.json")
}
