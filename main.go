package main

import (
	"os"

	"github.com/olegmymrin/mk1/service"
	"gopkg.in/yaml.v3"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	var config service.Config
	configFileData, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(configFileData, &config); err != nil {
		panic(err)
	}
	srv := service.NewService(&config)
	srv.ListenAndServe()
}
