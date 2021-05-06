package main

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

func main() {
	println("Object to YAML(Marshal)")
	p := Person{"Talha", "Asan", 20}

	y, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}

	println(string(y))

	println("Reading and printing config.yaml file")
	fileName := "./config.yaml"
	source, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}
	println(string(source))

	println("YAML to Object(Unmarshal)")
	c := string(source)

	o, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}

	println(string(o))

}

type Person struct {
	FirstName string `yaml:"first_name"`
	LastName  string `yaml:"last_name"`
	Age       int    `yaml:"age"`
}

type Config struct {
	User     string   `yaml:"user"`
	Database string   `yaml:"database"`
	Port     int      `yaml:"port"`
	Server   string   `yaml:"server"`
	Settings []string `yaml:"settings"`
}
