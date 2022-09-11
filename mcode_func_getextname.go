/*
	Copyright (c) 2019-2022 Ivan B. Rybko
	=====================================

	This program is a part of golang middleware functions library package mcode.

	This program is free software: you can redistribute it and/or modify 
	it under the terms of the GNU General Public License as published by 
	the Free Software Foundation, either version 3 of the License, or 
	(at your option) any later version. 

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of 
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. 
	See the GNU General Public License for more details.

	You should have received a copy of the GNU General Public License 
	along with this program. If not, see <https://www.gnu.org/licenses/>.

*/
package mcode

import (
	"log"
)

// GetExtName get extension name by content-type of path 
func GetExtName(path string) string {
	//
	var extname string = ""
	var flag bool = true
	var contenttypes = map[string]string{
		//
		"text/plain":			"txt",
		"text/html": 			"html",
		"text/css":  			"css",
		"text/javascript": 		"js",
		"text/csv":  			"csv",
		"application/json": 		"json",
		"application/octet-stream": 	"bin",
	}
	//
	ctype := GetContentType(path)
	extname, flag = contenttypes[ctype] 
	if !flag {
		log.Println("Error: No extension name found")
	}
	return extname
}

