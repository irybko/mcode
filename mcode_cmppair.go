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
	"strings"
)
/*
	DATA TYPES AND VARIABLES
*/
// Compares two instances of same type. 
func CmpPair(pm1 interface{}, pm2 interface{}) bool {
	//
	var flag bool = true
	//
	switch t1 := pm1.(type) {
	case []byte:
		switch t2 := pm2.(type) {
		case []byte:
			//
			l1 := len(t1)
			l2 := len(t2)
			//
			if l1 != l2 {
				return false
			}
			//
			count := 0
			for _, v1 := range t1 {
				for _, v2 := range t2 {
					if v1 == v2 { count++ }
				}
			}
			//
			flag = count == l1 

		}
	case string:
		switch t2 := pm2.(type) {
		case string:
			step1 := []byte(t1)
			step2 := []byte(t2)
		 	//
			flag = CmpPair(step1, step2)
		}
	case []string:
		switch t2 := pm2.(type) {
		case []string:
			step1 := strings.Join(t1,",")
			step2 := strings.Join(t2,",")
			flag = CmpPair(step1,step2)
		}
	}
	return flag
}


