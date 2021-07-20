/*
	Copyright (c) 2019-2021 Ivan B. Rybko
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
	"os"
    	"net/http"
)
//
// CORS Cross Origin Resource Sharing
//
func Corsheaders(method string,	resp http.ResponseWriter) {
	var acam string
	lm := len(method)
	AccessControlAllowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
	switch {
		case lm > 0:    acam = method
		case lm == 0:   acam = "POST, GET, OPTIONS, PUT, DELETE"
	}
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", acam)
	resp.Header().Set("Access-Control-Allow-Headers", AccessControlAllowHeaders)
}
//
// Get Content-Type by path
//
func GetContentType(path string) string {
	//
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	// Reset the read pointer if necessary.
	file.Seek(0, 0)
	// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
	return http.DetectContentType(buffer)
}
//
// GetExtName get extension name by content-type of path 
//
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

