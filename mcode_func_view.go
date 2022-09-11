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
	"bytes"
	"text/template"
)

/*
    VIEW INTERFACE FUNCTIONS
*/
var (
        jsontmpl        map[string]map[string]string
        prepstmts       map[string]map[string]string
)
//
type View interface {
        Subst(tmpltext string, kv Dict) string
        Listsubst(maintmpl string, embedtmpl string, key string, lstkv []Dict) string
}
//
func Subst(tmpltext string, kv Dict) string {
        t, _ := template.New("").Parse(tmpltext)
        buf := new(bytes.Buffer)
        t.Execute(buf, &kv)
        return buf.String()
}
//
func Listsubst(maintmpl string, embedtmpl string, key string, lstkv []Dict) string {
        var count int = 0
        arr := make([]string, 0, len(lstkv))

        for count < len(lstkv) {
                arr = append(arr, Subst(embedtmpl, lstkv[count]))
                count++
        }
        step := make(Dict)
        step[key] = arr
        return Subst(maintmpl, step)
}

