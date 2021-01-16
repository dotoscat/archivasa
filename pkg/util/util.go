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
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

//CopyFolder copies a folder from src to dst. You can optionally exclude a directory with exclude
func CopyFolder(src, dst, exclude string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		os.Mkdir(dst, os.ModePerm)
	}
	for _, file := range files {
		fmt.Println("copy file", file.Name())
		if file.IsDir() && file.Name() != exclude {
			folderSrc := filepath.Join(src, file.Name())
			folderDst := filepath.Join(dst, file.Name())
			CopyFolder(folderSrc, folderDst, "")
			continue
		} else if file.IsDir() && file.Name() == exclude {
			continue
		}
		fileSrc := filepath.Join(src, file.Name())
		fileDst := filepath.Join(dst, file.Name())
		fmt.Println("copy:", fileSrc, fileDst)
		CopyFile(fileSrc, fileDst)
	}
	return nil
}

//CopyFile copies a file from src to dst
func CopyFile(src, dst string) {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatalln(err)
	}
	dstFile, err := os.Create(dst)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		log.Fatalln(err)
	}
}
