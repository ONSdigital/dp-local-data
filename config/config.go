package config

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MongoURL    string   `yaml:"mongo-url"`
	Neo4jURL    string   `yaml:"neo4j-url"`
	MongoDBs    []string `yaml:"mongo-dbs"`
	Hierarchies []string `yaml:"hierarchies"`
	Codelists   []string `yaml:"codelists"`
}

func Get() (*Config, error) {
	path := os.Getenv("DP_LOCAL_DATA")
	if len(path) == 0 {
		return nil, errors.New("no config file specified")
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
