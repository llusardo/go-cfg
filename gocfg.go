package gocfg

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadConfiguration unmarshalls the content of "env_{env} file in out struct pointer"
func ReadConfiguration(env string, out interface{}) error {
	file, err := os.Open(fmt.Sprintf("env_%s.json", env))

	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(out)
	}
	return err
}

// MySQL connection representation
type MySQL struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Db       string `json:"db"`
}