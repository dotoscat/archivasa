package util

import (
	"strconv"
	"strings"
)

type Config map[string]string

func (c Config) GetString(field string) (string, bool) {
	value, ok := c[field]
	return value, ok
}

func (c Config) GetInt(field string) (int, error) {
	value, err := strconv.ParseInt(c[field], 10, 0)
	return int(value), err
}

func ReadConfigFromString(content string) Config {
	config := make(Config)
	lines := strings.Split(content, "\n")
	for _, row := range lines {
		if len(row) == 0 {
			continue
		}
		fieldValue := strings.Split(row, "=")
		field := strings.TrimSpace(fieldValue[0])
		if len(fieldValue) == 1 {
			config[field] = ""
		}
		value := strings.TrimSpace(fieldValue[1])
		config[field] = value
	}
	return config
}
