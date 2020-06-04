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

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dotoscat/archivasa/pkg/builder"
	"github.com/dotoscat/archivasa/pkg/config"
	"github.com/dotoscat/archivasa/pkg/content"
	"github.com/dotoscat/archivasa/pkg/theme"
)

func main() {
	cwd, error := os.Getwd()
	if error != nil {
		log.Fatal(error)
	}
	config := config.Read(cwd)
	theme := theme.Load(cwd)
	content := content.Read(cwd)
	fmt.Println(theme)
	fmt.Println(config.Title)
	fmt.Println(config.PostsPerPage)
	fmt.Println(content.Pages)
	fmt.Println(content.Posts)
	fmt.Println("Current working directory", cwd)
	builder.Run(config, &content, theme)
}
