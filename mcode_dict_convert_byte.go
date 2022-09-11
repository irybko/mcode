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
	"bytes"
	"encoding/gob"
)
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

