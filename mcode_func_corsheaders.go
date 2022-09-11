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
    	"net/http"
)
// CORS Cross Origin Resource Sharing
func Corsheaders(method string,	resp http.ResponseWriter) {
	var acam string
	lm := len(method)
	AccessControlAllowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
	switch {
		case lm > 0:    acam = method
		case lm == 0:   acam = "POST, GET, OPTIONS, PUT, DELETE"
	}
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", acam)
	resp.Header().Set("Access-Control-Allow-Headers", AccessControlAllowHeaders)
}
