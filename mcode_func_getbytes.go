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
	"bytes"
	"encoding/gob"
)
//
//	Convert funcs
//
// Makes byte`s sequence of some type instance
func GetBytes(key interface{}) []byte {
	var buf bytes.Buffer
	    enc := gob.NewEncoder(&buf)
	    err := enc.Encode(key)
	    if err != nil {
        	log.Fatal(err)
    	    }
	return buf.Bytes()
}

