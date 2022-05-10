/*
	Copyright (c) 2020-2022 Ivan B. Rybko
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
	"os"
	"log"
	"bytes"
	"strings"
	"encoding/json"
	"encoding/gob"
)
/*
	IMPORT FUNCTIONS
*/
//
// Convert byte sequence to Dict`s instance
//
func (inst Dict) DeserializeDict(tmp []byte) {
	//
	buff := bytes.NewBuffer(tmp)
	// creates a decoder object
	gobobjdec := gob.NewDecoder(buff)
	// decodes buffer and unmarshals it into a Message struct
	gobobjdec.Decode(inst)
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
//
// Convert list arguments to map, each element is a string which may contain or not '=' delimiter 
//
func (inst Dict) ArgsToMap(cmdline bool) {
	//
	res	:= make(map[string]string) 
	// 
	var lst = []string{}
	//
	if cmdline {
		lst = os.Args
	}
	lst = os.Environ()
	idx := 0
	for idx < len(lst) {
		v := lst[idx] 
		//
		if strings.Index(v,"=") != -1 {
			//
			step  	:= strings.Split(v,"=")
			key   	:= step[0]
			value 	:= step[1]
			res[key] = value
		} 
		idx++
	}
	log.Println("argstomap: ", res)
	inst.CopyContent(res)
}


