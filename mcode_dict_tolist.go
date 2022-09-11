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

import "log"
//
// Convert Dict`s instance to list 
//
func (inst Dict) ToList() [][]interface{} {
	//
	mlen, keys  := inst.GetKeys()
	index := 0
	log.Println("max: ", mlen)
	list  := make([][]interface{}, mlen)
	//
	for index < mlen {
		//
		k := keys[index]
		v := inst.GetValue(k)
		p := []interface{}{ k, v }
		list[index] = p
		index++
	}
	log.Println("current list content: ",list)
	return list
}

