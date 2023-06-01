package configload 

import (
  "io/ioutil"
  "gopkg.in/yaml.v2"

//  "log"
)

type Conf struct {
  ServerAddress string  `yaml:"server_address"`
  ServerPort int  `yaml:"server_port"`
  AgentMode string `yaml:"agent_mode"`
  AgentPort int `yaml:"agent_port"`
  HostName string `yaml:"host_name"`
  HostGroup string `yaml:"host_group"`
  AgentInterval string `yaml:"agent_interval"`
  AgentRetry int `yaml:"agent_retry"`
 
}


func LoadConfig(filename string) (*Conf, error) {
	// Read the YAML file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Create a new Conf instance
	config := &Conf{}

	// Unmarshal the YAML data into the Conf struct
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
