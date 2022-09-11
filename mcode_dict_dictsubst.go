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
	"strings"
)
//
// Substitutes templates with values dictionaries  
//
func (inst Dict) DictSubst(tmpl string) string {
	//
	for k, v := range inst {
		//
		if strings.Index(tmpl, k) != -1 {
			//
			tmpl = strings.ReplaceAll(tmpl, k, v.(string))
		}
	}
	return tmpl
}






