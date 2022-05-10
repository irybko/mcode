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
	"bytes"
	"encoding/json"
	"encoding/gob"
)
/*
	EXPORT FUNCTIONS
*/
//
// Convert Dict`s intstance to byte sequence
//
func (inst Dict) SerializeDict() []byte {
	//
	bin_buf := new(bytes.Buffer)
	// create a encoder object
	obj := gob.NewEncoder(bin_buf)
	// encode buffer and marshal it into a gob object
	obj.Encode(inst)
	return bin_buf.Bytes()
}
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
// Put Dict`s instance content to os.Args
//
func (inst Dict) DictToArgs() {
	//
	os.Args = inst.ToStrings()
}
//
// Put Dict`s instance content to os.Env
//
func (inst Dict) DictToEnv() {
	//
	idx := 0
	//
	mlen, keys := inst.GetKeys()
	//
	for idx < mlen {
		//
		k := keys[idx]
		v := inst.GetValue(k).(string)
		//
		os.Setenv(k,v)
		idx++
	}
}
 
