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

package content

import "path/filepath"

// Webpage represents a internet document with its URL
// A webpage belongs to a website
type Webpage struct {
	*Website
	Url        string
	outputPath string
}

func NewWebpage(site *Website, url, outputPath string) *Webpage {
	return &Webpage{site, url, outputPath}
}

func (w *Webpage) BuildURL(prefix, base string) {
	w.Url = filepath.Join(prefix, base)
}

func (w *Webpage) BuildOutputPath(prefix, base string) {
	w.outputPath = filepath.Join(prefix, base)
}

func (w *Webpage) URL() string {
	return w.Url
}

func (w *Webpage) OutputPath() string {
	return w.outputPath
}
