package main

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// DefaultPort is used whtn no valid port is found on the config file
const DefaultPort = 5000

// Replacement defines the source and target URL host, along with an optional path
type Replacement struct {
	HostSource string `yaml:"source"`
	HostTarget string `yaml:"target"`
	Path       string `yaml:"path"`
}

// Config holds the trigger definitions
type Config struct {
	Cmd struct {
		Port    int  `yaml:"port"`
		Verbose bool `yaml:"verbose"`
	} `yaml:"cmd"`
	Replacements []Replacement `yaml:"replacements"`
}

// ReadConfig parses `config.yaml` and returns a struct with the desired config
func ReadConfig(filePath string) (Config, error) {
	conf := Config{}

	// Read config.yaml
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		return conf, err
	}

	if conf.Cmd.Port == 0 {
		conf.Cmd.Port = DefaultPort
	}
	if len(conf.Replacements) < 1 {
		return conf, errors.Errorf("No replacements are defined")
	}

	// Append / at the end of the paths
	for _, replacement := range conf.Replacements {
		if replacement.Path == "" {
			replacement.Path = "/"
		}
	}

	err = checkConfig(conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func checkConfig(conf Config) error {
	var srcReplacements []string

	for idx, replacement := range conf.Replacements {
		if replacement.HostSource == "" {
			return errors.Errorf("[CONFIG] The source replacement at index %d is empty", idx)
		} else if replacement.HostTarget == "" {
			return errors.Errorf("[CONFIG] The target replacement at index %d is empty", idx)
		}
		for _, prevSrcReplacement := range srcReplacements {
			if prevSrcReplacement == replacement.HostSource {
				return errors.Errorf("[CONFIG] The source replacement %s is defined multiple times", prevSrcReplacement)
			}
		}
		srcReplacements = append(srcReplacements, replacement.HostSource)
	}
	return nil
}
