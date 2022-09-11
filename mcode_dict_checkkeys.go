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

//
// Check Dict`s instance keys
//
func (inst Dict) CheckKeys(hdrs []string) bool {
	var flag bool
	step := []string{}
	//
	_, keys := inst.GetKeys()
	//
	//
	if len(hdrs) == 0 {
		for _, k := range keys {
			step = append(step, k)
		}
		flag = true
	} else {
		if CmpPair(hdrs, keys) {
			flag = true
		} else {
			flag = false
		} 
	}
	return flag
}
