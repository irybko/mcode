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
	"fmt"	
	"testing"
	"reflect"
)
//
func TestDictSubst(t *testing.T) {
	//
	expected := "insert into persons(name,lastname,age,occupation) values(John,Smith,34,mechanic);"
	//
	tmplstr := "insert into persons(name,lastname,age,occupation) values(v1,v2,v3,v4);"
	//
	kv := NewDict()
	kv["v1"] = "John"
	kv["v2"] = "Smith"
	kv["v3"] = fmt.Sprintf("%d", 34)
	kv["v4"] = "mechanic"
	//
	got	 := kv.DictSubst(tmplstr)
	//
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got: %s, expected: %s", got, expected)
	}
}
//
func ExampleDictSubst() {
	//
	tmplstr := `insert into persons(name,lastname,age,occupation) values(v1,v2,v3,v4);`
	kv := NewDict()
	kv["v1"] = "John"
	kv["v2"] = "Smith"
	kv["v3"] = 34
	kv["v4"] = "mechanic"
	fmt.Printf("result of substitution %s", kv.DictSubst(tmplstr))
	// Output
	// "insert into persons(name,lastname,age,occupation) values(John,Smith,34,mechanic)"
}

