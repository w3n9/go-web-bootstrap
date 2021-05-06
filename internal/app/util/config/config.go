package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ReadConfigDirFromProgramArgument() string {
	dir := flag.String("config", "", "the config directory of this app")
	flag.Parse()
	return *dir
}

func Load(filename string, target interface{}) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, target)
	return err
}
