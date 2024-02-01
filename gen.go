//go:build ignore

package main

import (
	"os"
	"text/template"

	"github.com/spf13/viper"
)

var TMPL_FILE = "go.tmpl"
var OUT_FILE = "main_gen.go"
var CONFIG_FILE = "menu.yaml"

func getConfig() (config any, err error) {
	viper.SetConfigFile(CONFIG_FILE)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func main() {

	tmpl, err := template.New(TMPL_FILE).ParseFiles(TMPL_FILE)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(OUT_FILE)
	if err != nil {
		panic(err)
	}

	config, err := getConfig()
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, config)
	if err != nil {
		panic(err)
	}

}

