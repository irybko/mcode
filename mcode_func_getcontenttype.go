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
	"os"
    	"net/http"
)

// Get Content-Type by path
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

