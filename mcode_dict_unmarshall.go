/*
	Copyright (c) 2020-2021 Ivan B. Rybko
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
	"net/http"
	"io/ioutil"
	"encoding/json"
)
//
// Read request`s JSON formatted message string into Dict`s instance
//
func (inst Dict) UnmarshalRequest(request *http.Request) {
    	//
    	if request.Method == http.MethodPost {
		MemoryToRead, err := ioutil.ReadAll(request.Body)
	        if err != nil {
			log.Fatal("Error was occured during an attempt to read request Body: ", err)
	        }
        	//
	        err = json.Unmarshal(MemoryToRead, &inst)
        	if err != nil {
			log.Fatal("Error was occured during an attempt to parse json object: ", err)
	        }
    	}
}

