package datasaur

import (
	"encoding/json"
	"os"
)

type Host struct {
	Name string `json:"name"`
	Host string `json:"host"`
}

func GetHostsFromConfig(cfgPath string) map[string]Host {
	read, err := os.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}

	parsed := make(map[string]string)
	err = json.Unmarshal(read, &parsed)
	if err != nil {
		panic(err)
	}

	result := make(map[string]Host)
	for k, v := range parsed {
		result[k] = Host{
			Name: k,
			Host: v,
		}
	}
	return result
}
