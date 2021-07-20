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
	"log"
	"net/url"
    	"net/http"
	//
	"testing"
)
//
func TestParseURL(t *testing.T) {
	//
	expected := make(map[string]string)
	expected["scheme"] = "http"
	expected["host"]   = "localhost"
	expected["port"]   = "9000"
	expected["path"]   = "somepath"
	//
	got	 := ParseURL("http://localhost:9000/somepath")
	count	 := 0
	//
	for k1, v1 := range got {
		for k2, v2 := range expected {
			if k1 == k2 && v1 == v2 {
				count++
			}
		}
	}
	//
	if count != len(got) {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}
//
func TestFormToMap(t *testing.T) {
	//
	expected := make(map[string]string{
		"k1": "v1",
		"k2": "v2",
	})
	//
	got 	:= make(map[string]string)
	var keys = []string{ "k1", "k2" }		
	//
	formdata := url.Values{
		"k1": {"v1"},
		"k2": {"v2"},
	}
	//
	formpath := "/testformtomap"
	//
	go func() {
		http.HanleFunc(formpath, func(req *http.Request, resp http.ResponseWriter) {
			//
			FormToMap(req,keys,got)
		})
		http.ListenAndServe(":9000",nil)
	}()
	//
	_, err := http.PostForm(formpath, formdata)
	if err != nil {
		log.Fatal("During http.PostForm execution an error was occured: \n",err)
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}

