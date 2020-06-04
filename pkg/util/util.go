/*
archivasa - a static web generator, and only that
Copyright (C) 2020 Oscar Triano García

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
