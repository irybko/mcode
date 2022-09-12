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
	"encoding/json"
)
//
// Encode Dict`s instance into JSON-formatted string
//
func (inst Dict) EncodeJson() []byte {
	// encode the data     
    	jsob, err := json.Marshal(inst)
    	if err != nil {
       		panic(err)
    	}
	return jsob
}
//
// Decode JSON-formatted string into Dict`s instance
//
func (inst Dict) DecodeJson(data []byte) error {
	//
	var record interface{}
	err := json.Unmarshal(data, &record)
    	if err != nil {
        	return err
    	}
	inst.GetStructFields(record)
	return err
}

