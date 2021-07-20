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

/*
	DATA TYPES AND VARIABLES
*/
type Dictionary struct {

	Ru string
	En string 

}
//
func ChooseLang(lang string,viewdict map[string]Dictionary) map[string]string {

	kv := make(map[string]string) 
	if lang == "ru" {	for k,v := range viewdict {	kv[k] = v.Ru 	}}
	if lang == "en" {	for k,v := range viewdict {	kv[k] = v.En	}}
	return kv

}

