/*
	Copyright (c) 2020-2022 Ivan B. Rybko
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
	"reflect"
	"testing"
)
//
func TestArgsToMap(t *testing.T) {
	//
	var expected = NewDict()
	expected["arg1"] = "val1"
	expected["arg2"] = "val2"
	expected["arg3"] = "val3"
	//
	var testlist = []string{
		"arg1=val1",
		"arg2=val2",
		"arg3=val3",
	}
	//
	got := NewDict()
	got.ArgsToMap(testlist)
	//
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got: %q, expected %q", got, expected)
	}
}	

